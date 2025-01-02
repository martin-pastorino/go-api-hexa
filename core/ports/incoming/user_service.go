package incoming

import (
	"api/core/domain"
	"context"
)

type UserService interface {
	CreateUser(ctx context.Context, name, email string) (string, error)
	GetUser(ctx context.Context, email string) (domain.User, error)
	DeleteUser(ctx context.Context, email string) error
}
