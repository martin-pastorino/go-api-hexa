package mongomodel

import (
	"api/core/domain"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type UserMongoDB struct {
	ID      bson.ObjectID `bson:"_id, omitempty"`
	Name    string        `bson:"name"`
	Phone   string        `bson:"phone"`
	Address string        `bson:"address"`
	Email   string        `bson:"email"`
}

func (m UserMongoDB) ToMongoUserToDomainUser() domain.User {
	return domain.User{
		ID:      m.ID.Hex(),
		Name:    m.Name,
		Phone:   m.Phone,
		Address: m.Address,
		Email:   m.Email,
	}
}

func ToDomainUserToMongoDBUser(user domain.User) UserMongoDB {
	return UserMongoDB{
		ID:      bson.NewObjectID(),
		Name:    user.Name,
		Phone:   user.Phone,
		Address: user.Address,
		Email:   user.Email,
	}
}

type ProductMongoDB struct {
	ID      bson.ObjectID `bson:"_id, omitempty"`
	Sku     string        `bson:"sku"`
	Name    string        `bson:"name"`
	Price   float64       `bson:"price"`
}

func (m ProductMongoDB) ToMongoProductToDomainProduct() domain.Product {
	return domain.Product{
		ID:      m.ID.Hex(),
		Sku:     m.Sku,
		Name:    m.Name,
		Price:   m.Price,
	}
}

func ToDomainProductToMongoDBProduct(product domain.Product) ProductMongoDB {
	return ProductMongoDB{
		ID:      bson.NewObjectID(),
		Sku:     product.Sku,
		Name:    product.Name,
		Price:   product.Price,
	}
}
