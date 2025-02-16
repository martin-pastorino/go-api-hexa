package http

import (
	"api/adapters/dtos"
	core_errors "api/core/errors"
	"api/core/ports/incoming"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

type ProductHandler struct {
	productService incoming.ProductService
}

func NewProductHandler(productService incoming.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

func NewProductHandlerProvider(productService incoming.ProductService) *ProductHandler {
	return NewProductHandler(productService)
}

func (ph *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var productRequest dtos.Product

	if err := render.Bind(r, &productRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	productID, err := ph.productService.CreateProduct(r.Context(), productRequest.ToProductDomainModel())
	if err != nil {
		var alreadyExists *core_errors.AlreadyExists
		if errors.As(err, &alreadyExists) {
			http.Error(w, err.Error(), alreadyExists.Code)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"id": productID})

}

func (ph *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")

	product, err := ph.productService.GetProduct(r.Context(), productId)

	if err != nil || product.ID == "" {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dtos.ToProduct(product))

}

// Validate sku query parameter before use it
func (ph *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {

	productSKU := r.URL.Query().Get("sku")
	if productSKU == "" {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "sku query parameter is required",
		})
		return
	}
	result, err := ph.productService.DeleteProduct(r.Context(), productSKU)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type Result struct {
		Id string `json:"id"`
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Result{Id: result})

}
