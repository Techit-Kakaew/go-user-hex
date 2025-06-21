package repository

import (
	"errors"

	"github.com/yourusername/go-user-hex/internal/user/domain"
)

type userRepoMemory struct {
	users map[string]*domain.User
}

func NewUserRepoMemory() UserRepository {
	return &userRepoMemory{
		users: make(map[string]*domain.User),
	}
}

func (r *userRepoMemory) Create(user *domain.User) error {
	if _, ok := r.users[user.Email]; ok {
		return errors.New("user already exists")
	}
	r.users[user.Email] = user
	return nil
}

func (r *userRepoMemory) FindByEmail(email string) (*domain.User, error) {
	if user, ok := r.users[email]; ok {
		return user, nil
	}
	return nil, errors.New("user not found")
}
