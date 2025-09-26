package http

import (
	"api/adapters/dtos"
	"api/core/domain"
	"api/core/errors"
	"api/mocks"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestProductHandler_CreateProduct(t *testing.T) {
	t.Run("should return 201 when product is created", func(t *testing.T) {
		productService := new(mocks.ProductService)
		productHandler := NewProductHandler(productService)

		productDto := dtos.Product{
			Sku:   "any_sku",
			Name:  "any_name",
			Price: 10,
		}

		        productService.On("CreateProduct", mock.Anything, mock.Anything).Return("any_id", nil)
		
		        body, _ := json.Marshal(productDto)
		
		        req := httptest.NewRequest("POST", "/products", bytes.NewReader(body))
		        req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

		productHandler.CreateProduct(rr, req)

		assert.Equal(t, http.StatusCreated, rr.Code)
	})

	t.Run("should return 400 when body is invalid", func(t *testing.T) {
		productService := new(mocks.ProductService)
		productHandler := NewProductHandler(productService)

		req := httptest.NewRequest("POST", "/products", bytes.NewReader([]byte("invalid")))
		rr := httptest.NewRecorder()

		productHandler.CreateProduct(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should return 409 when product already exists", func(t *testing.T) {
		productService := new(mocks.ProductService)
		productHandler := NewProductHandler(productService)

		productDto := dtos.Product{
			Sku:   "any_sku",
			Name:  "any_name",
			Price: 10,
		}

		productService.On("CreateProduct", mock.Anything, mock.Anything).Return("", &errors.AlreadyExists{Code: http.StatusConflict, Message: "Product already exists"})

		body, _ := json.Marshal(productDto)

		req := httptest.NewRequest("POST", "/products", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()

		productHandler.CreateProduct(rr, req)

		assert.Equal(t, http.StatusConflict, rr.Code)
	})

	t.Run("should return 500 when something goes wrong", func(t *testing.T) {
		productService := new(mocks.ProductService)
		productHandler := NewProductHandler(productService)

		productDto := dtos.Product{
			Sku:   "any_sku",
			Name:  "any_name",
			Price: 10,
		}

		productService.On("CreateProduct", mock.Anything, mock.Anything).Return("", assert.AnError)

		body, _ := json.Marshal(productDto)

		req := httptest.NewRequest("POST", "/products", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()

		productHandler.CreateProduct(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})
}

func TestProductHandler_GetProduct(t *testing.T) {
	t.Run("should return 200 when product is found", func(t *testing.T) {
		productService := new(mocks.ProductService)
		productHandler := NewProductHandler(productService)

		product := domain.Product{
			ID:    "any_id",
			Sku:   "any_sku",
			Name:  "any_name",
			Price: 10,
		}

		productService.On("GetProduct", mock.Anything, "any_id").Return(product, nil)

		req := httptest.NewRequest("GET", "/products?id=any_id", nil)
		rr := httptest.NewRecorder()

		productHandler.GetProduct(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("should return 404 when product is not found", func(t *testing.T) {
		productService := new(mocks.ProductService)
		productHandler := NewProductHandler(productService)

		productService.On("GetProduct", mock.Anything, "any_id").Return(domain.Product{}, assert.AnError)

		req := httptest.NewRequest("GET", "/products?id=any_id", nil)
		rr := httptest.NewRecorder()

		productHandler.GetProduct(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
	})
}

func TestProductHandler_DeleteProduct(t *testing.T) {
	r := chi.NewRouter()
	productService := new(mocks.ProductService)
	productHandler := NewProductHandler(productService)
	r.Delete("/products", productHandler.DeleteProduct)

	t.Run("should return 200 when product is deleted", func(t *testing.T) {
		productService.On("DeleteProduct", mock.Anything, "any_sku").Return("any_id", nil).Once()

		req := httptest.NewRequest("DELETE", "/products?sku=any_sku", nil)
		rr := httptest.NewRecorder()

		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("should return 400 when sku is not provided", func(t *testing.T) {
		req := httptest.NewRequest("DELETE", "/products", nil)
		rr := httptest.NewRecorder()

		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should return 500 when something goes wrong", func(t *testing.T) {
		productService.On("DeleteProduct", mock.Anything, "any_sku").Return("", assert.AnError).Once()

		req := httptest.NewRequest("DELETE", "/products?sku=any_sku", nil)
		rr := httptest.NewRecorder()

		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})
}
