package users

import (
	"Pichat/pkg/ent"
	"Pichat/pkg/ent/user"
	"Pichat/pkg/entities"
	"context"
	"github.com/gofiber/fiber/v2"
)

type Repository interface {
	CreateUser(user *entities.User) (*entities.User, error)
	ReadUser() ([]*entities.User, error)
	GetUserById(id int) (*entities.User, error)
	GetUserByEmail(email string) (*entities.User, error)
	UpdateUser(user *entities.User, id int) (*entities.User, error)
	DeleteUser(id int) error
}

type repository struct {
	C       *ent.Client
	Context context.Context
}

func NewRepo(C *ent.Client, ctx context.Context) Repository {
	return &repository{
		C:       C,
		Context: ctx,
	}
}

func (r *repository) CreateUser(user *entities.User) (*entities.User, error) {

	u, err := r.C.User.Create().
		SetName(user.Name).
		SetEmail(user.Email).
		SetPassword(user.Password).
		Save(r.Context)

	return entities.FromEntUser(u), err
}

func (r *repository) ReadUser() ([]*entities.User, error) {
	users, err := r.C.User.Query().All(r.Context)
	if err != nil {
		return nil, err
	}

	usersEntities := make([]*entities.User, 0, len(users))
	for i, u := range users {
		usersEntities[i] = entities.FromEntUser(u)
	}

	return usersEntities, nil
}

func (r *repository) GetUserById(id int) (*entities.User, error) {

	u, err := r.C.User.Query().Where(user.IDEQ(id)).First(r.Context)
	if err != nil {
		return nil, err
	}

	return entities.FromEntUser(u), err
}

func (r *repository) GetUserByEmail(email string) (*entities.User, error) {

	u, err := r.C.User.Query().Where(user.EmailEQ(email)).First(r.Context)
	if err != nil {
		return nil, err
	}

	return entities.FromEntUser(u), err
}

func (r *repository) UpdateUser(update *entities.User, id int) (*entities.User, error) {

	query := r.C.User.Update().Where(user.IDEQ(id))

	if update.Name != "" {
		query = query.SetName(update.Name)
	}
	if update.Email != "" {
		query = query.SetEmail(update.Email)
	}
	if update.Password != "" {
		query = query.SetPassword(update.Password)
	}

	_, err := query.Save(r.Context)

	if err != nil {
		return nil, err
	}

	u, err := r.C.User.Query().Where(user.IDEQ(id)).First(r.Context)
	if err != nil {
		return nil, err
	}

	return entities.FromEntUser(u), err
}

func (r *repository) DeleteUser(id int) error {

	rows, err := r.C.User.Delete().Where(user.IDEQ(id)).Exec(r.Context)
	if err != nil {
		return err
	}

	if rows < 1 {
		return fiber.NewError(fiber.StatusNotFound, "User not found")
	}

	return nil
}

/*
User CRUD 란

User - 도메인, User 라는 기능을 담당한다는거
CRUD - Create, Read, Update, Delete 약자
     - 생성, 조회, 수정, 삭제
     -> RESTAPI 의 기본 요소
     -> POST, GET, PUT/PATCH, DELETE

암튼, CRUD 가 왜 중요하냐
API 가 정상적으로 동작하려면 이 4개는 필요함
*/

/*
백엔드 레이어드 아키텍쳐

 - Controller : 사용자에게 노출되는 API
 - Service : 비즈니스 로직
 - Repository : DB 에 직접 접근
*/
