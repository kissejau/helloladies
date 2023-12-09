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

func (r UsersRepositoryImpl) GetUserById(id string) (model.UserDto, error) {
	var userDto model.UserDto
	query := fmt.Sprintf("SELECT id, email, password, name, birth_date, is_admin FROM %s WHERE id=$1 LIMIT 1", usersTable)

	row := r.db.QueryRow(query, id)
	if err := row.Scan(&userDto.Id,
		&userDto.Email, &userDto.Password,
		&userDto.Name, &userDto.BirthDate, &userDto.IsAdmin); err != nil {
		return model.UserDto{}, fmt.Errorf("row.Scan: %w", err)
	}
	return userDto, nil
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

func (r UsersRepositoryImpl) UpdateUser(dto model.UserDto) (model.UserDto, error) {
	query := fmt.Sprintf("UPDATE %v SET email=:email, name=:name, birth_date=:birth_date WHERE id=:id",
		usersTable)
	if _, err := r.db.NamedExec(query, dto); err != nil {
		return model.UserDto{}, fmt.Errorf("r.db.NamedExec: %w", err)
	}
	return dto, nil
}

func (r UsersRepositoryImpl) DeleteUser(id string) error {
	query := fmt.Sprintf("DELETE FROM %v WHERE id=$1", usersTable)

	if _, err := r.db.Exec(query, id); err != nil {
		return fmt.Errorf("r.db.NamedExec: %w", err)
	}
	return nil
}

func (r UsersRepositoryImpl) List() ([]model.UserDto, error) {
	var userDtos []model.UserDto
	query := fmt.Sprintf("SELECT id, email, password, name, birth_date FROM %v", usersTable)

	rows, err := r.db.Query(query)
	if err != nil {
		return []model.UserDto{}, fmt.Errorf("r.db.Query: %w", err)
	}
	for rows.Next() {
		var userDto model.UserDto
		if err := rows.Scan(&userDto.Id,
			&userDto.Email, &userDto.Password,
			&userDto.Name, &userDto.BirthDate); err != nil {
			return []model.UserDto{}, fmt.Errorf("rows.Scan: %w", err)
		}
		userDtos = append(userDtos, userDto)
	}
	return userDtos, nil
}
