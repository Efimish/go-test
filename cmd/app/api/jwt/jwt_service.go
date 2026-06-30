package jwt

import (
	"crypto/ed25519"
	"crypto/sha256"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	privateKey ed25519.PrivateKey
}

func NewJWTService(jwtSecret string) JWTService {
	seed := sha256.Sum256([]byte(jwtSecret))
	privateKey := ed25519.NewKeyFromSeed(seed[:])
	return JWTService{privateKey: privateKey}
}

func (s JWTService) Sign(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, jwt.MapClaims{
		"sub": fmt.Sprintf("%d", userID),
	})
	return token.SignedString(s.privateKey)
}

func (s JWTService) Verify(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (any, error) {
			return s.privateKey, nil
		},
		jwt.WithValidMethods([]string{jwt.SigningMethodEdDSA.Alg()}),
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
