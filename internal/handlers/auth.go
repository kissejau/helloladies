package handlers

import (
	"helloladies/internal/model"
	"helloladies/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SignIn(c *gin.Context) {
	var signIn model.SignIn
	c.Bind(&signIn)

	token, err := h.services.AuthService.SignIn(signIn)
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusAccepted, token)
}

func (h *Handler) SignUp(c *gin.Context) {
	var signUp model.SignUp
	c.Bind(&signUp)

	token, err := h.services.AuthService.SignUp(signUp)
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusAccepted, token)
}

func getUserId(c *gin.Context) string {
	return c.GetString("userId")
}
