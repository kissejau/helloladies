package handlers

import (
	"helloladies/internal/model"
	"helloladies/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

//	@Summary	GetUser
//	@Security	Token
//	@Tags		users
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	model.User
//	@Failure	400	{object}	response.Error
//	@Router		/logged/users/get [get]
func (h *Handler) GetUser(c *gin.Context) {
	userId := getUserId(c)
	user, err := h.services.UsersService.GetUserById(userId)
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusAccepted, user)
}

//	@Summary	DeleteUser
//	@Security	Token
//	@Tags		users
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	response.Success
//	@Failure	400	{object}	response.Error
//	@Router		/logged/users/delete [delete]
func (h *Handler) DeleteUser(c *gin.Context) {
	userId := getUserId(c)
	if err := h.services.UsersService.DeleteUser(userId); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	response.NewSuccessResponse(c, http.StatusAccepted, "user was deleted successfully")
}

//	@Summary	UpdateUser
//	@Security	Token
//	@Tags		users
//	@Accept		json
//	@Produce	json
//	@Param		input	body		model.User	true	"user data"
//
//	@Success	200		{object}	model.User
//	@Failure	400		{object}	response.Error
//	@Router		/logged/users/update [put]
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

//	@Summary	ListUsers
//	@Security	Token
//	@Tags		users
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	[]model.User
//	@Failure	400	{object}	response.Error
//	@Router		/logged/users/all [get]
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
