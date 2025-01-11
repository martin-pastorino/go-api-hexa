package db

import (
	"api/core/domain"
	"api/core/errors"
	"api/core/ports/outgoing"
	"api/infra/config"
	"context"
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

const (
	KEY_USER_CACHE = "user:%s"
)

var ctx = context.Background()

type UserRepository struct {
	*LocalCache
	collection *mongo.Collection
}

func NewUserRepository(config *config.Config, cache *LocalCache, db *mongo.Database) *UserRepository {

	return &UserRepository{
		LocalCache: cache,
		collection: db.Collection("users"),
	}
}

// Provider for UserRepository
func NewUserRepositoryProvider(config *config.Config, cache *LocalCache, db *mongo.Database) outgoing.UserRepository {
	return NewUserRepository(config, cache, db)
}

func (r *UserRepository) Save(user domain.User) (string, error) {
	// Save user to database
	key := fmt.Sprintf(KEY_USER_CACHE, user.Email)


	savedUser, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
		return "", errors.NewAlreadyExists("user already exists")
		}
	}
	
	user.ID = savedUser.InsertedID.(bson.ObjectID).Hex()
	result, err := json.Marshal(user)
	if err != nil {
		return "", err
	}

	err = r.cache.Set(ctx, key, result, 0).Err()
	if err != nil {
		return "", err
	}

	fmt.Println("User saved to database")
	return user.ID, nil
}

// GetUser implements outgoing.UserRepository.
func (r *UserRepository) GetUser(email string) (domain.User, error) {
	var user domain.User
	result := r.cache.Get(ctx, fmt.Sprintf(KEY_USER_CACHE, email)).Val()

	if result == "" {
		return domain.User{}, fmt.Errorf("user not found")
	}

	err := json.Unmarshal([]byte(result), &user)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil

}

// DeleteUser implements outgoing.UserRepository.
func (r *UserRepository) DeleteUser(email string) error {
	key := fmt.Sprintf(KEY_USER_CACHE, email)
	err := r.cache.Del(ctx, key).Err()
	if err != nil {
		return err
	}

	fmt.Println("User deleted from database")
	return nil
}
