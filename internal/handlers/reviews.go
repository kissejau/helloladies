package handlers

import (
	"helloladies/internal/model"
	"helloladies/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

//	@Summary	CreateReview
//	@Security	Token
//	@Tags		reviews
//	@Accept		json
//	@Produce	json
//	@Param		teacher_code	query		string			true	"teacher's code"
//	@Param		input			body		model.Review	true	"review's data"
//	@Success	200				{object}	response.Success
//	@Failure	400				{object}	response.Error
//	@Router		/logged/reviews/create [post]
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

//	@Summary	ListReviews
//	@Security	Token
//	@Tags		reviews
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	[]model.Review
//	@Failure	400	{object}	response.Error
//	@Router		/logged/reviews/all [get]
func (h *Handler) ListReviews(c *gin.Context) {
	reviews, err := h.services.ReviewsService.List()
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusAccepted, reviews)
}

//	@Summary	GetReviewsByTeacher
//	@Security	Token
//	@Tags		reviews
//	@Accept		json
//	@Produce	json
//	@Param		teacher_code	query		string	true	"teacher's code"
//	@Success	200				{object}	[]model.Review
//	@Failure	400				{object}	response.Error
//	@Router		/logged/reviews/list [get]
func (h *Handler) GetReviewsByTeacher(c *gin.Context) {
	teacherCode := c.Query("teacher_code")

	reviews, err := h.services.ReviewsService.GetReviewsByTeacher(teacherCode)
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusAccepted, reviews)
}

//	@Summary	UpdateReview
//	@Security	Token
//	@Tags		reviews
//	@Accept		json
//	@Produce	json
//	@Param		teacher_code	query		string			true	"teacher's code"
//	@Param		input			body		model.Review	true	"review's data"
//	@Success	200				{object}	response.Success
//	@Failure	400				{object}	response.Error
//	@Router		/logged/reviews/update [put]
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

//	@Summary	DeleteReviews
//	@Security	Token
//	@Tags		reviews
//	@Accept		json
//	@Produce	json
//	@Param		review_code	query		string	true	"review's code"
//	@Success	200			{object}	response.Success
//	@Failure	400			{object}	response.Error
//	@Router		/logged/reviews/delete [delete]
func (h *Handler) DeleteReview(c *gin.Context) {
	reviewCode := c.Query("review_code")

	if err := h.services.ReviewsService.DeleteReview(reviewCode); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	response.NewSuccessResponse(c, http.StatusAccepted, "review was deleted")
}
