package mux

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

var port = 3000
var host = fmt.Sprintf("127.0.0.1:%d", port)
var port2 = 3001
var host2 = fmt.Sprintf("127.0.0.1:%d", port2)

type Notification struct {
	Message string `json:"message"`
}

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func TestMux() {
	go startHttpServer1()
	go startHttpServer2()
	select {}
}

func startHttpServer1() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /number/{number}", func(w http.ResponseWriter, r *http.Request) {
		str := r.PathValue("number")
		fmt.Printf("Основной сервер получил запрос: GET /%s\n", str)
		num, _ := strconv.Atoi(str)
		resp, _ := http.Get(fmt.Sprintf("http://%s/number/%d", host2, num))
		var resp_int int
		json.NewDecoder(resp.Body).Decode(&resp_int)
		json.NewEncoder(w).Encode(resp_int)
	})
	fmt.Printf("Основной HTTP сервер запущен на http://%s\n", host)
	http.ListenAndServe(host, CORS(mux))
}

func startHttpServer2() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /number/{number}", func(w http.ResponseWriter, r *http.Request) {
		str := r.PathValue("number")
		fmt.Printf("Вторичный сервер получил запрос: GET /%s\n", str)

		num, _ := strconv.Atoi(str)
		json.NewEncoder(w).Encode(num * 2)
	})
	fmt.Printf("Вторичный HTTP сервер запущен на http://%s\n", host2)
	http.ListenAndServe(host2, CORS(mux))
}
