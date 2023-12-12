package handlers

import (
	"github.com/gin-gonic/gin"
	"helloladies/internal/model"
	"helloladies/pkg/response"
	"net/http"
)

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

func (h *Handler) ListTeachers(c *gin.Context) {
	teachers, err := h.services.TeachersService.List()
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusAccepted, teachers)
}

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

func (h *Handler) DeleteTeacher(c *gin.Context) {
	teacherCode := c.Query("teacher_code")

	if err := h.services.TeachersService.DeleteTeacher(teacherCode); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	response.NewSuccessResponse(c, http.StatusAccepted, "teacher was deleted")
}

func (h *Handler) GetTeacherByUniv(c *gin.Context) {
	univCode := c.Query("univ_code")

	teachers, err := h.services.TeachersService.GetTeachersByUniv(univCode)
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusAccepted, teachers)
}
