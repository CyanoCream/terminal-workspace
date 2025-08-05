package service

import (
	"errors"
	"fmt"
	"time"

	terminalPort "terminal/internal/terminal/port"
	"terminal/internal/transaction/domain"
	transactionPort "terminal/internal/transaction/port"
	userPort "terminal/internal/user/port"
)

type transactionService struct {
	repo        transactionPort.TransactionRepository
	terminalSvc terminalPort.TerminalService
	userRepo    userPort.UserRepository
}

func NewTransactionService(
	repo transactionPort.TransactionRepository,
	terminalSvc terminalPort.TerminalService,
	userRepo userPort.UserRepository,
) transactionPort.TransactionService {
	return &transactionService{
		repo:        repo,
		terminalSvc: terminalSvc,
		userRepo:    userRepo,
	}
}

func (s *transactionService) CheckIn(cardNumber string, gateID uint) error {
	card, err := s.userRepo.FindCardByNumber(cardNumber)
	if err != nil {
		return errors.New("card not found")
	}

	activeCheckin, err := s.repo.FindActiveCheckIn(card.ID)
	if err == nil && activeCheckin != nil {
		return errors.New("already checked in")
	}

	transaction := &domain.Transaction{
		CardID:      card.ID,
		GateID:      gateID,
		Type:        "checkin",
		Timestamp:   time.Now().Unix(),
		IsCompleted: false,
		IsSynced:    true,
	}

	return s.repo.Create(transaction)
}

func (s *transactionService) CheckOut(cardNumber string, gateID uint) error {
	card, err := s.userRepo.FindCardByNumber(cardNumber)
	if err != nil {
		return errors.New("card not found")
	}

	checkin, err := s.repo.FindActiveCheckIn(card.ID)
	if err != nil || checkin == nil {
		return errors.New("no active checkin found")
	}

	fare, err := s.calculateFare(checkin.GateID, gateID)
	if err != nil {
		return fmt.Errorf("failed to calculate fare: %v", err)
	}

	if card.Balance < fare {
		return errors.New("insufficient balance")
	}

	card.Balance -= fare
	if err := s.userRepo.UpdateCardBalance(card, fare); err != nil {
		return errors.New("failed to update card balance")
	}

	checkout := &domain.Transaction{
		CardID:       card.ID,
		GateID:       gateID,
		Type:         "checkout",
		TerminalFrom: &checkin.GateID,
		Amount:       &fare,
		Timestamp:    time.Now().Unix(),
		IsCompleted:  true,
		IsSynced:     true,
	}

	checkin.IsCompleted = true
	if err := s.repo.Create(checkout); err != nil {
		return errors.New("failed to record checkout")
	}

	return nil
}

func (s *transactionService) calculateFare(fromGateID, toGateID uint) (float64, error) {
	fromGate, err := s.terminalSvc.FindGateByID(fromGateID)
	if err != nil {
		return 0, err
	}

	toGate, err := s.terminalSvc.FindGateByID(toGateID)
	if err != nil {
		return 0, err
	}

	fare, err := s.terminalSvc.GetTerminalPricing(fromGate.TerminalID, toGate.TerminalID)
	if err != nil {
		return 0, err
	}

	if s.isPeakHour(time.Now()) {
		fare *= 1.2
	}

	return fare, nil
}

func (s *transactionService) isPeakHour(t time.Time) bool {
	hour := t.Hour()
	return (hour >= 7 && hour < 9) || (hour >= 16 && hour < 19)
}

func (s *transactionService) SyncTransactions(transactions []domain.OfflineTransaction) error {
	for _, tx := range transactions {
		card, err := s.userRepo.FindCardByNumber(tx.CardNumber)
		if err != nil {
			continue
		}

		transaction := &domain.Transaction{
			CardID:      card.ID,
			GateID:      tx.GateID,
			Type:        tx.Type,
			Timestamp:   tx.Timestamp,
			IsSynced:    true,
			ReferenceID: fmt.Sprintf("offline-%d", time.Now().UnixNano()),
		}

		if tx.Type == "checkout" && tx.Amount > 0 {
			transaction.Amount = &tx.Amount
			transaction.IsCompleted = true
		}

		if err := s.repo.Create(transaction); err != nil {
			continue
		}
	}
	return nil
}

func (s *transactionService) GetTransactionHistory(userID uint) ([]domain.Transaction, error) {
	return s.repo.GetUserTransactions(userID)
}
