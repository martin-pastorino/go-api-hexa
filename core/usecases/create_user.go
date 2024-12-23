package usecases

import (
	"api/core/domain"
	"api/core/ports/outgoing"

	"github.com/google/uuid"
)


type CreateUser struct {
	UserRepository outgoing.UserRepository
	Notifier outgoing.Notifier
}

func NewCreateUser(userRepository outgoing.UserRepository, notifier outgoing.Notifier) *CreateUser {
	return &CreateUser{
		UserRepository: userRepository,
		Notifier: notifier,
	}
}

func (uc *CreateUser) CreateUser(name, email string) (string, error) {


	userID:= uuid.New().String()	

	user := domain.User{
		ID: userID,
		Name: name,
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