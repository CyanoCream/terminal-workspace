package domain

import "gorm.io/gorm"

type Terminal struct {
	gorm.Model
	Name      string         `gorm:"unique;not null"`
	Address   string         `gorm:"not null"`
	Latitude  float64        `gorm:"not null"`
	Longitude float64        `gorm:"not null"`
	Gates     []Gate         `gorm:"foreignKey:TerminalID"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Gate struct {
	gorm.Model
	TerminalID uint           `gorm:"not null"`
	Terminal   Terminal       `gorm:"foreignKey:TerminalID"`
	Name       string         `gorm:"not null"`
	IsActive   bool           `gorm:"default:true"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

type TerminalDistance struct {
	gorm.Model
	FromTerminalID uint     `gorm:"not null"`
	FromTerminal   Terminal `gorm:"foreignKey:FromTerminalID"`
	ToTerminalID   uint     `gorm:"not null"`
	ToTerminal     Terminal `gorm:"foreignKey:ToTerminalID"`
	Distance       float64  `gorm:"not null"` // in km
	BasePrice      float64  `gorm:"not null"` // base fare
}
