package usecase

import "github.com/Techit-Kakaew/go-user-hex/internal/user/domain"

type UserUseCase interface {
	Register(user *domain.User) error
}
