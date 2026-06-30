package auth

import (
	"encoding/json"
	"net/http"

	"github.com/efimish/go-test/cmd/app/api/jwt"
	"github.com/go-chi/chi/v5"
)

func AuthController(authService AuthService, jwtService jwt.JWTService) func(chi.Router) {
	return func(r chi.Router) {
		r.Post("/login", handlePostAuthLogin(authService, jwtService))
	}
}

func handlePostAuthLogin(authService AuthService, jwtService jwt.JWTService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, 1<<20)

		var credentials Credentials
		if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
			http.Error(w, "Invalid JSON or body too large", http.StatusBadRequest)
			return
		}

		user, found := authService.Login(credentials.Username, credentials.Password)
		if !found {
			http.Error(w, "Wrong credentials", http.StatusUnauthorized)
			return
		}

		token, err := jwtService.Sign(user.ID)
		if err != nil {
			http.Error(w, "Failed to sign token", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(TokenObject{Token: token})
	}
}
