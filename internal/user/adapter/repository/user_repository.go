package repository

import (
	"gorm.io/gorm"
	"terminal/internal/user/domain"
	"terminal/internal/user/port"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) port.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindByID(id uint) (*domain.User, error) {
	var user domain.User
	err := r.db.Preload("Cards").First(&user, id).Error
	return &user, err
}

func (r *userRepository) CreateCard(card *domain.Card) error {
	return r.db.Create(card).Error
}

func (r *userRepository) FindCardByNumber(number string) (*domain.Card, error) {
	var card domain.Card
	err := r.db.Preload("User").Where("number = ?", number).First(&card).Error
	return &card, err
}

func (r *userRepository) UpdateCardBalance(card *domain.Card, amount float64) error {
	card.Balance += amount
	return r.db.Save(card).Error
}

func (r *userRepository) CreateTransaction(tx *domain.CardTransaction) error {
	return r.db.Create(tx).Error
}
