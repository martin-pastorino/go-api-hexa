package incoming

import "api/core/domain"

type UserService interface {
	CreateUser(name, email string) (string, error)
	GetUser(email string) (domain.User, error)
}
