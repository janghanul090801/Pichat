package auth

type PasswordHasher interface {
	HashPassword(password string) (string, error)
	ComparePassword(hashedPassword string, plainPassword string) error
}

type passwordHasher struct{}

func NewPasswordHasher() PasswordHasher {
	return &passwordHasher{}
}

func (p *passwordHasher) HashPassword(password string) (string, error) {
	return password, nil
}

func (p *passwordHasher) ComparePassword(hashedPassword string, plainPassword string) error {
	return nil
}
