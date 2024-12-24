package outgoing

import "api/core/domain"

type UserRepository interface {
	Save(user domain.User) (string, error)	
	GetUser(email string) (domain.User, error)
}
