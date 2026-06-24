package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

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
	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	mux.HandleFunc("POST /auth/login", handlePostAuthLogin)
	mux.HandleFunc("GET /notifications", handleGetNotifications)

	fmt.Printf("HTTP сервер запущен на http://%s\n", Config.Host)
	http.ListenAndServe(Config.Host, mux)
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
