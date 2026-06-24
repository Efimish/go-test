package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/nats-io/nats.go"
)

var Config struct {
	Port      uint16
	Host      string
	JwtSecret string
}

func init() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	// PORT
	envPort, exists := os.LookupEnv("PORT")
	if !exists {
		Config.Port = 3000
	} else {
		port, _ := strconv.ParseUint(envPort, 10, 16)
		Config.Port = uint16(port)
	}
	// HOST
	Config.Host = fmt.Sprintf("127.0.0.1:%d", Config.Port)
	// JwtSecret
	envJwtSecret, exists := os.LookupEnv("JWT_SECRET")
	if !exists {
		panic("Missing JWT_SECRET env variable")
	} else {
		Config.JwtSecret = envJwtSecret
	}
}

const pets = "🐶🐱🐸🐷🐵🐔🦊🦁🐴🐝🦄🐳🦜🦔"
const flowers = "🌵🌴🌲🍄🌹🌻"
const ingredients = "🍎🍐🍊🍋🍋‍🟩🍌🍉🍇🍓🫐🍒🍑🥭🍍🥥🥝🍅🥑🥒🌶️🫑🌽🥕🧄🧅🥔🫚🍞🧀🥚🧈🥓🥜"
const dishes = "🥐🌭🍔🍟🍕🥪🌮🍣🍰🍿"
const other = "🔥⭐️"
const emojis = pets + flowers + ingredients + dishes + other

// func randomEmoji() string {
// 	r := emojis[rand.Int()%len(emojis)]
// 	min := r[0]
// 	max := r[1]
// 	n := rand.Intn(max-min+1) + min
// 	return html.UnescapeString("&#" + strconv.Itoa(n) + ";")
// }

func print(format string, a ...any) {
	fmt.Printf(
		"\033[32m[%s]\033[0m "+format,
		append(
			[]any{time.Now().Format("15:04:05.000000000")},
			a...,
		)...,
	)
}

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	subscribers := map[string]time.Duration{
		"1-🍎": 1000,
		"2-🍌": 1020,
		"3-🥝": 1200,
	}

	for name, ms := range subscribers {
		nc.QueueSubscribe("sub", "queue", func(m *nats.Msg) {
			print("%s) 🚀 Start: %s\n", name, m.Data)
			time.Sleep(time.Millisecond * ms)
			print("%s) 🎉 Finish: %s\n", name, m.Data)
		})
	}

	for i := 1; i <= 1000; i++ {
		// time.Sleep(time.Millisecond * 400)
		nc.Publish("sub", []byte(
			fmt.Sprintf("%d", i),
		))
	}
	select {}
}
