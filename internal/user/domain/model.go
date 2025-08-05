package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Role     string `gorm:"not null;default:'user'"`
	Cards    []Card `gorm:"foreignKey:UserID"`
}

type Card struct {
	gorm.Model
	Number    string         `gorm:"unique;not null"`
	Balance   float64        `gorm:"not null;default:0"`
	UserID    uint           `gorm:"not null"`
	User      User           `gorm:"foreignKey:UserID"`
	IsActive  bool           `gorm:"default:true"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type CardTransaction struct {
	gorm.Model
	CardID      uint    `gorm:"not null"`
	Card        Card    `gorm:"foreignKey:CardID"`
	Amount      float64 `gorm:"not null"`
	Type        string  `gorm:"not null"` // "topup", "payment", "refund"
	ReferenceID string  `gorm:"unique"`
	IsSynced    bool    `gorm:"default:false"`
}
