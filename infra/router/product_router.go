package router

import (
	"api/adapters/incoming/http"

	"github.com/go-chi/chi/v5"
)



func ProductAPIRouter(productHandler *http.ProductHandler) chi.Router {
	r := chi.NewRouter()
	r.Post("/", productHandler.CreateProduct)
	r.Get("/", productHandler.GetProduct)
	r.Delete("/", productHandler.DeleteProduct)
	return r
}