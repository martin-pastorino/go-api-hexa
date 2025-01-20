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
)


const (
	KEY_PRODUCT_CACHE = "product:%s"
	PROCUT_TTL            = 15
)

type ProductRepository struct {
	*LocalCache
	collection *mongo.Collection
}

func NewProductRepository(cache *LocalCache, db *mongo.Database) *ProductRepository {

	return &ProductRepository{
		LocalCache: cache,
		collection: db.Collection("products"),
	}
}

// Provider for UserRepository
func NewProductRepositoryProvider(cache *LocalCache, db *mongo.Database) outgoing.ProductRepository {
	return NewProductRepository(cache, db)
}

func (r *ProductRepository) Save(ctx context.Context, product domain.Product) (string, error) {
	// Save user to database
	key := fmt.Sprintf(KEY_PRODUCT_CACHE, product.Sku)

	savedProduct, err := r.collection.InsertOne(ctx, mongomodel.ToDomainProductToMongoDBProduct(product))
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return "", errors.NewAlreadyExists("product already exists")
		}
	}

	product.ID = savedProduct.InsertedID.(bson.ObjectID).Hex()
	result, err := json.Marshal(product)
	if err != nil {
		return "", err
	}

	err = r.cache.Set(ctx, key, result, time.Minute*PROCUT_TTL).Err()
	if err != nil {
		return "", err
	}

	fmt.Println("prouct saved to database")
	return product.ID, nil
}

// GetProduct implements outgoing.ProductRepository.
func (r *ProductRepository) GetProduct(ctx context.Context, sku string) (domain.Product, error) {
	var product domain.Product
	result := r.cache.Get(ctx, fmt.Sprintf(KEY_PRODUCT_CACHE, sku)).Val()

	if result == "" {
		var productDb mongomodel.ProductMongoDB
		filter := bson.D{{Key: "sku", Value: sku}}
		r.collection.FindOne(context.TODO(), filter).Decode(&productDb)
		product = productDb.ToMongoProductToDomainProduct()
		if product.Sku == "" {
			return domain.Product{}, fmt.Errorf("product not found")
		}

		return product, nil
	}

	err := json.Unmarshal([]byte(result), &product)
	if err != nil {
		return domain.Product{}, err
	}

	return product, nil

}

// DeleteProduct implements outgoing.ProductRepository.
func (r *ProductRepository) DeleteProduct(ctx context.Context, sku string) (string, error) {
	key := fmt.Sprintf(KEY_PRODUCT_CACHE, sku)
	err := r.cache.Del(ctx, key).Err()
	if err != nil {
		return "", err
	}	

	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": sku})
	if err != nil {
		return "", err
	}

	return sku, nil
}

func (r *ProductRepository) Search(ctx context.Context, name string) ([]domain.Product, error) {
	var products []domain.Product
	cursor, err := r.collection.Find(ctx, bson.M{"name": bson.M{"$regex": name, "$options": "i"}})
	if err != nil {
		return products, err
	}

	for cursor.Next(ctx) {
		var product mongomodel.ProductMongoDB
		err := cursor.Decode(&product)
		if err != nil {
			return products, err
		}		
		products = append(products, product.ToMongoProductToDomainProduct())
	}
	return products, nil
}					