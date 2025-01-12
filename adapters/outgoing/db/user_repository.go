package db

import (
	mongomodel "api/adapters/outgoing/db/mongo_model"
	"api/core/domain"
	"api/core/errors"
	"api/core/ports/outgoing"
	"api/infra/config"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

const (
	KEY_USER_CACHE = "user:%s"
	TTL            = 15
)

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

func (r *UserRepository) Save(ctx context.Context, user domain.User) (string, error) {
	// Save user to database
	key := fmt.Sprintf(KEY_USER_CACHE, user.Email)

	savedUser, err := r.collection.InsertOne(ctx, mongomodel.ToMongoDB(user))
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

	err = r.cache.Set(ctx, key, result, time.Minute*TTL).Err()
	if err != nil {
		return "", err
	}

	fmt.Println("User saved to database")
	return user.ID, nil
}

// GetUser implements outgoing.UserRepository.
func (r *UserRepository) GetUser(ctx context.Context, email string) (domain.User, error) {
	var user domain.User
	result := r.cache.Get(ctx, fmt.Sprintf(KEY_USER_CACHE, email)).Val()

	if result == "" {
		var userDb mongomodel.MongoDB
		filter := bson.D{{Key: "email", Value: email}}
		r.collection.FindOne(context.TODO(), filter).Decode(&userDb)
		user = userDb.ToDomain()
		if user.Email == "" {
			return domain.User{}, fmt.Errorf("user not found")
		}

		return user, nil
	}

	err := json.Unmarshal([]byte(result), &user)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil

}

// DeleteUser implements outgoing.UserRepository.
func (r *UserRepository) DeleteUser(ctx context.Context, email string) (string, error) {
	key := fmt.Sprintf(KEY_USER_CACHE, email)
	err := r.cache.Del(ctx, key).Err()
	if err != nil {
		return "", err
	}

	_, err = r.collection.DeleteOne(ctx, bson.M{"email": email})
	if err != nil {
		return "", err
	}

	return email, nil
}
