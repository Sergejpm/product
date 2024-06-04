package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/sergejpm/product/internal/domain/model"
)

type ProductDBRepository struct {
	db *sqlx.DB
}

func NewProductDBRepository(db *sqlx.DB) *ProductDBRepository {
	return &ProductDBRepository{db: db}
}

func (r ProductDBRepository) GetProduct(ctx context.Context, name string) (*model.Product, error) {
	product := &model.Product{}
	return product, r.db.GetContext(ctx, product, "SELECT * FROM products WHERE name = $1", name)
}
