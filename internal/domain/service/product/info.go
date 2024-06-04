package product

import (
	"context"
	"github.com/sergejpm/product/internal/domain/model"
	"github.com/sergejpm/product/internal/domain/repository"
)

type Service struct {
	products repository.ProductRepository
}

func NewService(products repository.ProductRepository) *Service {
	return &Service{products: products}
}

func (s Service) GetProduct(ctx context.Context, name string) (*model.Product, error) {
	return s.products.GetProduct(ctx, name)
}
