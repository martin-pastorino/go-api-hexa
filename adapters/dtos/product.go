package dtos

import (
	"api/core/domain"
	"errors"
	"net/http"
)

type Product struct {
	ID      string `json:"id"`
	Sku     string `json:"sku"`
	Name    string `json:"name"`
	Price   float64 `json:"price"`
}

func (p *Product) Bind(r *http.Request) error	 {

	if p.Sku == "" {
		return errors.New("sku is required")
	}

	if p.Name == "" {
		return errors.New("name is required")
	}

	if p.Price == 0 || p.Price < 0 {
		return errors.New("price is required or must be greater than 0")
	}	

	return nil
}

func (p Product) ToProductDomainModel() domain.Product {
	return domain.Product{
		ID:      p.ID,
		Sku:     p.Sku,
		Name:    p.Name,
		Price:   p.Price,
	}
}

func ToProduct(product domain.Product) Product {
	return Product{
		ID:      product.ID,
		Sku:     product.Sku,
		Name:    product.Name,
		Price:   product.Price,
	}
}
