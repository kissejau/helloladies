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

type Services struct {
	AuthService
}

func New(repos *repository.Repositories, log *logrus.Logger, jwtConfig jwt.Config) *Services {
	return &Services{
		AuthService: NewAuthService(repos.UsersRepository, jwtConfig),
	}
}
