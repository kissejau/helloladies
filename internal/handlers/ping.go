package handlers

import (
	"helloladies/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary	Ping
// @Tags		api
// @Accept		json
// @Produce	json
// @Success	200	{object}	response.Success
// @Router		/ping [get]
func (h *Handler) Ping(c *gin.Context) {
	response.NewSuccessResponse(c, http.StatusAccepted, "ping")
}
