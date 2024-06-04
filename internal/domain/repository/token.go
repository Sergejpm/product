package repository

import "context"

type TokenRepository interface {
	CreateToken(ctx context.Context, userId uint, token string) error
	GetToken(ctx context.Context, userId uint) (string, error)
	FindToken(ctx context.Context, token string) (uint, error)
}
