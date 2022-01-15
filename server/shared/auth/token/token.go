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

// NewJWTTokenVerifier creates a jwt token verifier.
func NewJWTTokenVerifier(pubKey *rsa.PublicKey) *JWTTokenVerifier {
	return &JWTTokenVerifier{
		PublicKey: pubKey,
	}
}

// VerifyAccessToken verifes a access token and returns account id.
func (v *JWTTokenVerifier) VerifyAccessToken(token string) (string, error) {
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

// JWTRefreshClaims defines a refresh claims.
type JWTRefreshClaims struct {
	jwt.StandardClaims
	AID string `json:"aid"` // store account id
}

// JWTRefreshTokenVerifier verifies jwt refresh tokens
type JWTRefreshTokenVerifier struct {
	refreshKey string
}

// NewJWTRefreshTokenVerifier creates a jwt token verifier.
func NewJWTRefreshTokenVerifier(refreshKey string) *JWTRefreshTokenVerifier {
	return &JWTRefreshTokenVerifier{
		refreshKey: refreshKey,
	}
}

// VerifyRefreshToken verifes a refresh token and returns account id.
func (v *JWTRefreshTokenVerifier) VerifyRefreshToken(token string) (string, error) {
	var clm JWTRefreshClaims
	_, err := jwt.ParseWithClaims(token, &clm, func(*jwt.Token) (interface{}, error) {
		return v.refreshKey, nil
	})
	if err != nil {
		return "", fmt.Errorf("cannot parse token: %v", err)
	}

	if err = clm.Valid(); err != nil {
		return "", fmt.Errorf("claim not valid: %v", err)
	}

	return clm.AID, nil
}
