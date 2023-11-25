package repository

import (
	"fmt"
	"helloladies/apps/backend/internal/model"

	"github.com/jmoiron/sqlx"
)

type UsersRepositoryImpl struct {
	db *sqlx.DB
}

func NewUsersRepository(db *sqlx.DB) UsersRepositoryImpl {
	return UsersRepositoryImpl{
		db: db,
	}
}

func (r UsersRepositoryImpl) CreateUser(signUpDto model.UserDto) error {
	query := fmt.Sprintf("INSERT INTO %s (id, email, password, name, birth_date) VALUES (:id, :email, :password, :name, :birth_date)", usersTable)

	if _, err := r.db.NamedExec(query, signUpDto); err != nil {
		return fmt.Errorf("r.db.NamedExec: %w", err)
	}
	return nil
}

func (r UsersRepositoryImpl) GetUserByEmail(email string) (model.UserDto, error) {
	var userDto model.UserDto
	query := fmt.Sprintf("SELECT id, email, password, name, birth_date FROM %s WHERE email=$1 LIMIT 1", usersTable)

	row := r.db.QueryRow(query, email)
	if err := row.Scan(&userDto.Id,
		&userDto.Email, &userDto.Password,
		&userDto.Name, &userDto.BirthDate); err != nil {
		return model.UserDto{}, err
	}
	return userDto, nil
}
