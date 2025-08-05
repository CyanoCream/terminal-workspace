package port

import (
	"github.com/gofiber/fiber/v2"
	"terminal/internal/transaction/domain"
)

type TransactionService interface {
	CheckIn(cardNumber string, gateID uint) error
	CheckOut(cardNumber string, gateID uint) error
	SyncTransactions(transactions []domain.OfflineTransaction) error
	GetTransactionHistory(userID uint) ([]domain.Transaction, error)
}

type TransactionHandler interface {
	CheckIn(c *fiber.Ctx) error
	CheckOut(c *fiber.Ctx) error
	SyncTransactions(c *fiber.Ctx) error
	GetTransactionHistory(c *fiber.Ctx) error
}
