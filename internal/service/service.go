package service

import (
	"fmt"
	"helloladies/internal/lib/jwt"
	"helloladies/internal/model"
	"helloladies/internal/repository"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

const (
	errIncorrectCode = "incorrect code"
)

type AuthService interface {
	SignIn(model.SignIn) (model.TokenOut, error)
	SignUp(model.SignUp) (model.TokenOut, error)
	Logout() error
	Refresh() (model.TokenOut, error)
}

type UsersService interface {
	GetUserById(string) (model.User, error)
	UpdateUser(string, model.User) (model.User, error)
	DeleteUser(string) error
	List() ([]model.User, error)
	IsAdmin(string) bool
}

type CitiesService interface {
	CreateCity(model.City) error
	List() ([]model.City, error)
	UpdateCity(string, model.City) (model.City, error)
	DeleteCity(string) error
}

type Services struct {
	AuthService
	UsersService
	CitiesService
}

func New(repos *repository.Repositories, log *logrus.Logger, jwtConfig jwt.Config) *Services {
	return &Services{
		AuthService:   NewAuthService(repos.UsersRepository, jwtConfig),
		UsersService:  NewUserService(repos.UsersRepository, log),
		CitiesService: NewCitiesService(repos.CitiesRepository, log),
	}
}

func generateCode(data string) string {
	return fmt.Sprintf("%s-%s", uuid.NewString()[:3], data)
}
