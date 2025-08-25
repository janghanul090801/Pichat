package auth

import "Pichat/pkg/entities"

type Service interface {
	GenerateJWToken(user *entities.User) (string, error)
}
