package service

import (
	"errors"
	"terminal/internal/user/domain"
	"terminal/internal/user/port"
)

type userService struct {
	repo port.UserRepository
}

func NewUserService(repo port.UserRepository) port.UserService {
	return &userService{repo: repo}
}

func (s *userService) GetProfile(userID uint) (*domain.User, error) {
	return s.repo.FindByID(userID)
}

func (s *userService) CreateCard(userID uint, card *domain.Card) error {
	card.UserID = userID
	return s.repo.CreateCard(card)
}

func (s *userService) TopUpCard(userID uint, cardNumber string, amount float64) error {
	card, err := s.repo.FindCardByNumber(cardNumber)
	if err != nil {
		return err
	}

	if card.UserID != userID {
		return errors.New("unauthorized")
	}

	return s.repo.UpdateCardBalance(card, amount)
}

func (s *userService) GetCardBalance(userID uint, cardNumber string) (float64, error) {
	card, err := s.repo.FindCardByNumber(cardNumber)
	if err != nil {
		return 0, err
	}

	if card.UserID != userID {
		return 0, errors.New("unauthorized")
	}

	return card.Balance, nil
}
