package usecases

import (
	"api/core/domain"
	"api/core/ports/incoming"
	"api/core/ports/outgoing"

	"github.com/google/uuid"
)

type CreateUser struct {
	UserRepository outgoing.UserRepository
	Notifier       outgoing.Notifier
}

// Provider for CreateUser use case
func NewCreateUserProvider(userRepo outgoing.UserRepository, notifier outgoing.Notifier) incoming.UserService {
	return NewCreateUser(userRepo, notifier)
}

func NewCreateUser(userRepository outgoing.UserRepository, notifier outgoing.Notifier) *CreateUser {
	return &CreateUser{
		UserRepository: userRepository,
		Notifier:       notifier,
	}
}

func (uc *CreateUser) CreateUser(name, email string) (string, error) {

	userID := uuid.New().String()

	user := domain.User{
		ID:    userID,
		Name:  name,
		Email: email,
	}

	id, err := uc.UserRepository.Save(user)
	if err != nil {
		return "", err
	}

	err = uc.Notifier.SendWelcomeEmail(email)
	if err != nil {
		return "", err
	}

	return id, nil
}

// GetUser implements incoming.UserService.
func (uc *CreateUser) GetUser(email string) (domain.User, error) {

	user, err := uc.UserRepository.GetUser(email)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

// DeleteUser implements incoming.UserService.
func (uc *CreateUser) DeleteUser(email string) error {
	return uc.UserRepository.DeleteUser(email)
}
