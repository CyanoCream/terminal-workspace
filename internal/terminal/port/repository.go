package port

import "terminal/internal/terminal/domain"

type TerminalRepository interface {
	CreateTerminal(terminal *domain.Terminal) error
	GetAllTerminals() ([]domain.Terminal, error)
	AddGate(gate *domain.Gate) error
	SetPricing(distance *domain.TerminalDistance) error
	GetPricing(fromID, toID uint) (*domain.TerminalDistance, error)
	FindGateByID(id uint) (*domain.Gate, error)
}
