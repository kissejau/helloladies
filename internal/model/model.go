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
}
