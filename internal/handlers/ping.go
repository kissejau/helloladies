package handlers

import (
	"helloladies/apps/backend/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Ping(c *gin.Context) {
	response.NewSuccessResponse(c, http.StatusAccepted, "ping")
}
