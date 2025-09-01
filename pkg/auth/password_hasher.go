package auth

import "golang.org/x/crypto/bcrypt"

type PasswordHasher interface {
	HashPassword(password string) (string, error)
	ComparePassword(hashedPassword string, plainPassword string) bool
}

type passwordHasher struct{}

func NewPasswordHasher() PasswordHasher {
	return &passwordHasher{}
}

func (p *passwordHasher) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (p *passwordHasher) ComparePassword(hashedPassword string, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}
