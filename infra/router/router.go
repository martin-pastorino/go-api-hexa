package router

import (
	"api/adapters/incoming/http"
	"api/infra/middleware"

	"github.com/go-chi/chi/v5"
)

func New(userHandler *http.UserHandler, productHandler *http.ProductHandler) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.ErrorHandler)

	r.Mount("/users", UsersAPIRouter(userHandler))
	r.Mount("/products", ProductAPIRouter(productHandler))

	return r
}
