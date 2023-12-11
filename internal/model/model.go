package model

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type TokenIn struct {
	UserId string
}

type TokenClaims struct {
	jwt.StandardClaims
	UserId string
}

type TokenOut struct {
	Token string `json:"token"`
}

type SignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUp struct {
	Email          string `json:"email"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeat_password"`
}

type SignUpDto struct {
	Id       string `db:"id"`
	Email    string `db:"email"`
	Password string `db:"password"`
}

type User struct {
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	BirthDate time.Time `json:"birth_date"`
}

type UserDto struct {
	Id        string    `db:"id"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Name      string    `db:"name"`
	BirthDate time.Time `db:"birth_date"`
	IsAdmin   bool      `db:"is_admin"`
}

func UserDtoToUser(dto UserDto) User {
	return User{
		Email:     dto.Email,
		Name:      dto.Name,
		BirthDate: dto.BirthDate,
	}
}

type City struct {
	Code  string `json:"code"`
	Title string `json:"title"`
}

type CityDto struct {
	Id        string `db:"id"`
	Code      string `db:"code"`
	Title     string `db:"title"`
	Confirmed bool   `db:"confirmed"`
}

func CityDtoToCity(cityDto CityDto) City {
	return City{
		Code:  cityDto.Code,
		Title: cityDto.Title,
	}
}

type Univ struct {
	Code  string `json:"code"`
	Title string `json:"title"`
}

type UnivDto struct {
	Id        string `db:"id"`
	Code      string `db:"code"`
	Title     string `db:"title"`
	CityId    string `db:"city_id"`
	Confirmed bool   `db:"confirmed"`
}

func UnivDtoToUniv(univDto UnivDto) Univ {
	return Univ{
		Code:  univDto.Code,
		Title: univDto.Title,
	}
}

type Teacher struct {
	Code       string `json:"code"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Degree     string `json:"degree"`
}

type TeacherDto struct {
	Id         string `db:"id"`
	Code       string `db:"code"`
	Name       string `db:"name"`
	Surname    string `db:"surname"`
	Patronymic string `db:"patronymic"`
	Degree     string `db:"degree"`
	UnivId     string `db:"univ_id"`
	Confirmed  bool   `db:"confirmed"`
}

type Review struct {
	Code        string `json:"code"`
	Text        string `json:"description"`
	Description string `json:"rating"`
	TeacherId   string `json:"teacher_id"`
	UserId      string `json:"user_id"`
}

type ReviewDto struct {
	Id          string `db:"id"`
	Code        string `db:"code"`
	Description string `db:"description"`
	Rating      string `db:"rating"`
	TeacherId   string `db:"teacher_id"`
	UserId      string `db:"user_id"`
}
