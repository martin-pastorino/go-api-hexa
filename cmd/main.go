package main

import (
	"api/cmd/app"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

var port int16 = 8080
func main() {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	 userHandler := app.InitializeUserHandler()	

	r.Post("/users", userHandler.CreateUser)
	r.Get("/users", userHandler.GetUser)

	http.ListenAndServe( fmt.Sprintf(":%d",port), r)
}
