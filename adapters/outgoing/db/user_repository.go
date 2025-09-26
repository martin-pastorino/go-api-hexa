package db

import (
	mongomodel "api/adapters/outgoing/db/mongo_model"
	"api/core/domain"
	"api/core/errors"
	"api/core/ports/outgoing"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const (
	KEY_USER_CACHE = "user:%s"
	TTL            = 15
)

type UserRepository struct {
	*LocalCache
	collection *mongo.Collection
}

func (r *UserRepository) cacheUser(ctx context.Context, user domain.User) error {
	key := fmt.Sprintf(KEY_USER_CACHE, user.Email)
	result, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return r.cache.Set(ctx, key, result, time.Minute*TTL).Err()
}

func NewUserRepository(cache *LocalCache, db *mongo.Database) *UserRepository {

	return &UserRepository{
		LocalCache: cache,
		collection: db.Collection("users"),
	}
}

// Provider for UserRepository
func NewUserRepositoryProvider(cache *LocalCache, db *mongo.Database) outgoing.UserRepository {
	return NewUserRepository(cache, db)
}

func (r *UserRepository) Save(ctx context.Context, user domain.User) (string, error) {
	// Save user to database
	savedUser, err := r.collection.InsertOne(ctx, mongomodel.ToDomainUserToMongoDBUser(user))
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return "", errors.NewAlreadyExists("user already exists")
		}
	}

	user.ID = savedUser.InsertedID.(bson.ObjectID).Hex()

	err = r.cacheUser(ctx, user)
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
		var userDb mongomodel.UserMongoDB
		filter := bson.D{{Key: "email", Value: email}}
		r.collection.FindOne(context.TODO(), filter).Decode(&userDb)
		user = userDb.ToMongoUserToDomainUser()
		if user.Email == "" {
			return domain.User{}, fmt.Errorf("user not found")
		}
		err := r.cacheUser(ctx, user)
		if err != nil {
			return domain.User{}, err
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

func (r *UserRepository) Search(ctx context.Context, email string) ([]domain.User, error) {
	var users []domain.User
	findOptions := options.Find()
	findOptions.SetLimit(5)

	cursor, err := r.collection.Find(ctx, bson.M{"email": bson.M{"$regex": email, "$options": "i"}}, findOptions)
	if err != nil {
		return users, err
	}

	for cursor.Next(ctx) {
		var user mongomodel.UserMongoDB
		err := cursor.Decode(&user)
		if err != nil {
			return users, err
		}
		users = append(users, user.ToMongoUserToDomainUser())
	}
	return users, nil
}
