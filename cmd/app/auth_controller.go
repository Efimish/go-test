package main

import (
	"encoding/json"
	"net/http"
)

func handlePostAuthLogin(w http.ResponseWriter, r *http.Request) {
	// Ограничиваем максимальный размер тела (защита от DDOS)
	r.Body = http.MaxBytesReader(w, r.Body, 1048576) // 1 МБ

	var credentials Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, "Invalid JSON or body too large", http.StatusBadRequest)
		return
	}

	user, found := NewAuthService().Login(credentials.Username, credentials.Password)
	if !found {
		http.Error(w, "Wrong credentials", http.StatusBadRequest)
		return
	}

	token, err := NewJwtService(Config.JwtSecret).Sign(user.ID)
	if err != nil {
		http.Error(w, "Wrong credentials", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(TokenObject{Token: token})
}
