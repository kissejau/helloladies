package service

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"helloladies/internal/model"
	"helloladies/internal/repository"
)

const (
	errCreateReview        = "error while creating review"
	errListReviews         = "error while listing reviews"
	errUpdateReview        = "error while updating review"
	errDeleteReview        = "error while deleting review"
	errGetReviewsByTeacher = "error while getting reviews by teacher"
)

type ReviewsServiceImpl struct {
	reviewsRepo repository.ReviewsRepository
	teacherRepo repository.TeachersRepository
	log         *logrus.Logger
}

func NewReviewsService(reviewsRepo repository.ReviewsRepository, teacherRepo repository.TeachersRepository, log *logrus.Logger) *ReviewsServiceImpl {
	return &ReviewsServiceImpl{
		reviewsRepo: reviewsRepo,
		teacherRepo: teacherRepo,
		log:         log,
	}
}

func (s *ReviewsServiceImpl) CreateReview(userId, teacherCode string, review model.Review) error {
	teacherId, err := s.teacherRepo.GetIdByCode(teacherCode)
	if err != nil {
		s.log.Printf("s.teacherRepo.GetIdByCode: %s", err.Error())
		return fmt.Errorf(errIncorrectCode)
	}

	reviewDto := model.ReviewDto{
		Id:          uuid.NewString(),
		Code:        generateCode(fmt.Sprintf("%s%d", review.Description[:1], review.Rating)),
		Description: review.Description,
		Rating:      review.Rating,
		UserId:      userId,
		TeacherId:   teacherId,
	}
	if err := s.reviewsRepo.CreateReview(reviewDto); err != nil {
		s.log.Printf("s.reviewsRepo.CreateReview: %s", err.Error())
		return fmt.Errorf(errCreateReview)
	}
	return nil
}

func (s *ReviewsServiceImpl) List() ([]model.Review, error) {
	reviewDtos, err := s.reviewsRepo.List()
	if err != nil {
		s.log.Printf("s.reviewsRepo.List: %s", err.Error())
		return []model.Review{}, fmt.Errorf(errListReviews)
	}

	var reviews []model.Review
	for _, reviewDto := range reviewDtos {
		reviews = append(reviews, model.ReviewDtoToReview(reviewDto))
	}
	return reviews, nil
}

func (s *ReviewsServiceImpl) GetReviewsByTeacher(teacherCode string) ([]model.Review, error) {
	teacherId, err := s.teacherRepo.GetIdByCode(teacherCode)
	if err != nil {
		s.log.Printf("s.teacherRepo.GetIdByCode: %s", err.Error())
		return []model.Review{}, fmt.Errorf(errIncorrectCode)
	}

	reviewDtos, err := s.reviewsRepo.GetReviewsByTeacher(teacherId)
	if err != nil {
		s.log.Printf("s.reviewsRepo.GetReviewsByTeacher: %s", err.Error())
		return []model.Review{}, fmt.Errorf(errGetReviewsByTeacher)
	}

	var reviews []model.Review
	for _, reviewDto := range reviewDtos {
		reviews = append(reviews, model.ReviewDtoToReview(reviewDto))
	}
	return reviews, nil
}

func (s *ReviewsServiceImpl) UpdateReview(userId, teacherCode string, review model.Review) (model.Review, error) {
	teacherId, err := s.teacherRepo.GetIdByCode(teacherCode)
	if err != nil {
		s.log.Printf("s.teacherRepo.GetIdByCode: %s", err.Error())
		return model.Review{}, fmt.Errorf(errIncorrectCode)
	}

	reviewId, err := s.reviewsRepo.GetIdByCode(review.Code)
	if err != nil {
		s.log.Printf("s.reviewsRepo.GetIdByCode: %s", err.Error())
		return model.Review{}, fmt.Errorf(errIncorrectCode)
	}

	reviewDto := model.ReviewDto{
		Id:          reviewId,
		Code:        review.Code,
		Description: review.Description,
		Rating:      review.Rating,
		UserId:      userId,
		TeacherId:   teacherId,
	}
	_, err = s.reviewsRepo.UpdateReview(reviewDto)
	if err != nil {
		s.log.Printf("s.reviewsRepo.UpdateReview: %s", err.Error())
		return model.Review{}, fmt.Errorf(errUpdateReview)
	}
	return review, nil
}

func (s *ReviewsServiceImpl) DeleteReview(reviewCode string) error {
	reviewId, err := s.reviewsRepo.GetIdByCode(reviewCode)
	if err != nil {
		s.log.Printf("s.reviewsRepo.GetIdByCode: %s", err.Error())
		return fmt.Errorf(errIncorrectCode)
	}

	if err := s.reviewsRepo.DeleteReview(reviewId); err != nil {
		s.log.Printf("s.reviews.DeleteReview: %s", err.Error())
		return fmt.Errorf(errDeleteCity)
	}
	return nil
}
