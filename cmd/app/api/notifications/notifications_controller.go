package notifications

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/golang-jwt/jwt/v5"
)

func NotificationsController(notificationService NotificationService, tokenAuth *jwtauth.JWTAuth) func(chi.Router) {
	return func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator(tokenAuth))
		r.Get("/", handleGetNotifications(notificationService))
	}
}

func handleGetNotifications(notificationService NotificationService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, claims, _ := jwtauth.FromContext(r.Context())
		tokenClaims := jwt.MapClaims(claims)

		userID, err := userIDFromClaims(tokenClaims)
		if err != nil {
			http.Error(w, "Invalid token subject", http.StatusUnauthorized)
			return
		}

		json.NewEncoder(w).Encode(notificationService.ListByUserID(userID))
	}
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
