// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"api/adapters/incoming/http"
	"api/adapters/outgoing/db"
	"api/adapters/outgoing/smtp"
	"api/core/usecases"
)

// Injectors from wire.go:

func InitializeUserHandler() *http.UserHandler {
	userRepository := db.NewUserRepositoryProvider()
	notifier := smtp.NewNotifierProvider()
	userService := usecases.NewCreateUserProvider(userRepository, notifier)
	userHandler := http.NewUserHandlerProvider(userService)
	return userHandler
}