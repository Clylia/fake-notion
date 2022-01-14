package token

import (
	"crypto/rsa"
	"fmt"

	"github.com/golang-jwt/jwt"
)

// JWTTokenVerifier verifies jwt access tokens.
type JWTTokenVerifier struct {
	PublicKey *rsa.PublicKey
}

// Verify verifes a token and returns account id.
func (v *JWTTokenVerifier) Verify(token string) (string, error) {
	var clm jwt.StandardClaims
	_, err := jwt.ParseWithClaims(token, &clm, func(*jwt.Token) (interface{}, error) {
		return v.PublicKey, nil
	})
	if err != nil {
		return "", fmt.Errorf("cannot parse token: %v", err)
	}
	
	if err = clm.Valid(); err != nil {
		return "", fmt.Errorf("claim not valid: %v", err)
	}

	return clm.Subject, nil
}
