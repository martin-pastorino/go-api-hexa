package db

import (
	"api/core/domain"
	"api/core/ports/outgoing"
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type UserRepository struct {
	redisCLient *redis.Client
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		redisCLient: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		}),
	}
}

// Provider for UserRepository
func NewUserRepositoryProvider() outgoing.UserRepository {
	return NewUserRepository()
}

func (r *UserRepository) Save(user domain.User) (string, error) {
	// Save user to database
	key := fmt.Sprintf("user:%s", user.Email)

	if exists := r.redisCLient.Exists(ctx, key); exists.Val() == 1 {
		return "", fmt.Errorf("user already exists")
	}

	result, err := json.Marshal(user)
	if err != nil {
		return "", err
	}

	err = r.redisCLient.Set(ctx, key, result, 0).Err()
	if err != nil {
		return "", err
	}

	fmt.Println("User saved to database")
	return "user_id", nil
}