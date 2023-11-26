package middleware

import (
	"helloladies/apps/backend/internal/model"
	"helloladies/apps/backend/pkg/response"
	"net/http"
	"strings"

	jwtGen "helloladies/apps/backend/internal/lib/jwt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

const (
	errMissedToken    = "token wasn't recieved"
	errIncorrectToken = "incorrect token"
	errAuth           = "error while auth"
	errInvalidToken   = "invalid token"
)

type AuthMiddlewareImpl struct {
	cfg jwtGen.Config
}

func NewAuthMiddleware(cfg jwtGen.Config) *AuthMiddlewareImpl {
	return &AuthMiddlewareImpl{
		cfg: cfg,
	}
}

func (m AuthMiddlewareImpl) VerifyToken(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if len(header) == 0 {
		response.NewErrorResponse(c, http.StatusBadRequest, errMissedToken)
		return
	}

	tokenData := strings.Split(header, " ")
	if len(tokenData) != 2 {
		response.NewErrorResponse(c, http.StatusBadRequest, errIncorrectToken)
		return
	}

	token, err := jwt.ParseWithClaims(tokenData[1], &model.TokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(m.cfg.SecretKey), nil
	})
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if claims, ok := token.Claims.(*model.TokenClaims); ok && token.Valid {
		c.Set("userId", claims.Id)
		c.Next()
		return
	}

	response.NewErrorResponse(c, http.StatusBadRequest, errInvalidToken)
}
