package repository

import (
	"gorm.io/gorm"
	"terminal/internal/transaction/domain"
	"terminal/internal/transaction/port"
)

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) port.TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) Create(transaction *domain.Transaction) error {
	return r.db.Create(transaction).Error
}

func (r *transactionRepository) FindActiveCheckIn(cardID uint) (*domain.Transaction, error) {
	var transaction domain.Transaction
	err := r.db.Where("card_id = ? AND type = 'checkin' AND is_completed = false", cardID).First(&transaction).Error
	return &transaction, err
}

func (r *transactionRepository) GetUserTransactions(userID uint) ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	err := r.db.Joins("JOIN cards ON cards.id = transactions.card_id").
		Where("cards.user_id = ?", userID).
		Find(&transactions).Error
	return transactions, err
}
