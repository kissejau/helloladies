package repository

import (
	"helloladies/apps/backend/internal/model"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

const (
	usersTable = "users"
)

type UsersRepository interface {
	CreateUser(model.UserDto) error
	GetUserById(string) (model.UserDto, error)
	GetUserByEmail(string) (model.UserDto, error)
	UpdateUser(model.UserDto) (model.UserDto, error)
	DeleteUser(string) error
	List() ([]model.UserDto, error)
}

type AuthRepository interface {
}

type Repositories struct {
	AuthRepository
	UsersRepository
}

func NewRepositories(db *sqlx.DB, log *logrus.Logger) *Repositories {
	return &Repositories{
		AuthRepository:  NewAuthRepository(db),
		UsersRepository: NewUsersRepository(db),
	}
}
