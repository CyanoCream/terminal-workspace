package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"terminal/internal/auth/domain"
	"terminal/internal/auth/port"
	"terminal/pkg/jwt"
)

type authService struct {
	repo      port.AuthRepository
	jwtSecret string
	jwtExpiry int
}

func NewAuthService(repo port.AuthRepository, jwtSecret string, jwtExpiry int) port.AuthService {
	return &authService{
		repo:      repo,
		jwtSecret: jwtSecret,
		jwtExpiry: jwtExpiry,
	}
}

func (s *authService) Login(username, password string) (string, error) {
	user, err := s.repo.FindByUsername(username)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := jwt.GenerateToken(user.ID, user.Role, s.jwtSecret, s.jwtExpiry)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}

func (s *authService) Register(user *domain.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	return s.repo.Create(user)
}
