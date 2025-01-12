package main

import (
	"api/cmd/app"
	"api/infra/router"
	"fmt"
	"net/http"

	localMiddleware "api/infra/middleware"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

var port int16 = 8080

func main() {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(localMiddleware.Recovery)
	r.Use(localMiddleware.RequiredHeaders)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Mount("/users", router.UsersAPIRouter(app.InitializeUsersHandler()))

	fmt.Printf("Server running on port %d\n", port)
	fmt.Println((http.ListenAndServe(fmt.Sprintf(":%d", port), r)))
}
