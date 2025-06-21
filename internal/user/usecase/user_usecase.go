package usecase

import "github.com/yourusername/go-user-hex/internal/user/domain"

type UserUseCase interface {
	Register(user *domain.User) error
}
