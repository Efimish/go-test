package auth

import (
	"encoding/json"
	"net/http"

	"github.com/efimish/go-test/internal/api/jwt"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	authService Service
	jwtService  jwt.Service
}

func NewHandler(authService Service, jwtService jwt.Service) Handler {
	return Handler{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (h Handler) Routes(r chi.Router) {
	r.Post("/login", h.postLogin)
}

func (h Handler) postLogin(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)

	var credentials Credentials
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid JSON or body too large", http.StatusBadRequest)
		return
	}

	user, found := h.authService.Login(credentials.Username, credentials.Password)
	if !found {
		http.Error(w, "Wrong credentials", http.StatusUnauthorized)
		return
	}

	token, err := h.jwtService.Sign(user.ID)
	if err != nil {
		http.Error(w, "Failed to sign token", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(TokenObject{Token: token})
}
