package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/efimish/go-test/cmd/app/api/auth"
	"github.com/efimish/go-test/cmd/app/api/jwt"
	"github.com/efimish/go-test/cmd/app/api/notifications"
	"github.com/efimish/go-test/cmd/app/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	cfg := config.Load()

	authService := auth.NewAuthService()
	jwtService := jwt.NewJWTService(cfg.JWTSecret)
	notificationService := notifications.NewNotificationService(cfg.PublicURL)
	tokenAuth := jwt.NewTokenAuth(cfg.JWTSecret)

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.ClientIPFromRemoteAddr)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(middleware.Compress(5, "application/json"))

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		MaxAge:         300,
	}))

	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	r.Route("/auth", auth.AuthController(authService, jwtService))
	r.Route("/notifications", notifications.NotificationsController(notificationService, tokenAuth))

	fmt.Printf("HTTP сервер запущен на %s\n", cfg.PublicURL)
	log.Fatal(http.ListenAndServe(cfg.HostPort, r))
}
