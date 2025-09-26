package usecases

import (
	"api/core/domain"
	"api/core/ports/incoming"
	"api/core/ports/outgoing"
	"context"
)

type UserUseCase struct {
	userRepository outgoing.UserRepository
	notifier       outgoing.Notifier
}

// Provider for CreateUser use case
func NewUserUseCaseProvider(userRepo outgoing.UserRepository, notifier outgoing.Notifier) incoming.UserService {
	return NewUserUseCase(userRepo, notifier)
}

func NewUserUseCase(userRepository outgoing.UserRepository, notifier outgoing.Notifier) *UserUseCase {
	return &UserUseCase{
		userRepository: userRepository,
		notifier:       notifier,
	}
}

// // CreateUser implements incoming.UserService.
func (uc *UserUseCase) CreateUser(ctx context.Context, user domain.User) (string, error) {

	id, err := uc.userRepository.Save(ctx, user)
	if err != nil {
		return "", err
	}

	err = uc.notifier.SendWelcomeEmail(ctx, user.Email)
	if err != nil {
		return "", err
	}

	return id, nil
}

// GetUser implements incoming.UserService.
func (uc *UserUseCase) GetUser(ctx context.Context, email string) (domain.User, error) {

	user, err := uc.userRepository.GetUser(ctx, email)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

// DeleteUser implements incoming.UserService.
func (uc *UserUseCase) DeleteUser(ctx context.Context, email string) (string, error) {
	return uc.userRepository.DeleteUser(ctx, email)
}

func (uc *UserUseCase) Search(ctx context.Context, email string) ([]domain.User, error) {
	return uc.userRepository.Search(ctx, email)
}
