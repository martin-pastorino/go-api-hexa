package dtos

import "api/core/domain"

type User struct {
	ID      string `json:"id,omitempty"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

func (u User) ToUserDomainModel() domain.User {
	return domain.User{	
		ID:      u.ID,
		Name:    u.Name,
		Email:   u.Email,
		Phone:   u.Phone,
		Address: u.Address,
	}
}
