package numberovernats

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/nats-io/nats.go"
)

var port = 3000
var host = fmt.Sprintf("127.0.0.1:%d", port)

func Test() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	go startHttpServer(nc)
	go startService(nc)

	select {}
}

func startHttpServer(nc *nats.Conn) {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{message}", func(w http.ResponseWriter, r *http.Request) {
		message := r.PathValue("message")
		fmt.Printf("Сервер получил запрос: GET /%s\n", message)

		reply, err := nc.Request("service.query", []byte(message), 2*time.Second)
		if err != nil {
			log.Printf("[HTTP] Request failed: %v", err)
			http.Error(w, "Service timeout or error", http.StatusGatewayTimeout)
			return
		}

		number, _ := strconv.Atoi(string(reply.Data))
		json.NewEncoder(w).Encode(number)
	})
	fmt.Printf("HTTP сервер запущен на http://%s\n", host)

	http.ListenAndServe(host, mux)
}

func startService(nc *nats.Conn) {
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
