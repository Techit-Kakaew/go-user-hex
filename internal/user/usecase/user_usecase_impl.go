package usecase

import (
	"time"

	"github.com/google/uuid"
	"github.com/yourusername/go-user-hex/internal/user/domain"
	"github.com/yourusername/go-user-hex/internal/user/repository"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	repo repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{repo: repo}
}

func (uc *userUseCase) Register(user *domain.User) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.ID = uuid.New().String()
	user.Password = string(hashed)
	user.CreatedAt = time.Now()

	return uc.repo.Create(user)
}
