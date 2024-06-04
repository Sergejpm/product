package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/sergejpm/product/internal/domain/model"
)

type TokenDBRepository struct {
	db *sqlx.DB
}

func NewTokenDBRepository(db *sqlx.DB) *TokenDBRepository {
	return &TokenDBRepository{db: db}
}

func (r TokenDBRepository) GetToken(ctx context.Context, userId uint) (string, error) {
	token := &model.Token{}
	err := r.db.GetContext(ctx, &token.Token, "SELECT token FROM tokens WHERE user_id = $1 AND expired_at > now()", userId)
	return token.Token, err
}

func (r TokenDBRepository) CreateToken(ctx context.Context, userId uint, token string) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO tokens (token,user_id) VALUES ($1,$2)", token, userId)
	return err
}

func (r TokenDBRepository) FindToken(ctx context.Context, token string) (uint, error) {
	var userId uint
	err := r.db.GetContext(ctx, &userId, "SELECT user_id FROM tokens WHERE token = $1 AND expired_at > now()", token)
	return userId, err
}
