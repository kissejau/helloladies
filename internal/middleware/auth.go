package middleware

import (
	"database/sql"
	"errors"
	"helloladies/apps/backend/internal/model"
	"helloladies/apps/backend/internal/service"
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
	errUserNotExists  = "user does not exists"
)

type AuthMiddlewareImpl struct {
	cfg          jwtGen.Config
	usersService service.UsersService
}

func NewAuthMiddleware(cfg jwtGen.Config, usersRepo service.UsersService) *AuthMiddlewareImpl {
	return &AuthMiddlewareImpl{
		cfg:          cfg,
		usersService: usersRepo,
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
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(m.cfg.SecretKey), nil
	})
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, errInvalidToken)
		return
	}

	if claims, ok := token.Claims.(*model.TokenClaims); ok && token.Valid {
		c.Set("userId", claims.UserId)
		if err := m.verifyUser(claims.UserId); err != nil {
			response.NewErrorResponse(c, http.StatusBadRequest, errUserNotExists)
			return
		}
		c.Next()
		return
	}

	response.NewErrorResponse(c, http.StatusBadRequest, errInvalidToken)
}

func (m AuthMiddlewareImpl) verifyUser(id string) error {
	if _, err := m.usersService.GetUserById(id); err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return err
		}
		// TODO:
		// handle it
		return err
	}
	return nil
}
