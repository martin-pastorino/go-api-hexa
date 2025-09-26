package router

import (
	incomingHTTP "api/adapters/incoming/http"
	"net/http"
	"time"

	localMiddleware "api/infra/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"github.com/go-chi/chi/middleware"
)

func New(userHandler *incomingHTTP.UserHandler, productHandler *incomingHTTP.ProductHandler) chi.Router {
	r := chi.NewRouter()

	// Middleware stack
	r.Use(localMiddleware.RequiredHeaders)
	r.Use(middleware.RequestID)
	r.Use(middleware.Timeout(300 * time.Millisecond))
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(localMiddleware.ErrorHandler)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// Health check endpoint
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, map[string]string{"status": "ok"})
	})

	r.Mount("/users", UsersAPIRouter(userHandler))
	r.Mount("/products", ProductAPIRouter(productHandler))

	return r
}
