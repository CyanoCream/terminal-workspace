package port

import "terminal/internal/transaction/domain"

type TransactionRepository interface {
	Create(transaction *domain.Transaction) error
	FindActiveCheckIn(cardID uint) (*domain.Transaction, error)
	GetUserTransactions(userID uint) ([]domain.Transaction, error)
}
