package handlers

import (
	"helloladies/apps/backend/internal/model"
	"helloladies/apps/backend/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	errIncorrectBody = "incorrect body"
)

func (h *Handler) GetUser(c *gin.Context) {
	userId := getUserId(c)
	user, err := h.services.UsersService.GetUserById(userId)
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusAccepted, user)
}

func (h *Handler) DeleteUser(c *gin.Context) {
	userId := getUserId(c)
	if err := h.services.UsersService.DeleteUser(userId); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	response.NewSuccessResponse(c, http.StatusAccepted, "user was deleted successfully")
}

func (h *Handler) UpdateUser(c *gin.Context) {
	var user model.User

	userId := getUserId(c)
	if err := c.Bind(&user); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, errIncorrectBody)
		return
	}

	user, err := h.services.UsersService.UpdateUser(userId, user)
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusAccepted, user)
}

func (h *Handler) ListUsers(c *gin.Context) {
	// TODO:
	// add permission
	users, err := h.services.UsersService.List()
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusAccepted, users)
}
