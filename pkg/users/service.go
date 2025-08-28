package users

import (
	"Pichat/pkg/ent"
	"Pichat/pkg/entities"
)

type Service interface {
	InsertUser(user *entities.User) (*entities.User, error)
	ListUsers() ([]*entities.User, error)
	GetUserByID(id int) (*entities.User, error)
	GetUserByEmail(email string) (*entities.User, error)
	UpdateUser(user *entities.User, id int) (*entities.User, error)
	RemoveUser(id int) error
	CheckDuplicateEmail(email string) (bool, error)
}

type service struct {
	repo Repository
}

func (s *service) InsertUser(user *entities.User) (*entities.User, error) {
	return s.repo.CreateUser(user)
}

func (s *service) ListUsers() ([]*entities.User, error) {
	return s.repo.ReadUser()
}

func (s *service) GetUserByID(id int) (*entities.User, error) {
	return s.repo.GetUserById(id)
}

func (s *service) GetUserByEmail(email string) (*entities.User, error) {
	return s.repo.GetUserByEmail(email)
}

func (s *service) UpdateUser(user *entities.User, id int) (*entities.User, error) {
	return s.repo.UpdateUser(user, id)
}

func (s *service) RemoveUser(id int) error {
	return s.repo.DeleteUser(id)
}

func (s *service) CheckDuplicateEmail(email string) (bool, error) {
	_, err := s.repo.GetUserByEmail(email)

	if err != nil {
		if ent.IsNotFound(err) {
			return false, nil
		} else {
			return true, err
		}
	}

	return true, nil
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}

/*
Go 의 에러처리
 - err 값을 함수 return
 - err 중요하지 않음 -> 일반적인 에러처리 로직
 - err 존나 심각함 -> panic 으로 프로그램을 종료시킴
*/
