package jwt

import (
	"helloladies/internal/model"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTGenerator struct {
	cfg Config
}

func NewJWTGenerator(cfg Config) *JWTGenerator {
	return &JWTGenerator{cfg: cfg}
}

func (gen *JWTGenerator) GenerateToken(tokenIn model.TokenIn) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, model.TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(gen.cfg.TokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: tokenIn.UserId,
	})

	return token.SignedString([]byte(gen.cfg.SecretKey))
}
