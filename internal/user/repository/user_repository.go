package repository

import "github.com/Techit-Kakaew/go-user-hex/internal/user/domain"

type UserRepository interface {
	Create(user *domain.User) error
	FindByEmail(email string) (*domain.User, error)
	GetByID(id string) (*domain.User, error)
}
