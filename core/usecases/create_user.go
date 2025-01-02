package usecases

import (
	"api/core/domain"
	"api/core/ports/incoming"
	"api/core/ports/outgoing"
	"context"

	"github.com/google/uuid"
)

type CreateUser struct {
	userRepository outgoing.UserRepository
	notifier       outgoing.Notifier
}

// Provider for CreateUser use case
func NewCreateUserProvider(userRepo outgoing.UserRepository, notifier outgoing.Notifier) incoming.UserService {
	return NewCreateUser(userRepo, notifier)
}

func NewCreateUser(userRepository outgoing.UserRepository, notifier outgoing.Notifier) *CreateUser {
	return &CreateUser{
		userRepository: userRepository,
		notifier:       notifier,
	}
}

func (uc *CreateUser) CreateUser(ctx context.Context, name, email string) (string, error) {

	userID := uuid.New().String()

	user := domain.User{
		ID:    userID,
		Name:  name,
		Email: email,
	}

	id, err := uc.userRepository.Save(user)
	if err != nil {
		return "", err
	}

	err = uc.notifier.SendWelcomeEmail(email)
	if err != nil {
		return "", err
	}

	return id, nil
}

// GetUser implements incoming.UserService.
func (uc *CreateUser) GetUser(ctx context.Context, email string) (domain.User, error) {

	user, err := uc.userRepository.GetUser(email)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

// DeleteUser implements incoming.UserService.
func (uc *CreateUser) DeleteUser(ctx context.Context ,email string) error {
	return uc.userRepository.DeleteUser(email)
}
