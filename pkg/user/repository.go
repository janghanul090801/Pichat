package user

import (
	"Pichat/api/presenter"
	"Pichat/pkg/ent"
	"Pichat/pkg/entities"
	"context"
)

type Repository interface {
	CreateUser(user *entities.User) (*entities.User, error)
	ReadUser() (*[]presenter.UserResponse, error)
	GetUserById(id int) (*entities.User, error)
	GetUserByEmail(email string) (*entities.User, error)
	UpdateUser(user *entities.User, id int) (*entities.User, error)
	DeleteUser(id int) error
}

type repository struct {
	DBConn  *ent.Client
	Context context.Context
}

func NewRepo(dbconn *ent.Client, ctx context.Context) Repository {
	return &repository{
		DBConn:  dbconn,
		Context: ctx,
	}
}

func (r *repository) CreateUser(user *entities.User) (*entities.User, error) {
	return nil, nil
}

func (r *repository) ReadUser() (*[]presenter.UserResponse, error) {
	return nil, nil
}

func (r *repository) GetUserById(id int) (*entities.User, error) {
	return nil, nil
}

func (r *repository) GetUserByEmail(email string) (*entities.User, error) {
	return nil, nil
}

func (r *repository) UpdateUser(user *entities.User, id int) (*entities.User, error) {
	return nil, nil
}

func (r *repository) DeleteUser(id int) error {
	return nil
}
