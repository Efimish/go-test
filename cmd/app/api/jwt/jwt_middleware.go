package jwt

import (
	"crypto/ed25519"
	"crypto/sha256"

	"github.com/go-chi/jwtauth/v5"
)

func NewTokenAuth(jwtSecret string) *jwtauth.JWTAuth {
	seed := sha256.Sum256([]byte(jwtSecret))
	key := ed25519.NewKeyFromSeed(seed[:])
	return jwtauth.New("EdDSA", key, key.Public())
}
