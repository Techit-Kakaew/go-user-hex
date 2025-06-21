package usecase

import (
	"errors"
	"time"

	"github.com/Techit-Kakaew/go-user-hex/internal/auth"
	"github.com/Techit-Kakaew/go-user-hex/internal/user/domain"
	"github.com/Techit-Kakaew/go-user-hex/internal/user/repository"
	"github.com/google/uuid"
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

func (uc *userUseCase) Login(email, password string) (string, error) {
	user, err := uc.repo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := auth.GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (uc *userUseCase) GetByID(id string) (*domain.User, error) {
	return uc.repo.GetByID(id)
}
