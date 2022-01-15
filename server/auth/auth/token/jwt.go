package token

import (
	"crypto/rsa"
	"fmt"
	tokenutil "notion/shared/auth/token"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

// JWTTokenGen generates a JWT token.
type JWTTokenGen struct {
	privateKey *rsa.PrivateKey
	refreshKey string
	issuer     string
	nowFunc    func() time.Time
}

// NewJWTokenGen creates a JWTTokenGen.
func NewJWTTokenGen(issuer string, privateKey *rsa.PrivateKey, refreshKey string) *JWTTokenGen {
	return &JWTTokenGen{
		issuer:     issuer,
		nowFunc:    time.Now,
		privateKey: privateKey,
		refreshKey: refreshKey,
	}
}

// GenAccessToken generates a access token.
func (t *JWTTokenGen) GenAccessToken(accountID string, exprie time.Duration) (string, error) {
	nowSec := t.nowFunc().Unix()
	tkn := jwt.NewWithClaims(jwt.SigningMethodRS512, jwt.StandardClaims{
		Issuer:    t.issuer,
		IssuedAt:  nowSec,
		ExpiresAt: nowSec + int64(exprie.Seconds()),
		Subject:   accountID,
	})

	return tkn.SignedString(t.privateKey)
}

// GenRefreshToken generates a refresh token.
func (t *JWTTokenGen) GenRefreshToken(accountID string, exprie time.Duration) (string, error) {
	tid, err := uuid.NewRandom()
	if err != nil {
		return "", fmt.Errorf("cannot gen uuid %w", err)
	}
	nowSec := t.nowFunc().Unix()
	tkn := jwt.NewWithClaims(jwt.SigningMethodRS512, tokenutil.JWTRefreshClaims{
		AID: accountID,
		StandardClaims: jwt.StandardClaims{
			Issuer:    t.issuer,
			IssuedAt:  nowSec,
			ExpiresAt: nowSec + int64(exprie.Seconds()),
			Subject:   tid.String(),
		},
	})

	return tkn.SignedString(t.refreshKey)
}
