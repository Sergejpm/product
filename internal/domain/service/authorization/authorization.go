package authorization

import (
	"context"
	"database/sql"
	"errors"
	"github.com/sergejpm/product/internal/domain/repository"
	"github.com/sergejpm/product/internal/domain/service/hash"
)

type Service struct {
	tokens   repository.TokenRepository
	users    repository.UserRepository
	tokenKey []byte
	hasher   *hash.Hasher
}

func NewService(tokens repository.TokenRepository, users repository.UserRepository, tokenKey []byte) *Service {
	return &Service{tokens: tokens, users: users, tokenKey: tokenKey, hasher: hash.NewHasher()}
}

func (s Service) GetToken(ctx context.Context, username, password string) (string, error) {
	user, err := s.users.FindUser(ctx, username)
	if err != nil {
		return "", err
	}

	err = s.hasher.CheckUserPassword(user, password)
	if err != nil {
		return "", err
	}

	token, err := s.tokens.GetToken(ctx, user.Id)
	if err == nil {
		return token, nil
	}

	if !errors.Is(err, sql.ErrNoRows) {
		return "", err
	}

	tokenString, err := s.hasher.GenerateToken(s.tokenKey)
	if err != nil {
		return "", err
	}

	err = s.tokens.CreateToken(ctx, user.Id, tokenString)
	if err != nil {
		return "", err
	}

	return tokenString, err
}

func (s Service) Authorize(ctx context.Context, token string) (uint, error) {
	userId, err := s.tokens.FindToken(ctx, token)
	if err == nil {
		return userId, nil
	}

	if errors.Is(err, sql.ErrNoRows) {
		return 0, errors.New("unauthorized")
	}
	return 0, err
}
