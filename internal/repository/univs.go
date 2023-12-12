package repository

import (
	"fmt"
	"helloladies/internal/model"

	"github.com/jmoiron/sqlx"
)

type UnivsRepositoryImpl struct {
	db *sqlx.DB
}

func NewUnivsRepository(db *sqlx.DB) UnivsRepositoryImpl {
	return UnivsRepositoryImpl{
		db: db,
	}
}

func (r UnivsRepositoryImpl) CreateUniv(univDto model.UnivDto) error {
	query := fmt.Sprintf("INSERT INTO %s (id, code, title, city_id, confirmed) VALUES (:id, :code, :title, :city_id, :confirmed)", univsTable)

	if _, err := r.db.NamedExec(query, univDto); err != nil {
		return fmt.Errorf("r.db.NamedExec: %w", err)
	}
	return nil
}

func (r UnivsRepositoryImpl) List() ([]model.UnivDto, error) {
	var univDtos []model.UnivDto
	query := fmt.Sprintf("SELECT id, code, title, city_id, confirmed FROM %s", univsTable)

	rows, err := r.db.Query(query)
	if err != nil {
		return []model.UnivDto{}, fmt.Errorf("r.db.Query: %w", err)
	}

	for rows.Next() {
		var univDto model.UnivDto
		if err := rows.Scan(&univDto.Id,
			&univDto.Code, &univDto.Title,
			&univDto.CityId, &univDto.Confirmed); err != nil {
			return []model.UnivDto{}, fmt.Errorf("rows.Scan: %w", err)
		}
		univDtos = append(univDtos, univDto)
	}
	return univDtos, nil
}

func (r UnivsRepositoryImpl) UpdateUniv(univDto model.UnivDto) (model.UnivDto, error) {
	query := fmt.Sprintf("UPDATE %s SET code=:code, title=:title, confirmed=:confirmed WHERE id=:id", univsTable)

	if _, err := r.db.NamedExec(query, univDto); err != nil {
		return model.UnivDto{}, fmt.Errorf("r.db.NamedExec: %w", err)
	}
	return univDto, nil
}

func (r UnivsRepositoryImpl) DeleteUniv(id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", univsTable)

	if _, err := r.db.Exec(query, id); err != nil {
		return fmt.Errorf("r.db.Exec: %w", err)
	}
	return nil
}

func (r UnivsRepositoryImpl) GetIdByCode(univCode string) (string, error) {
	var id string
	query := fmt.Sprintf("SELECT id FROM %s WHERE code=$1", univsTable)

	rows, err := r.db.Query(query, univCode)
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

func (r UnivsRepositoryImpl) GetUnivsByCity(cityCode string) ([]model.UnivDto, error) {
	var univDtos []model.UnivDto
	query := fmt.Sprintf(
		`SELECT u.id, u.code, u.title, u.city_id, u.confirmed 
		FROM %s AS c
		JOIN %s AS u ON u.city_id = c.id
		WHERE c.code=$1`, citiesTable, univsTable)

	rows, err := r.db.Query(query, cityCode)
	if err != nil {
		return []model.UnivDto{}, fmt.Errorf("r.db.Query: %w", err)
	}

	for rows.Next() {
		var univDto model.UnivDto
		if err := rows.Scan(&univDto.Id,
			&univDto.Code, &univDto.Title,
			&univDto.CityId, &univDto.Confirmed); err != nil {
			return []model.UnivDto{}, fmt.Errorf("rows.Scan: %w", err)
		}
		univDtos = append(univDtos, univDto)
	}
	return univDtos, nil
}
