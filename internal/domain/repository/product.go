package repository

import (
	"context"
	"github.com/sergejpm/product/internal/domain/model"
)

type ProductRepository interface {
	GetProduct(ctx context.Context, name string) (*model.Product, error)
}
