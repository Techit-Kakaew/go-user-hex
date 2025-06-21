package usecase

import "github.com/Techit-Kakaew/go-user-hex/internal/user/domain"

type UserUseCase interface {
	Register(user *domain.User) error
	Login(email, password string) (string, error) // return JWT
	GetByID(id string) (*domain.User, error)
}
