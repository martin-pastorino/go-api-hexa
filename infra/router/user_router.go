package router

import (
	"api/adapters/incoming/http"

	"github.com/go-chi/chi/v5"
)

func UsersAPIRouter(userHandler *http.UserHandler) chi.Router {
	r := chi.NewRouter()
	r.Post("/", userHandler.CreateUser)
	r.Get("/", userHandler.GetUser)
	r.Delete("/", userHandler.DeleteUser)
	return r
}
