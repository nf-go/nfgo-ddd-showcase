package util

import (
	"time"

	gojwt "github.com/pascaldekloe/jwt"
	"nfgo.ga/nfgo/nutil/jwt"
)

const (
	// JwtSecret -
	JwtSecret = "9U47#+ze=r4p2aCa6Hlm@trl#?-ph3-*egUPrAT+"
)

// NewToken -
func NewToken(subject string, expiration time.Time, set map[string]interface{}) (string, error) {
	return jwt.NewToken(JwtSecret, subject, expiration, set)
}

// ParseToken -
func ParseToken(token string) (*gojwt.Claims, error) {
	return jwt.ParseToken(JwtSecret, token)
}

// ValidateToken -
func ValidateToken(token string) (*gojwt.Claims, error) {
	return jwt.ValidateToken(JwtSecret, token)
}
