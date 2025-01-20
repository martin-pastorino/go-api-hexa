package incoming

import (
	"api/core/domain"
	"context"
)

type ProductService interface {
	CreateProduct(ctx context.Context, product domain.Product) (string, error)
	GetProduct(ctx context.Context, sku string) (domain.Product, error)
	DeleteProduct(ctx context.Context, sku string) (string, error)
}