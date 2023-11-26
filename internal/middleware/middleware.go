package middleware

import (
	jwtGen "helloladies/apps/backend/internal/lib/jwt"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware interface {
	VerifyToken(c *gin.Context)
}

type Middlewares struct {
	AuthMiddleware
}

func NewMiddlewares(cfg jwtGen.Config) Middlewares {
	return Middlewares{
		AuthMiddleware: NewAuthMiddleware(cfg),
	}
}
