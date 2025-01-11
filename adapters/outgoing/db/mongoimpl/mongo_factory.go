package mongoimpl

import (
	"api/infra/config"
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func NewMongoClient(config *config.Config) *mongo.Database {
	client, err := mongo.Connect(options.Client().ApplyURI(config.MongoUrl))
	if err != nil {
		return nil
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil

	}

	db := client.Database(DATABASE_NAME)
	shouldReturn, result := applyIndexes(db)
	if shouldReturn {
		return result
	}

	return client.Database(DATABASE_NAME)
}

func applyIndexes(db *mongo.Database) (bool, *mongo.Database) {
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: FIELD_EMAIL, Value: 1}},
		Options: options.Index().SetUnique(true),
	}
	_, err := db.Collection(COLLECTION_USERS).Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		return true, nil
	}
	return false, nil
}

func NewMongoClientProvider(config *config.Config) *mongo.Database {
	return NewMongoClient(config)
}
