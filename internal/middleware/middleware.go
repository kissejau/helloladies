package middleware

import (
	jwtGen "helloladies/apps/backend/internal/lib/jwt"
	"helloladies/apps/backend/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware interface {
	VerifyToken(c *gin.Context)
	VerifyAdminPermissions(c *gin.Context)
}

type Middlewares struct {
	AuthMiddleware
}

func NewMiddlewares(cfg jwtGen.Config, usersService service.UsersService) Middlewares {
	return Middlewares{
		AuthMiddleware: NewAuthMiddleware(cfg, usersService),
	}
}
