package port

import "terminal/internal/auth/domain"

type AuthRepository interface {
	FindByUsername(username string) (*domain.User, error)
	Create(user *domain.User) error
}
