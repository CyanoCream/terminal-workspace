package port

import (
	"github.com/gofiber/fiber/v2"
	"terminal/internal/terminal/domain"
)

type TerminalService interface {
	CreateTerminal(terminal *domain.Terminal) error
	GetAllTerminals() ([]domain.Terminal, error)
	AddGate(gate *domain.Gate) error
	SetPricing(distance *domain.TerminalDistance) error
	SetTerminalPricing(fromID, toID uint, distance, price float64) error
	GetTerminalPricing(fromID, toID uint) (float64, error)
	FindGateByID(gateID uint) (*domain.Gate, error)
}

type TerminalHandler interface {
	GetAllTerminals(c *fiber.Ctx) error
	CreateTerminal(c *fiber.Ctx) error
	AddGate(c *fiber.Ctx) error
	SetPricing(c *fiber.Ctx) error
}
