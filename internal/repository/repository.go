package repository

import (
	"helloladies/internal/model"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

const (
	usersTable    = "users"
	citiesTable   = "cities"
	univsTable    = "univs"
	teachersTable = "teachers"
	reviewsTable  = "reviews"
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
	CreateUniv(univDto model.UnivDto) error
	List() ([]model.UnivDto, error)
	UpdateUniv(univDto model.UnivDto) (model.UnivDto, error)
	DeleteUniv(id string) error
	GetIdByCode(univCode string) (string, error)
	GetUnivsByCity(cityCode string) ([]model.UnivDto, error)
}

type TeachersRepository interface {
	CreateTeacher(teacherDto model.TeacherDto) error
	List() ([]model.TeacherDto, error)
	UpdateTeacher(teacherDto model.TeacherDto) (model.TeacherDto, error)
	DeleteTeacher(id string) error
	GetIdByCode(teacherCode string) (string, error)
	GetTeachersByUniv(univCode string) ([]model.TeacherDto, error)
}

type ReviewsRepository interface {
	CreateReview(review model.ReviewDto) error
	List() ([]model.ReviewDto, error)
	GetReviewsByTeacher(teacherId string) ([]model.ReviewDto, error)
	UpdateReview(reviewDto model.ReviewDto) (model.ReviewDto, error)
	DeleteReview(reviewId string) error
	GetIdByCode(reviewCode string) (string, error)
}

type AuthRepository interface {
}

type Repositories struct {
	AuthRepository
	UsersRepository
	CitiesRepository
	UnivsRepository
	TeachersRepository
	ReviewsRepository
}

func NewRepositories(db *sqlx.DB, log *logrus.Logger) *Repositories {
	return &Repositories{
		AuthRepository:     NewAuthRepository(db),
		UsersRepository:    NewUsersRepository(db),
		CitiesRepository:   NewCitiesRepository(db),
		UnivsRepository:    NewUnivsRepository(db),
		TeachersRepository: NewTeachersRepository(db),
		ReviewsRepository:  NewReviewsRepository(db),
	}
}
