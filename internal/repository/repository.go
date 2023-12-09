package repository

import (
	"helloladies/internal/model"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

const (
	usersTable  = "users"
	citiesTable = "cities"
)

type UsersRepository interface {
	CreateUser(model.UserDto) error
	GetUserById(string) (model.UserDto, error)
	GetUserByEmail(string) (model.UserDto, error)
	UpdateUser(model.UserDto) (model.UserDto, error)
	DeleteUser(string) error
	List() ([]model.UserDto, error)
}

type CitiesRepository interface {
	CreateCity(model.CityDto) error
	List() ([]model.CityDto, error)
	UpdateCity(string, model.CityDto) (model.CityDto, error)
	DeleteCity(string) error
	GetIdByCode(string) (string, error)
}

type AuthRepository interface {
}

type Repositories struct {
	AuthRepository
	UsersRepository
	CitiesRepository
}

func NewRepositories(db *sqlx.DB, log *logrus.Logger) *Repositories {
	return &Repositories{
		AuthRepository:   NewAuthRepository(db),
		UsersRepository:  NewUsersRepository(db),
		CitiesRepository: NewCitiesRepository(db),
	}
}
