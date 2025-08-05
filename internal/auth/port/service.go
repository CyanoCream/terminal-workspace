package port

import (
	"github.com/gofiber/fiber/v2"
	"terminal/internal/auth/domain"
)

type AuthService interface {
	Login(username, password string) (string, error)
	Register(user *domain.User) error
}

type AuthHandler interface {
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
}
