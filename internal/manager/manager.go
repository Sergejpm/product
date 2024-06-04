package manager

import (
	"context"
	"github.com/jmoiron/sqlx"
	productcommands "github.com/sergejpm/product/internal/commands/product-commands"
	usercommands "github.com/sergejpm/product/internal/commands/user-commands"
	"github.com/sergejpm/product/internal/domain/model"
	"github.com/sergejpm/product/internal/domain/service/authorization"
	productService "github.com/sergejpm/product/internal/domain/service/product"
	"github.com/sergejpm/product/internal/domain/service/registration"
	producthandlers "github.com/sergejpm/product/internal/handlers/product-handlers"
	"github.com/sergejpm/product/internal/handlers/user-handlers"
	"github.com/sergejpm/product/internal/infra/log"
	"github.com/sergejpm/product/internal/infra/repository"
	"github.com/sergejpm/product/pkg/api/product"
)

type Manager struct {
	db       *sqlx.DB
	tokenKey []byte
}

func NewManager(db *sqlx.DB, tokenKey []byte) *Manager {
	return &Manager{db: db, tokenKey: tokenKey}
}

func (m Manager) Registration(ctx context.Context, request *product.RegistrationRequest) error {
	h := userhandlers.NewRegistrationHandler(registration.NewService(repository.NewUserDBRepository(m.db)))
	cmd := &usercommands.RegistrationCommand{
		Username: request.GetUsername(),
		Password: request.GetPassword(),
	}
	err := h.Handle(ctx, cmd)
	if err != nil {
		log.Logger().Errorf("error while handling registration command: %v", err)
	}
	return err
}

func (m Manager) Authorization(ctx context.Context, request *product.AuthorizationRequest) (string, error) {
	h := userhandlers.NewAuthorizationHandler(authorization.NewService(
		repository.NewTokenDBRepository(m.db),
		repository.NewUserDBRepository(m.db),
		m.tokenKey,
	),
	)
	cmd := &usercommands.AuthorizationCommand{
		Username: request.GetUsername(),
		Password: request.GetPassword(),
	}

	result, err := h.Handle(ctx, cmd)
	if err != nil {
		log.Logger().Errorf("error while handling authorization command: %v", err)
	}
	return result, err
}

func (m Manager) ProductInfo(ctx context.Context, request *product.ProductInfoRequest) (*model.Product, error) {
	h := producthandlers.NewInfoHandler(productService.NewService(
		repository.NewProductDBRepository(m.db)),
		authorization.NewService(
			repository.NewTokenDBRepository(m.db),
			repository.NewUserDBRepository(m.db),
			m.tokenKey,
		),
	)
	cmd := &productcommands.InfoCommand{
		Name: request.GetName(),
	}
	result, err := h.Handle(ctx, cmd)
	if err != nil {
		log.Logger().Errorf("error while handling product info command: %v", err)
	}
	return result, err
}
