package service

import (
	"helloladies/apps/backend/internal/lib/jwt"
	"helloladies/apps/backend/internal/model"
	"helloladies/apps/backend/internal/repository"

	"github.com/sirupsen/logrus"
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

type Services struct {
	AuthService
	UsersService
}

func New(repos *repository.Repositories, log *logrus.Logger, jwtConfig jwt.Config) *Services {
	return &Services{
		AuthService:  NewAuthService(repos.UsersRepository, jwtConfig),
		UsersService: NewUserService(repos.UsersRepository, log),
	}
}
