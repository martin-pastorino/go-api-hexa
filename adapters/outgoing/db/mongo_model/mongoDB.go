package mongomodel

import (
	"api/core/domain"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type MongoDB struct {
	ID      bson.ObjectID `bson:"_id, omitempty"`
	Name    string        `bson:"name"`
	Phone   string        `bson:"phone"`
	Address string        `bson:"address"`
	Email   string        `bson:"email"`
}

func (m MongoDB) ToDomain() domain.User {
	return domain.User{
		ID:      m.ID.Hex(),
		Name:    m.Name,
		Phone:   m.Phone,
		Address: m.Address,
		Email:   m.Email,
	}
}

func ToMongoDB(user domain.User) MongoDB {
	return MongoDB{
		ID:      bson.NewObjectID(),
		Name:    user.Name,
		Phone:   user.Phone,
		Address: user.Address,
		Email:   user.Email,
	}
}
