package notifications

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/golang-jwt/jwt/v5"
)

type Handler struct {
	notificationService Service
	tokenAuth           *jwtauth.JWTAuth
}

func NewHandler(notificationService Service, tokenAuth *jwtauth.JWTAuth) Handler {
	return Handler{
		notificationService: notificationService,
		tokenAuth:           tokenAuth,
	}
}

func (h Handler) Routes(r chi.Router) {
	r.Use(jwtauth.Verifier(h.tokenAuth))
	r.Use(jwtauth.Authenticator(h.tokenAuth))
	r.Get("/amount", h.getNotificationsAmount)
	r.Get("/list", h.getNotificationsList)
}

func (h Handler) getNotificationsAmount(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	tokenClaims := jwt.MapClaims(claims)

	userID, err := userIDFromClaims(tokenClaims)
	if err != nil {
		http.Error(w, "Invalid token subject", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(h.notificationService.AmountByUserID(userID))
}

func (h Handler) getNotificationsList(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	tokenClaims := jwt.MapClaims(claims)

	userID, err := userIDFromClaims(tokenClaims)
	if err != nil {
		http.Error(w, "Invalid token subject", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(h.notificationService.ListByUserID(userID))
}

func userIDFromClaims(claims jwt.MapClaims) (uint, error) {
	subject, err := claims.GetSubject()
	if err != nil {
		return 0, err
	}

	userID, err := strconv.ParseUint(subject, 10, 0)
	if err != nil {
		return 0, err
	}

	return uint(userID), nil
}
