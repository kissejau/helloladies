package handlers

import (
	"helloladies/internal/model"
	"helloladies/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

//	@Summary	CreateTeacher
//	@Security	Token
//	@Tags		teachers
//	@Accept		json
//	@Produce	json
//	@Param		univ_code	query		string			true	"university's code"
//	@Param		input		body		model.Teacher	true	"teacher's data"
//	@Success	200			{object}	response.Success
//	@Failure	400			{object}	response.Error
//	@Router		/logged/teachers/create [post]
func (h *Handler) CreateTeacher(c *gin.Context) {
	var teacher model.Teacher

	univCode := c.Query("univ_code")

	if err := c.Bind(&teacher); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, errIncorrectBody)
		return
	}

	if err := h.services.TeachersService.CreateTeacher(univCode, teacher); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	response.NewSuccessResponse(c, http.StatusCreated, "teacher was created")
}

//	@Summary	ListTeachers
//	@Security	Token
//	@Tags		teachers
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	[]model.Teacher
//	@Failure	400	{object}	response.Error
//	@Router		/logged/teachers/all [get]
func (h *Handler) ListTeachers(c *gin.Context) {
	teachers, err := h.services.TeachersService.List()
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusAccepted, teachers)
}

//	@Summary	UpdateTeacher
//	@Security	Token
//	@Tags		teachers
//	@Accept		json
//	@Produce	json
//	@Param		univ_code	query		string			true	"university's code"
//	@Param		input		body		model.Teacher	true	"teacher's data"
//	@Success	200			{object}	response.Success
//	@Failure	400			{object}	response.Error
//	@Router		/logged/teachers/update [put]
func (h *Handler) UpdateTeacher(c *gin.Context) {
	univCode := c.Query("univ_code")

	var teacher model.Teacher
	if err := c.Bind(&teacher); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, errIncorrectBody)
		return
	}

	if _, err := h.services.TeachersService.UpdateTeacher(univCode, teacher); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	response.NewSuccessResponse(c, http.StatusAccepted, "teacher was updated")
}

//	@Summary	DeleteTeacher
//	@Security	Token
//	@Tags		teachers
//	@Accept		json
//	@Produce	json
//	@Param		teacher_code	query		string	true	"teacher's code"
//	@Success	200				{object}	response.Success
//	@Failure	400				{object}	response.Error
//	@Router		/logged/teachers/delete [delete]
func (h *Handler) DeleteTeacher(c *gin.Context) {
	teacherCode := c.Query("teacher_code")

	if err := h.services.TeachersService.DeleteTeacher(teacherCode); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	response.NewSuccessResponse(c, http.StatusAccepted, "teacher was deleted")
}

//	@Summary	GetTeachersByUniv
//	@Security	Token
//	@Tags		teachers
//	@Accept		json
//	@Produce	json
//	@Param		univ_code	query		string	true	"univ's code"
//	@Success	200			{object}	[]model.Teacher
//	@Failure	400			{object}	response.Error
//	@Router		/logged/teachers/list [get]
func (h *Handler) GetTeacherByUniv(c *gin.Context) {
	univCode := c.Query("univ_code")

	teachers, err := h.services.TeachersService.GetTeachersByUniv(univCode)
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusAccepted, teachers)
}
