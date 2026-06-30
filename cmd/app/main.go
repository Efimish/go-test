package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/efimish/go-test/internal/api/auth"
	"github.com/efimish/go-test/internal/api/jwt"
	"github.com/efimish/go-test/internal/api/notifications"
	"github.com/efimish/go-test/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	cfg := config.Load()

	authService := auth.NewService()
	jwtService := jwt.NewService(cfg.JWTSecret)
	notificationService := notifications.NewService(cfg.PublicURL)
	tokenAuth := jwt.NewTokenAuth(cfg.JWTSecret)

	authHandler := auth.NewHandler(authService, jwtService)
	notificationsHandler := notifications.NewHandler(notificationService, tokenAuth)

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
	r.Route("/auth", authHandler.Routes)
	r.Route("/notifications", notificationsHandler.Routes)

	fmt.Printf("HTTP сервер запущен на %s\n", cfg.PublicURL)
	log.Fatal(http.ListenAndServe(cfg.HostPort, r))
}
