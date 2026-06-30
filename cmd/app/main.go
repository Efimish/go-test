package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	. "github.com/efimish/go-test/cmd/app/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
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
	r.Route("/auth", authController)
	r.Route("/notifications", notificationsController)
	r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(notifications[1])
	})

	fmt.Printf("HTTP сервер запущен на http://%s\n", Config.Host)
	http.ListenAndServe(Config.Host, r)
}
