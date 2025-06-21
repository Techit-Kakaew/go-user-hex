package repository

import (
	"context"
	"errors"
	"time"

	"github.com/Techit-Kakaew/go-user-hex/internal/user/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepoMongo struct {
	collection *mongo.Collection
}

func NewUserRepoMongo(db *mongo.Database) UserRepository {
	return &userRepoMongo{
		collection: db.Collection("users"),
	}
}

func (r *userRepoMongo) Create(user *domain.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// ตรวจซ้ำก่อนสร้าง
	_, err := r.FindByEmail(user.Email)
	if err == nil {
		return errors.New("user already exists")
	}

	_, err = r.collection.InsertOne(ctx, user)
	return err
}

func (r *userRepoMongo) FindByEmail(email string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user domain.User
	err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func (r *userRepoMongo) GetByID(id string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user domain.User
	err := r.collection.FindOne(ctx, bson.M{"id": id}).Decode(&user)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}
