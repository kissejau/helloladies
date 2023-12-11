package repository

import (
	"helloladies/internal/model"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

const (
	usersTable  = "users"
	citiesTable = "cities"
	univsTable  = "univs"
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

type UnivsRepository interface {
	CreateUniv(model.UnivDto) error
	List() ([]model.UnivDto, error)
	UpdateUniv(univDtos model.UnivDto) (model.UnivDto, error)
	DeleteUniv(id string) error
	GetIdByCode(univCode string) (string, error)
	GetUnivsByCity(cityCode string) ([]model.UnivDto, error)
}

type AuthRepository interface {
}

type Repositories struct {
	AuthRepository
	UsersRepository
	CitiesRepository
	UnivsRepository
}

func NewRepositories(db *sqlx.DB, log *logrus.Logger) *Repositories {
	return &Repositories{
		AuthRepository:   NewAuthRepository(db),
		UsersRepository:  NewUsersRepository(db),
		CitiesRepository: NewCitiesRepository(db),
		UnivsRepository:  NewUnivsRepository(db),
	}
}
