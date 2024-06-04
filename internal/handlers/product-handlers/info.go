package producthandlers

import (
	"context"
	productcommands "github.com/sergejpm/product/internal/commands/product-commands"
	"github.com/sergejpm/product/internal/domain/model"
	"github.com/sergejpm/product/internal/domain/service/authorization"
	"github.com/sergejpm/product/internal/domain/service/product"
)

type InfoHandler struct {
	productService *product.Service
	authService    *authorization.Service
}

func NewInfoHandler(service *product.Service, authService *authorization.Service) *InfoHandler {
	return &InfoHandler{productService: service, authService: authService}
}

func (h InfoHandler) Handle(ctx context.Context, cmd *productcommands.InfoCommand) (*model.Product, error) {
	return h.productService.GetProduct(ctx, cmd.Name)
}
