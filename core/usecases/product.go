package usecases

import (
	"api/core/domain"
	"api/core/ports/incoming"
	"api/core/ports/outgoing"
	"context"
)


type ProductUserCase struct {
	productRepository outgoing.ProductRepository
}

func NewProductUseCaseProvider(productRepository outgoing.ProductRepository) incoming.ProductService {
	return &ProductUserCase{
		productRepository: productRepository,
	}
}

func NewProductUseCase (productRepository outgoing.ProductRepository) *ProductUserCase {
	return &ProductUserCase{
		productRepository: productRepository,
	}
}

// CreateProduct implements incoming.ProductService.
func (uc *ProductUserCase) CreateProduct(ctx context.Context, product domain.Product) (string, error) {
	id, err := uc.productRepository.Save(ctx, product)
	if err != nil {
		return "", err
	}
	return id, nil
}

// GetProduct implements incoming.ProductService.
func (uc *ProductUserCase) GetProduct(ctx context.Context, sku string) (domain.Product, error) {
	product, err := uc.productRepository.GetProduct(ctx, sku)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

// DeleteProduct implements incoming.ProductService.
func (uc *ProductUserCase) DeleteProduct(ctx context.Context, id string) (string, error) {
	return uc.productRepository.DeleteProduct(ctx, id)
}

