package presentation

import (
	"context"
	"github.com/sergejpm/product/internal/domain/model"
	"github.com/sergejpm/product/pkg/api/product"
)

type Presenter struct{}

func NewPresenter() *Presenter {
	return &Presenter{}
}

func (p Presenter) Registration(ctx context.Context, success bool) *product.RegistrationResponse {
	return &product.RegistrationResponse{Success: success}
}

func (p Presenter) Authorization(ctx context.Context, token string) *product.AuthorizationResponse {
	return &product.AuthorizationResponse{
		Token: token,
	}
}

func (p Presenter) ProductInfo(ctx context.Context, productItem *model.Product) *product.ProductInfoResponse {
	return &product.ProductInfoResponse{
		Id:          productItem.Id,
		Name:        productItem.Name,
		Description: productItem.Description,
		Price:       productItem.Price.String(),
	}
}
