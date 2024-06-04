package hash

import (
	"github.com/golang-jwt/jwt"
	"github.com/sergejpm/product/internal/domain/model"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Hasher struct{}

func NewHasher() *Hasher {
	return &Hasher{}
}

func (h Hasher) CheckUserPassword(user *model.User, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func (h Hasher) GeneratePasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (h Hasher) GenerateToken(key []byte) (string, error) {
	t := jwt.New(jwt.SigningMethodHS256)
	s := time.Now().String() + string(key)
	tokenString, err := t.SignedString([]byte(s))
	return tokenString, err
}
