//go:build wireinject
// +build wireinject

package app
import (
	"api/adapters/incoming/http"
	"api/adapters/outgoing/db"
	"api/adapters/outgoing/smtp"
	"api/core/usecases"
	"api/infra/config"
	"github.com/google/wire"

)

func InitializeUsersHandler() *http.UserHandler {
	wire.Build(
		db.NewUserRepositoryProvider,
		smtp.NewNotifierProvider,
		usecases.NewUserUseCaseProvider,
		http.NewUserHandlerProvider,
		config.NewConfigProvider,
	)
	 return &http.UserHandler{}
}