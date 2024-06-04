package repository

import (
	"context"
	"github.com/sergejpm/product/internal/domain/model"
)

type UserRepository interface {
	CreateUser(ctx context.Context, username string, password string) error
	FindUser(ctx context.Context, username string) (*model.User, error)
}
