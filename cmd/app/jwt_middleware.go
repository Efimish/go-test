package main

import (
	"crypto/ed25519"
	"crypto/sha256"

	. "github.com/efimish/go-test/cmd/app/config"
	"github.com/go-chi/jwtauth/v5"
)

var tokenAuth *jwtauth.JWTAuth

func init() {
	seed := sha256.Sum256([]byte(Config.JwtSecret))
	key := ed25519.NewKeyFromSeed(seed[:])
	tokenAuth = jwtauth.New("EdDSA", key, key.Public())
}
