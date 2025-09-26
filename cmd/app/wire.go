//go:build wireinject
// +build wireinject

package app
import (
	"api/adapters/incoming/http"
	"api/adapters/outgoing/db"
	"api/adapters/outgoing/db/mongoimpl"
	"api/adapters/outgoing/smtp"
	"api/core/usecases"
	"api/infra/config"
	"github.com/google/wire"

)

func InitializeUsersHandler() *http.UserHandler {
	wire.Build(
		db.NewUserRepositoryProvider,
		db.NewCacheProvider,
		mongoimpl.NewMongoClientProvider,
		smtp.NewNotifierProvider,
		usecases.NewUserUseCaseProvider,
		http.NewUserHandlerProvider,
		config.NewConfigProvider,

		
	)
	 return &http.UserHandler{}
}

func InitializeProductsHandler() *http.ProductHandler {
	wire.Build(
		db.NewProductRepositoryProvider,
		db.NewCacheProvider,
		mongoimpl.NewMongoClientProvider,
		usecases.NewProductUseCaseProvider,
		http.NewProductHandlerProvider,
		config.NewConfigProvider,

		
	)
	 return &http.ProductHandler{}
}