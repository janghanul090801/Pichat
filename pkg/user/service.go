package user

import "Pichat/pkg/entities"

type Service interface {
	InsertUser(user *entities.User) (*entities.User, error)
	ReadAllUsers() ([]*entities.User, error)
	ReadUser(id int) (*entities.User, error)
	UpdateUser(user *entities.User) (*entities.User, error)
	RemoveUser(id int) error
}

type service struct {
	repo Repository
}

func (s service) InsertUser(user *entities.User) (*entities.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) ReadAllUsers() ([]*entities.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) ReadUser(id int) (*entities.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) UpdateUser(user *entities.User) (*entities.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) RemoveUser(id int) error {
	//TODO implement me
	panic("implement me")
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}
