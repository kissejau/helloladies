package repository

import (
	"fmt"
	"helloladies/internal/model"

	"github.com/jmoiron/sqlx"
)

type CitiesRepositoryImpl struct {
	db *sqlx.DB
}

func NewCitiesRepository(db *sqlx.DB) CitiesRepositoryImpl {
	return CitiesRepositoryImpl{
		db: db,
	}
}

func (r CitiesRepositoryImpl) CreateCity(cityDto model.CityDto) error {
	query := fmt.Sprintf("INSERT INTO %s (id, code, title, confirmed) VALUES (:id, :code, :title, :confirmed)", citiesTable)

	if _, err := r.db.NamedExec(query, cityDto); err != nil {
		return fmt.Errorf("r.db.NamedExec: %w", err)
	}
	return nil
}

func (r CitiesRepositoryImpl) List() ([]model.CityDto, error) {
	var cityDtos []model.CityDto
	query := fmt.Sprintf("SELECT id, code, title, confirmed FROM %s", citiesTable)

	rows, err := r.db.Query(query)
	if err != nil {
		return []model.CityDto{}, fmt.Errorf("r.db.Query: %w", err)
	}

	for rows.Next() {
		var cityDto model.CityDto
		if err := rows.Scan(&cityDto.Id, &cityDto.Code,
			&cityDto.Title, &cityDto.Confirmed); err != nil {
			return []model.CityDto{}, fmt.Errorf("rows.Scan: %w", err)
		}
		cityDtos = append(cityDtos, cityDto)
	}
	return cityDtos, nil
}

func (r CitiesRepositoryImpl) UpdateCity(id string, cityDto model.CityDto) (model.CityDto, error) {
	query := fmt.Sprintf("UPDATE %s SET code=:code, title=:title, confirmed=:confirmed WHERE id=:id", citiesTable)

	if _, err := r.db.NamedExec(query, cityDto); err != nil {
		return model.CityDto{}, fmt.Errorf("r.db.NamedExec: %w", err)
	}
	return cityDto, nil
}

func (r CitiesRepositoryImpl) DeleteCity(id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", citiesTable)

	if _, err := r.db.Exec(query, id); err != nil {
		return fmt.Errorf("r.db.Exec: %w", err)
	}
	return nil
}

func (r CitiesRepositoryImpl) GetIdByCode(code string) (string, error) {
	var id string
	query := fmt.Sprintf("SELECT id FROM %s WHERE code=$1", citiesTable)

	fmt.Println(query, code)
	rows, err := r.db.Query(query, code)
	if err != nil {
		return "", fmt.Errorf("r.db.Query: %w", err)
	}

	fl := rows.Next()
	if !fl {
		return "", fmt.Errorf("incorrect id")
	}

	if err := rows.Scan(&id); err != nil {
		return "", fmt.Errorf("r.db.Query: %w", err)
	}
	return id, nil
}
