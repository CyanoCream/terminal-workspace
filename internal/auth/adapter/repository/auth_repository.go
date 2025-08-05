package repository

import (
	"gorm.io/gorm"
	"terminal/internal/auth/domain"
	"terminal/internal/auth/port"
)

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) port.AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) FindByUsername(username string) (*domain.User, error) {
	var user domain.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *authRepository) Create(user *domain.User) error {
	return r.db.Create(user).Error
}
