package response

import (
	"github.com/gin-gonic/gin"
)

type Error struct {
	Message string `json:"error"`
}

type Success struct {
	Message string `json:"message"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, Error{Message: message})
}

func NewSuccessResponse(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, Success{Message: message})
}
