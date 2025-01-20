package dtos

import (
	"api/core/domain"
	"errors"
	"net/http"
)

type User struct {
	ID      string `json:"id,omitempty"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

func (u *User) Bind(r *http.Request) error {
	if u.Address == "" {
		return errors.New("address is required")
	}

	if u.Email == "" {
		return errors.New("email is required")
	}

	if u.Name == "" {
		return errors.New("name is required")
	}

	if u.Phone == "" {
		return errors.New("phone is required")
	}

	return nil
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

func ToUser(user domain.User) User {
	return User{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		Phone:   user.Phone,
		Address: user.Address,
	}
}

func ToUsers(user []domain.User) []User {
	var users []User
	for _, u := range user {
		users = append(users, ToUser(u))

	}
	return users
}
