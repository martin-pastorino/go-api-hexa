package http

import (
	"api/adapters/dtos"
	core_errors "api/core/errors"
	"api/core/ports/incoming"
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
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	productID, err := ph.productService.CreateProduct(r.Context(), productRequest.ToProductDomainModel())
	if err != nil {
		var alreadyExists *core_errors.AlreadyExists
		if errors.As(err, &alreadyExists) {
			render.Render(w, r, ErrAlreadyExists(err))
			return
		}
		render.Render(w, r, ErrInternalServer(err))
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, map[string]string{"id": productID})
}

func (ph *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")

	product, err := ph.productService.GetProduct(r.Context(), productId)

	if err != nil {
		render.Render(w, r, ErrInternalServer(err))
		return
	}

	if product.ID == "" {
		render.Render(w, r, ErrNotFound(errors.New("product not found")))
		return
	}

	render.JSON(w, r, dtos.ToProduct(product))

}

// Validate sku query parameter before use it
func (ph *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {

	productSKU := r.URL.Query().Get("sku")
	if productSKU == "" {
		render.Render(w, r, ErrInvalidRequest(errors.New("sku query parameter is required")))
		return
	}
	result, err := ph.productService.DeleteProduct(r.Context(), productSKU)
	if err != nil {
		render.Render(w, r, ErrInternalServer(err))
		return
	}

	type Result struct {
		Id string `json:"id"`
	}

	render.JSON(w, r, Result{Id: result})

}
