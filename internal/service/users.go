package service

import (
	"fmt"
	"helloladies/apps/backend/internal/model"
	"helloladies/apps/backend/internal/repository"

	"github.com/sirupsen/logrus"
)

const (
	errServer     = "server error"
	errDeleteUser = "error while deleting user"
	errUpdateUser = "error while updating user"
	errListUsers  = "error while listing users"
)

type UsersServiceImpl struct {
	usersRepo repository.UsersRepository
	log       *logrus.Logger
}

func NewUserService(usersRepo repository.UsersRepository, log *logrus.Logger) *UsersServiceImpl {
	return &UsersServiceImpl{
		usersRepo: usersRepo,
		log:       log,
	}
}

func (s *UsersServiceImpl) GetUserById(id string) (model.User, error) {
	var user model.User
	userDto, err := s.usersRepo.GetUserById(id)
	if err != nil {
		s.log.Printf("s.usersRepo.GetUserById: %s\n", err.Error())
		return model.User{}, fmt.Errorf(errServer)
	}

	user = model.UserDtoToUser(userDto)
	return user, nil
}

func (s *UsersServiceImpl) UpdateUser(id string, user model.User) (model.User, error) {
	dto := model.UserDto{
		Id:        id,
		Email:     user.Email,
		Name:      user.Name,
		BirthDate: user.BirthDate,
	}
	dto, err := s.usersRepo.UpdateUser(dto)
	if err != nil {
		s.log.Printf("s.usersRepo.UpdateUser: %s", err.Error())
		return model.User{}, fmt.Errorf(errUpdateUser)
	}
	return model.UserDtoToUser(dto), nil
}

func (s *UsersServiceImpl) DeleteUser(id string) error {
	if err := s.usersRepo.DeleteUser(id); err != nil {
		s.log.Printf("s.DeleteUser: %s", err.Error())
		return fmt.Errorf(errDeleteUser)
	}
	return nil
}

func (s *UsersServiceImpl) List() ([]model.User, error) {
	var users []model.User
	userDtos, err := s.usersRepo.List()
	if err != nil {
		s.log.Printf("s.usersRepo.List: %s", err.Error())
		return []model.User{}, fmt.Errorf(errListUsers)
	}

	for _, userDto := range userDtos {
		user := model.UserDtoToUser(userDto)
		users = append(users, user)
	}
	return users, nil
}
