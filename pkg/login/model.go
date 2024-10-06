package login

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type SSHLogin struct {
	IP       string `json:"ip"`
	Username string `json:"username"`
	Port     string `json:"port"`
	Password string `json:"password"`
}

type Claims struct {
	Username  string `json:"username"`
	SessionID string `json:"session_id"`
	jwt.RegisteredClaims
}

func (c *Claims) Valid() error {
	if c.ExpiresAt != nil && time.Now().After(c.ExpiresAt.Time) {
		return jwt.NewValidationError("Token is expired", jwt.ValidationErrorExpired)
	}
	return nil
}
