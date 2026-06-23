package main

import (
	"crypto/ed25519"
	"crypto/sha256"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type JwtService struct {
	privateKey ed25519.PrivateKey
}

func NewJwtService(jwtSecret string) JwtService {
	seed := sha256.Sum256([]byte(jwtSecret))
	privateKey := ed25519.NewKeyFromSeed(seed[:])
	return JwtService{privateKey}
}

func (s JwtService) Sign(userId uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, jwt.MapClaims{
		"sub": fmt.Sprintf("%d", userId),
	})
	return token.SignedString(s.privateKey)
}

func (s JwtService) Verify(tokenString string) (jwt.MapClaims, error) {
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
		return nil, err
	}

	return claims, nil
}
