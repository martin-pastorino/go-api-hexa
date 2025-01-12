package mongomodel

import "api/core/domain"

type MongoDB struct {
	ID      string `bson:"_id"`
	Name    string `bson:"name"`
	Phone   string `bson:"phone"`
	Address string `bson:"address"`
	Email   string `bson:"email"`
}

func (m MongoDB) ToDomain() domain.User {
	return domain.User{
		ID:      m.ID,
		Name:    m.Name,
		Phone:   m.Phone,
		Address: m.Address,
		Email:   m.Email,
	}
}
