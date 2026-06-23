package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func handleGetNotifications(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if len(authHeader) == 0 {
		http.Error(w, "Отсутствует header Authorization", http.StatusUnauthorized)
		return
	}
	if !strings.HasPrefix(authHeader, "Bearer ") {
		http.Error(w, "Header Authorization должен начинаться с 'Bearer '", http.StatusUnauthorized)
	}
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	claims, err := NewJwtService(Config.JwtSecret).Verify(tokenString)
	if err != nil {
		http.Error(w, "Ошибка проверки токена", http.StatusUnauthorized)
		return
	}
	userIdStr, err := claims.GetSubject()
	if err != nil {
		http.Error(w, "Ошибка получения sub", http.StatusUnauthorized)
		return
	}
	userId, err := strconv.ParseUint(userIdStr, 10, 0)
	if err != nil {
		http.Error(w, "Ошибка parseuint", http.StatusUnauthorized)
		return
	}

	userNotifications := notifications[uint(userId)]
	json.NewEncoder(w).Encode(userNotifications)
}
