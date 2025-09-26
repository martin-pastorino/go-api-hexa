package dtos

import (
	"api/core/domain"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProduct_Bind(t *testing.T) {
	t.Run("should return error when sku is empty", func(t *testing.T) {
		p := Product{
			Sku: "",
		}
		err := p.Bind(&http.Request{})
		assert.NotNil(t, err)
		assert.Equal(t, "sku is required", err.Error())
	})

	t.Run("should return error when name is empty", func(t *testing.T) {
		p := Product{
			Sku:  "any_sku",
			Name: "",
		}
		err := p.Bind(&http.Request{})
		assert.NotNil(t, err)
		assert.Equal(t, "name is required", err.Error())
	})

	t.Run("should return error when price is 0", func(t *testing.T) {
		p := Product{
			Sku:   "any_sku",
			Name:  "any_name",
			Price: 0,
		}
		err := p.Bind(&http.Request{})
		assert.NotNil(t, err)
		assert.Equal(t, "price is required or must be greater than 0", err.Error())
	})

	t.Run("should return error when price is less than 0", func(t *testing.T) {
		p := Product{
			Sku:   "any_sku",
			Name:  "any_name",
			Price: -1,
		}
		err := p.Bind(&http.Request{})
		assert.NotNil(t, err)
		assert.Equal(t, "price is required or must be greater than 0", err.Error())
	})

	t.Run("should return nil when all fields are valid", func(t *testing.T) {
		p := Product{
			Sku:   "any_sku",
			Name:  "any_name",
			Price: 1,
		}
		err := p.Bind(&http.Request{})
		assert.Nil(t, err)
	})
}

func TestProduct_ToProductDomainModel(t *testing.T) {
	t.Run("should convert Product to domain.Product", func(t *testing.T) {
		p := Product{
			ID:    "any_id",
			Sku:   "any_sku",
			Name:  "any_name",
			Price: 1,
		}
		domainProduct := p.ToProductDomainModel()
		assert.Equal(t, p.ID, domainProduct.ID)
		assert.Equal(t, p.Sku, domainProduct.Sku)
		assert.Equal(t, p.Name, domainProduct.Name)
		assert.Equal(t, p.Price, domainProduct.Price)
	})
}

func TestToProduct(t *testing.T) {
	t.Run("should convert domain.Product to Product", func(t *testing.T) {
		domainProduct := domain.Product{
			ID:    "any_id",
			Sku:   "any_sku",
			Name:  "any_name",
			Price: 1,
		}
		p := ToProduct(domainProduct)
		assert.Equal(t, domainProduct.ID, p.ID)
		assert.Equal(t, domainProduct.Sku, p.Sku)
		assert.Equal(t, domainProduct.Name, p.Name)
		assert.Equal(t, domainProduct.Price, p.Price)
	})
}
