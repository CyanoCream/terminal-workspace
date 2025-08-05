package port

import (
	"github.com/gofiber/fiber/v2"
	"terminal/internal/user/domain"
)

type UserService interface {
	GetProfile(userID uint) (*domain.User, error)
	CreateCard(userID uint, card *domain.Card) error
	TopUpCard(userID uint, cardNumber string, amount float64) error
	GetCardBalance(userID uint, cardNumber string) (float64, error)
}

type UserHandler interface {
	GetProfile(c *fiber.Ctx) error
	CreateCard(c *fiber.Ctx) error
	TopUpCard(c *fiber.Ctx) error
}
