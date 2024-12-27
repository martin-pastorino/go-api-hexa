package db

import (
	"api/core/domain"
	"api/core/ports/outgoing"
	"api/infra/config"
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
)

const (
	// RedisHost is the host of the Redis server.
	RedisHost = "localhost:6379"
	KEY_USER_CACHE = "user:%s"
)

var ctx = context.Background()

type UserRepository struct {
	redisCLient *redis.Client
}

func NewUserRepository(config  *config.Config) *UserRepository {
	fmt.Println(config)
	return &UserRepository{
		redisCLient: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", config.RedisHost , config.RedisPort),
			Password: config.RedisPassword, // no password set
			Username: "default", // use default username
			DB:       0,  // use default DB
		}),
	}
}

// Provider for UserRepository
func NewUserRepositoryProvider(config *config.Config) outgoing.UserRepository {
	return NewUserRepository(config)
}

func (r *UserRepository) Save(user domain.User) (string, error) {
	// Save user to database
	key := fmt.Sprintf(KEY_USER_CACHE, user.Email)

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
	return user.ID, nil
}

// GetUser implements outgoing.UserRepository.
func (r *UserRepository) GetUser(email string) (domain.User, error) {
	var user domain.User
	result := r.redisCLient.Get(ctx, fmt.Sprintf(KEY_USER_CACHE, email)).Val()

	if result == "" {
		return domain.User{}, fmt.Errorf("user not found")
	}

	err := json.Unmarshal([]byte(result),&user)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
	
}

// DeleteUser implements outgoing.UserRepository.
func (r *UserRepository) DeleteUser(email string) error {
	key := fmt.Sprintf(KEY_USER_CACHE, email)
	err := r.redisCLient.Del(ctx, key).Err()
	if err != nil {
		return err
	}
	
	fmt.Println("User deleted from database")
	return nil
}
