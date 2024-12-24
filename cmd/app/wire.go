//go:build wireinject
// +build wireinject

package app

import (
	"api/adapters/incoming/http"
	"api/adapters/outgoing/db"
	"api/adapters/outgoing/smtp"
	"api/core/usecases"
	"github.com/google/wire"

)

func InitializeUserHandler() *http.UserHandler {
	wire.Build(
		db.NewUserRepositoryProvider,
		smtp.NewNotifierProvider,
		usecases.NewCreateUserProvider,
		http.NewUserHandlerProvider,
	)
	 return &http.UserHandler{}
}