package outgoing

import (
	"api/core/domain"
	"context"
)

type UserRepository interface {
	Save(ctx context.Context, user domain.User) (string, error)
	GetUser(ctx context.Context, email string) (domain.User, error)
	DeleteUser(ctx context.Context, email string) (string, error)
	Search(ctx context.Context, email string) ([]domain.User, error)
}
