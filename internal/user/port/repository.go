package port

import "terminal/internal/user/domain"

type UserRepository interface {
	FindByID(id uint) (*domain.User, error)
	CreateCard(card *domain.Card) error
	FindCardByNumber(number string) (*domain.Card, error)
	UpdateCardBalance(card *domain.Card, amount float64) error
	CreateTransaction(tx *domain.CardTransaction) error
}
