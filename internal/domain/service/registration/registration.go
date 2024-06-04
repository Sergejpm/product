package registration

import (
	"context"
	"database/sql"
	"errors"
	"github.com/sergejpm/product/internal/domain/repository"
	"github.com/sergejpm/product/internal/domain/service/hash"
)

type Service struct {
	users  repository.UserRepository
	hasher *hash.Hasher
}

func NewService(users repository.UserRepository) *Service {
	return &Service{users: users, hasher: hash.NewHasher()}
}

func (s Service) CreateUser(ctx context.Context, username, password string) error {
	user, err := s.users.FindUser(ctx, username)
	if user.Id > 0 {
		return errors.New("user already exists")
	} else if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	passwordHash, err := s.hasher.GeneratePasswordHash(password)
	if err != nil {
		return err
	}
	return s.users.CreateUser(ctx, username, passwordHash)
}
