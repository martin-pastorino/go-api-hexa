package outgoing

import (
	"api/core/domain"
	"context"
)


type ProductRepository interface {
	Save(ctx context.Context, product domain.Product) (string, error)
	GetProduct(ctx context.Context, sku string) (domain.Product, error)
	DeleteProduct(ctx context.Context, sku string) (string, error)
}