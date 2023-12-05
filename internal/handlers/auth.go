package handlers

import (
	"helloladies/internal/model"
	"helloladies/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

//	@Summary	SignIn
//	@Tags		auth
//	@Accept		json
//	@Produce	json
//	@Param		input	body		model.SignIn	true	"sign in data"
//	@Success	200		{object}	model.TokenOut
//	@Failure	400		{object}	response.Error
//	@Router		/auth/sign-in [post]
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

//	@Summary	SignUp
//	@Tags		auth
//	@Accept		json
//	@Produce	json
//	@Param		input	body		model.SignUp	true	"sign up data"
//	@Success	200		{object}	model.TokenOut
//	@Failure	400		{object}	response.Error
//	@Router		/auth/sign-up [post]
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
