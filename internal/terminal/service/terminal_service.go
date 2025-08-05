package service

import (
	"terminal/internal/terminal/domain"
	"terminal/internal/terminal/port"
)

type terminalService struct {
	repo port.TerminalRepository
}

func NewTerminalService(repo port.TerminalRepository) port.TerminalService {
	return &terminalService{repo: repo}
}
func (s *terminalService) SetPricing(distance *domain.TerminalDistance) error {
	return s.repo.SetPricing(distance)
}
func (s *terminalService) CreateTerminal(terminal *domain.Terminal) error {
	return s.repo.CreateTerminal(terminal)
}

func (s *terminalService) GetAllTerminals() ([]domain.Terminal, error) {
	return s.repo.GetAllTerminals()
}

func (s *terminalService) AddGate(gate *domain.Gate) error {
	return s.repo.AddGate(gate)
}

func (s *terminalService) SetTerminalPricing(fromID, toID uint, distance, price float64) error {
	pricing := &domain.TerminalDistance{
		FromTerminalID: fromID,
		ToTerminalID:   toID,
		Distance:       distance,
		BasePrice:      price,
	}
	return s.repo.SetPricing(pricing)
}

func (s *terminalService) GetTerminalPricing(fromID, toID uint) (float64, error) {
	pricing, err := s.repo.GetPricing(fromID, toID)
	if err != nil {
		return 0, err
	}
	return pricing.BasePrice, nil
}

func (s *terminalService) FindGateByID(gateID uint) (*domain.Gate, error) {
	return s.repo.FindGateByID(gateID)
}
