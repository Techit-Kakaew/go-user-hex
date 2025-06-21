package repository

import "github.com/yourusername/go-user-hex/internal/user/domain"

type UserRepository interface {
	Create(user *domain.User) error
	FindByEmail(email string) (*domain.User, error)
}
