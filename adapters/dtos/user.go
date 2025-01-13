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

func  ToUser(user domain.User) User {
	return User {
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Phone: user.Phone,
		Address: user.Address,
	}
}

func  ToUsers(user []domain.User) []User {
	var users []User
	for _, u := range user {
		users = append(users, ToUser(u))
		
	}
	return users
}
	
