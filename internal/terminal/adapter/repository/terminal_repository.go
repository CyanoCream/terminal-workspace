package repository

import (
	"gorm.io/gorm"
	"terminal/internal/terminal/domain"
	"terminal/internal/terminal/port"
)

type terminalRepository struct {
	db *gorm.DB
}

func NewTerminalRepository(db *gorm.DB) port.TerminalRepository {
	return &terminalRepository{db: db}
}

func (r *terminalRepository) CreateTerminal(terminal *domain.Terminal) error {
	return r.db.Create(terminal).Error
}

func (r *terminalRepository) GetAllTerminals() ([]domain.Terminal, error) {
	var terminals []domain.Terminal
	err := r.db.Preload("Gates").Find(&terminals).Error
	return terminals, err
}

func (r *terminalRepository) AddGate(gate *domain.Gate) error {
	return r.db.Create(gate).Error
}

func (r *terminalRepository) SetPricing(distance *domain.TerminalDistance) error {

	var existing domain.TerminalDistance
	err := r.db.Where("from_terminal_id = ? AND to_terminal_id = ?",
		distance.FromTerminalID, distance.ToTerminalID).First(&existing).Error

	if err == nil {
		// Update existing pricing
		return r.db.Model(&existing).Updates(distance).Error
	}

	// Create new pricing
	return r.db.Create(distance).Error
}

func (r *terminalRepository) GetPricing(fromID, toID uint) (*domain.TerminalDistance, error) {
	var pricing domain.TerminalDistance
	err := r.db.Where("from_terminal_id = ? AND to_terminal_id = ?", fromID, toID).First(&pricing).Error
	return &pricing, err
}

func (r *terminalRepository) FindGateByID(id uint) (*domain.Gate, error) {
	var gate domain.Gate
	err := r.db.Preload("Terminal").First(&gate, id).Error
	return &gate, err
}
