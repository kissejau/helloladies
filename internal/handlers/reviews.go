package handlers

import (
	"github.com/gin-gonic/gin"
	"helloladies/internal/model"
	"helloladies/pkg/response"
	"net/http"
)

func (h *Handler) CreateReview(c *gin.Context) {
	userId := getUserId(c)
	teacherCode := c.Query("teacher_code")

	var review model.Review
	if err := c.Bind(&review); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, errIncorrectBody)
		return
	}

	if err := h.services.ReviewsService.CreateReview(userId, teacherCode, review); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	response.NewSuccessResponse(c, http.StatusCreated, "review was created")
}

func (h *Handler) ListReviews(c *gin.Context) {
	reviews, err := h.services.ReviewsService.List()
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusAccepted, reviews)
}

func (h *Handler) GetReviewsByTeacher(c *gin.Context) {
	teacherCode := c.Query("teacher_code")

	reviews, err := h.services.ReviewsService.GetReviewsByTeacher(teacherCode)
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusAccepted, reviews)
}

func (h *Handler) UpdateReview(c *gin.Context) {
	userId := getUserId(c)
	teacherCode := c.Query("teacher_code")

	var review model.Review
	if err := c.Bind(&review); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, errIncorrectBody)
		return
	}

	if _, err := h.services.ReviewsService.UpdateReview(userId, teacherCode, review); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	response.NewSuccessResponse(c, http.StatusAccepted, "review was updated")
}

func (h *Handler) DeleteReview(c *gin.Context) {
	reviewCode := c.Query("review_code")

	if err := h.services.ReviewsService.DeleteReview(reviewCode); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	response.NewSuccessResponse(c, http.StatusAccepted, "review was deleted")
}
