package auth

import (
	"Pichat/pkg/users"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type Service interface {
	GenerateJWToken(email string) (string, error)
	CheckPassword(password string, hash string) bool
	HashPassword(password string) (string, error)
}

type service struct {
	userRepository users.Repository
	passwordHasher PasswordHasher
}

func (s *service) GenerateJWToken(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return t, nil
}

func (s *service) CheckPassword(password string, hash string) bool {
	return s.passwordHasher.ComparePassword(hash, password)
}

func (s *service) HashPassword(password string) (string, error) {
	return s.passwordHasher.HashPassword(password)
}

func NewService(userRepo users.Repository, passwordHasher PasswordHasher) Service {
	return &service{
		userRepository: userRepo,
		passwordHasher: passwordHasher,
	}
}
