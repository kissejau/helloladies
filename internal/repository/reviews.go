package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"helloladies/internal/model"
)

type ReviewsRepositoryImpl struct {
	db *sqlx.DB
}

func NewReviewsRepository(db *sqlx.DB) ReviewsRepositoryImpl {
	return ReviewsRepositoryImpl{
		db: db,
	}
}

func (r ReviewsRepositoryImpl) CreateReview(review model.ReviewDto) error {
	query := fmt.Sprintf(`
		INSERT INTO %s (id, code, description, rating, teacher_id, user_id)
		VALUES (:id, :code, :description, :rating, :teacher_id, :user_id)`, reviewsTable)

	if _, err := r.db.NamedExec(query, review); err != nil {
		return fmt.Errorf("r.db.NamedExec: %w", err)
	}
	return nil
}

func (r ReviewsRepositoryImpl) List() ([]model.ReviewDto, error) {
	query := fmt.Sprintf(`SELECT id, code, description, rating, teacher_id, user_id FROM %s`, reviewsTable)

	rows, err := r.db.Query(query)
	if err != nil {
		return []model.ReviewDto{}, fmt.Errorf("r.db.Query: %w", err)
	}

	var reviewDtos []model.ReviewDto
	for rows.Next() {
		var reviewDto model.ReviewDto
		if err := rows.Scan(&reviewDto.Id, &reviewDto.Code,
			&reviewDto.Description, &reviewDto.Rating,
			&reviewDto.TeacherId, &reviewDto.UserId); err != nil {
			return []model.ReviewDto{}, fmt.Errorf("rows.Scan: %w", err)
		}
		reviewDtos = append(reviewDtos, reviewDto)
	}
	return reviewDtos, nil
}

func (r ReviewsRepositoryImpl) GetReviewsByTeacher(teacherId string) ([]model.ReviewDto, error) {
	query := fmt.Sprintf(`
		SELECT id, code, description, rating, teacher_id, user_id FROM %s
		WHERE teacher_id=$1`, reviewsTable)

	rows, err := r.db.Query(query, teacherId)
	if err != nil {
		return []model.ReviewDto{}, fmt.Errorf("r.db.Query: %w", err)
	}

	var reviewDtos []model.ReviewDto
	for rows.Next() {
		var reviewDto model.ReviewDto
		if err := rows.Scan(&reviewDto.Id, &reviewDto.Code,
			&reviewDto.Description, &reviewDto.Rating,
			&reviewDto.TeacherId, &reviewDto.UserId); err != nil {
			return []model.ReviewDto{}, fmt.Errorf("rows.Scan: %w", err)
		}
		reviewDtos = append(reviewDtos, reviewDto)
	}
	return reviewDtos, nil
}

func (r ReviewsRepositoryImpl) UpdateReview(reviewDto model.ReviewDto) (model.ReviewDto, error) {
	query := fmt.Sprintf(`
		UPDATE %s 
		SET code=:code, description=:description, rating=:rating, teacher_id=:teacher_id, user_id=:user_id 
		WHERE id=:id`, reviewsTable)

	if _, err := r.db.NamedExec(query, reviewDto); err != nil {
		return model.ReviewDto{}, fmt.Errorf("r.db.NamedExec: %w", err)
	}
	return reviewDto, nil
}

func (r ReviewsRepositoryImpl) DeleteReview(reviewId string) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id=$1`, reviewsTable)

	if _, err := r.db.Exec(query, reviewId); err != nil {
		return fmt.Errorf("r.db.Exec: %w", err)
	}
	return nil
}

func (r ReviewsRepositoryImpl) GetIdByCode(reviewCode string) (string, error) {
	var id string
	query := fmt.Sprintf("SELECT id FROM %s WHERE code=$1", reviewsTable)

	rows, err := r.db.Query(query, reviewCode)
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
