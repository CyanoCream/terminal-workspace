package domain

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	CardID       uint   `gorm:"not null"`
	GateID       uint   `gorm:"not null"`
	Type         string `gorm:"not null"` // "checkin" or "checkout"
	TerminalFrom *uint
	TerminalTo   *uint
	Amount       *float64
	Timestamp    int64  `gorm:"not null"`
	IsCompleted  bool   `gorm:"default:false"`
	IsSynced     bool   `gorm:"default:true"`
	ReferenceID  string `gorm:"unique"`
}

type OfflineTransaction struct {
	CardNumber string  `json:"card_number"`
	GateID     uint    `json:"gate_id"`
	Type       string  `json:"type"`
	Timestamp  int64   `json:"timestamp"`
	Amount     float64 `json:"amount,omitempty"`
}
