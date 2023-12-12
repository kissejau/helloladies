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
	UpdateCity(model.City) (model.City, error)
	DeleteCity(string) error
}

type UnivsService interface {
	CreateUniv(cityCode string, univ model.Univ) error
	List() ([]model.Univ, error)
	UpdateUniv(cityCode string, univ model.Univ) (model.Univ, error)
	DeleteUniv(univCode string) error
	GetUnivsByCity(cityCode string) ([]model.Univ, error)
}

type TeachersService interface {
	CreateTeacher(univCode string, teacher model.Teacher) error
	List() ([]model.Teacher, error)
	UpdateTeacher(univCode string, teacher model.Teacher) (model.Teacher, error)
	DeleteTeacher(teacherCode string) error
	GetTeachersByUniv(univCode string) ([]model.Teacher, error)
}

type Services struct {
	AuthService
	UsersService
	CitiesService
	UnivsService
	TeachersService
}

func New(repos *repository.Repositories, log *logrus.Logger, jwtConfig jwt.Config) *Services {
	return &Services{
		AuthService:     NewAuthService(repos.UsersRepository, jwtConfig),
		UsersService:    NewUserService(repos.UsersRepository, log),
		CitiesService:   NewCitiesService(repos.CitiesRepository, log),
		UnivsService:    NewUnivsService(repos.UnivsRepository, repos.CitiesRepository, log),
		TeachersService: NewTeachersService(repos.UnivsRepository, repos.TeachersRepository, log),
	}
}

func generateCode(data string) string {
	return fmt.Sprintf("%s-%s", uuid.NewString()[:3], data)
}
