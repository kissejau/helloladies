package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"helloladies/internal/model"
)

type TeachersRepositoryImpl struct {
	db *sqlx.DB
}

func NewTeachersRepository(db *sqlx.DB) TeachersRepositoryImpl {
	return TeachersRepositoryImpl{
		db: db,
	}
}

func (r TeachersRepositoryImpl) CreateTeacher(teacherDto model.TeacherDto) error {
	query := fmt.Sprintf(`INSERT INTO %s (id, code, name, surname, patronymic, degree, univ_id, confirmed)
		VALUES (:id, :code, :name, :surname, :patronymic, :degree, :univ_id, :confirmed)`, teachersTable)

	if _, err := r.db.NamedExec(query, teacherDto); err != nil {
		return fmt.Errorf("r.db.NamedExec: %w", err)
	}
	return nil
}

func (r TeachersRepositoryImpl) List() ([]model.TeacherDto, error) {
	query := fmt.Sprintf(`SELECT id, code, name, surname, patronymic, degree, univ_id, confirmed FROM %s`, teachersTable)

	rows, err := r.db.Query(query)
	if err != nil {
		return []model.TeacherDto{}, fmt.Errorf("r.db.Query: %w", err)
	}

	var teacherDtos []model.TeacherDto
	for rows.Next() {
		var teacherDto model.TeacherDto
		if err := rows.Scan(&teacherDto.Id, &teacherDto.Code,
			&teacherDto.Name, &teacherDto.Surname,
			&teacherDto.Patronymic, &teacherDto.Degree,
			&teacherDto.UnivId, &teacherDto.Confirmed); err != nil {
			return []model.TeacherDto{}, fmt.Errorf("rows.Scan: %w", err)
		}
		teacherDtos = append(teacherDtos, teacherDto)
	}
	return teacherDtos, nil
}

func (r TeachersRepositoryImpl) UpdateTeacher(teacherDto model.TeacherDto) (model.TeacherDto, error) {
	query := fmt.Sprintf(`UPDATE %s SET code=:code, name=:name, surname=:surname,
	 patronymic=:patronymic, degree=:degree, univ_id=:univ_id, confirmed=:confirmed WHERE id=:id`, teachersTable)

	if _, err := r.db.NamedExec(query, teacherDto); err != nil {
		return model.TeacherDto{}, fmt.Errorf("r.db.NamedExec: %w", err)
	}
	return teacherDto, nil
}

func (r TeachersRepositoryImpl) DeleteTeacher(id string) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id=$1`, teachersTable)

	if _, err := r.db.Exec(query, id); err != nil {
		return fmt.Errorf("r.db.Exec: %w", err)
	}
	return nil
}

func (r TeachersRepositoryImpl) GetIdByCode(teacherCode string) (string, error) {
	query := fmt.Sprintf(`SELECT id FROM %s WHERE code=$1`, teachersTable)

	rows, err := r.db.Query(query, teacherCode)
	if err != nil {
		return "", fmt.Errorf("r.db.Query: %w", err)
	}

	fl := rows.Next()
	if !fl {
		return "", fmt.Errorf("incorrect id")
	}

	var id string
	if err := rows.Scan(&id); err != nil {
		return "", fmt.Errorf("r.db.Query: %w", err)
	}
	return id, nil
}

func (r TeachersRepositoryImpl) GetTeachersByUniv(univCode string) ([]model.TeacherDto, error) {
	query := fmt.Sprintf(
		`SELECT t.id, t.code, t.name, t.surname, t.patronymic, t.degree, t.univ_id, t.confirmed
		FROM %s AS t
		JOIN %s AS u ON u.id = t.univ_id
		WHERE u.code=$1`, teachersTable, univsTable)

	rows, err := r.db.Query(query, univCode)
	if err != nil {
		return []model.TeacherDto{}, fmt.Errorf("r.db.Query: %w", err)
	}

	var teacherDtos []model.TeacherDto
	for rows.Next() {
		var teacherDto model.TeacherDto
		if err := rows.Scan(&teacherDto.Id, &teacherDto.Code,
			&teacherDto.Name, &teacherDto.Surname,
			&teacherDto.Patronymic, &teacherDto.Degree,
			&teacherDto.UnivId, &teacherDto.Confirmed); err != nil {
			return []model.TeacherDto{}, fmt.Errorf("rows.Scan: %w", err)
		}
		teacherDtos = append(teacherDtos, teacherDto)
	}
	return teacherDtos, nil
}
