package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	. "github.com/efimish/go-test/cmd/app/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/nats-io/nats.go"
)

func handleRoot() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	go startHttpServer(nc, &notifications)
	go startService(nc, &notifications)

	select {}
}

func startHttpServer(nc *nats.Conn, notifications *map[uint][]Notification) {
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

	r.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	r.Route("/auth", authController)
	r.Route("/notifications", notificationsController)

	fmt.Printf("HTTP сервер запущен на http://%s\n", Config.Host)
	http.ListenAndServe(Config.Host, r)
}

func startService(nc *nats.Conn, notifications *map[uint][]Notification) {
	_, err := nc.Subscribe("service.query", func(m *nats.Msg) {
		number, _ := strconv.Atoi(string(m.Data))
		err := m.Respond([]byte(strconv.Itoa(number * 2)))
		if err != nil {
			log.Printf("[Service] Failed to respond: %v", err)
		}
	})
	if err != nil {
		log.Fatalf("Subscription error: %v", err)
	}

	log.Println("Service is ready to process requests...")
}
