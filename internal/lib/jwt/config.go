package jwt

import "time"

type Config struct {
	SecretKey string
	TokenTTL  time.Duration
}
