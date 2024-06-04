package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/sergejpm/product/internal/domain/model"
	"strings"
)

type UserDBRepository struct {
	db *sqlx.DB
}

func NewUserDBRepository(db *sqlx.DB) *UserDBRepository {
	return &UserDBRepository{db: db}
}

func (r UserDBRepository) CreateUser(ctx context.Context, username string, passwordHash string) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO users(username,password) VALUES ($1,$2)", strings.ToLower(username), passwordHash)
	return err
}

func (r UserDBRepository) FindUser(ctx context.Context, username string) (*model.User, error) {
	user := &model.User{}
	err := r.db.GetContext(ctx, user, "SELECT id,username,password FROM users WHERE username = $1", strings.ToLower(username))
	return user, err
}
