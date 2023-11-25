package service

import (
	"database/sql"
	"errors"
	"fmt"
	"helloladies/apps/backend/internal/lib/jwt"
	"helloladies/apps/backend/internal/model"
	"helloladies/apps/backend/internal/repository"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const (
	errPaswordIsEmpty    = "empty password"
	errEmailIsEmpty      = "empty email"
	errPasswordsNotMatch = "passwords not match"
	errInvalidPassword   = "invalid password"
	errIncorrectPassword = "incorrect password"
	errUserExists        = "user with this email already exists"
	errUserNotExists     = "user with this email not exists"
	errGenerateToken     = "cant generate token"
	errDefualt           = "server error"
)

type AuthServiceImpl struct {
	usersRepo repository.UsersRepository
	tokenGen  *jwt.JWTGenerator
}

func NewAuthService(usersRepo repository.UsersRepository, jwtConfig jwt.Config) *AuthServiceImpl {
	return &AuthServiceImpl{
		usersRepo: usersRepo,
		tokenGen:  jwt.NewJWTGenerator(jwtConfig),
	}
}

func (s *AuthServiceImpl) SignIn(signIn model.SignIn) (model.TokenOut, error) {
	if err := checkEmptyPasswordOrEmail(signIn.Password, signIn.Email); err != nil {
		return model.TokenOut{}, err
	}

	userDto, err := s.usersRepo.GetUserByEmail(signIn.Email)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return model.TokenOut{}, fmt.Errorf(errUserNotExists)
		}
		return model.TokenOut{}, fmt.Errorf(errDefualt)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userDto.Password), []byte(signIn.Password)); err != nil {
		return model.TokenOut{}, fmt.Errorf(errIncorrectPassword)
	}

	token, err := s.tokenGen.GenerateToken(model.TokenIn{UserId: userDto.Id})
	if err != nil {
		return model.TokenOut{}, fmt.Errorf(errGenerateToken)
	}

	return model.TokenOut{Token: token}, nil
}

func (s *AuthServiceImpl) SignUp(signUp model.SignUp) (model.TokenOut, error) {
	if err := checkEmptyPasswordOrEmail(signUp.Password, signUp.Email); err != nil {
		return model.TokenOut{}, err
	}

	if signUp.Password != signUp.RepeatPassword {
		return model.TokenOut{}, fmt.Errorf(errPasswordsNotMatch)
	}

	userDto, err := s.usersRepo.GetUserByEmail(signUp.Email)
	if err == nil {
		return model.TokenOut{}, fmt.Errorf(errUserExists)

	}

	if !errors.Is(sql.ErrNoRows, err) {
		return model.TokenOut{}, fmt.Errorf(errDefualt)
	}

	cryptedPassword, err := bcrypt.GenerateFromPassword([]byte(signUp.Password), 12)
	if err != nil {
		return model.TokenOut{}, fmt.Errorf(errInvalidPassword)
	}

	userDto = model.UserDto{
		Id:       uuid.NewString(),
		Email:    signUp.Email,
		Password: string(cryptedPassword),
	}
	if err := s.usersRepo.CreateUser(userDto); err != nil {
		return model.TokenOut{}, fmt.Errorf(errDefualt)
	}

	token, err := s.tokenGen.GenerateToken(model.TokenIn{
		UserId: userDto.Id,
	})
	if err != nil {
		return model.TokenOut{}, fmt.Errorf(errGenerateToken)
	}

	return model.TokenOut{Token: token}, nil
}

func (s *AuthServiceImpl) Logout() error {
	return nil
}

func (s *AuthServiceImpl) Refresh() (model.TokenOut, error) {
	return model.TokenOut{}, nil
}

func checkEmptyPasswordOrEmail(password, email string) error {
	if password == "" {
		return fmt.Errorf(errPaswordIsEmpty)
	}
	if email == "" {
		return fmt.Errorf(errEmailIsEmpty)
	}
	return nil
}
