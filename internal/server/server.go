package server

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/sergejpm/product/internal/manager"
	"github.com/sergejpm/product/internal/presentation"
	"github.com/sergejpm/product/pkg/api/product"
)

type ProductServer struct {
	product.UnimplementedProductServer
	manager   *manager.Manager
	presenter *presentation.Presenter
}

func NewServer(db *sqlx.DB, tokenKey []byte) product.ProductServer {
	return &ProductServer{
		manager:   manager.NewManager(db, tokenKey),
		presenter: presentation.NewPresenter(),
	}
}

func (s *ProductServer) Registration(ctx context.Context, req *product.RegistrationRequest) (*product.RegistrationResponse, error) {
	err := s.manager.Registration(ctx, req)
	var success bool
	if err == nil {
		success = true
	}
	response := s.presenter.Registration(ctx, success)

	return response, err
}

func (s *ProductServer) Authorize(ctx context.Context, req *product.AuthorizationRequest) (*product.AuthorizationResponse, error) {
	token, err := s.manager.Authorization(ctx, req)
	if err != nil {
		return nil, err
	}
	return s.presenter.Authorization(ctx, token), nil
}

func (s *ProductServer) ProductInfo(ctx context.Context, req *product.ProductInfoRequest) (*product.ProductInfoResponse, error) {
	productItem, err := s.manager.ProductInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	return s.presenter.ProductInfo(ctx, productItem), nil
}
