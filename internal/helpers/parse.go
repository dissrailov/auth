package helpers

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

func JwtParseKey(t *jwt.Token) (*jwt.Token, error) {
	_, ok := t.Method.(*jwt.SigningMethodHMAC)
	if !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
	}
	return t, nil
}
