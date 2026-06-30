package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Config struct {
	Port      uint16
	Host      string
	JwtSecret string
}

func init() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
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
	Config.Host = fmt.Sprintf("192.168.10.20:%d", Config.Port)
	// JwtSecret
	envJwtSecret, exists := os.LookupEnv("JWT_SECRET")
	if !exists {
		panic("Missing JWT_SECRET env variable")
	} else {
		Config.JwtSecret = envJwtSecret
	}
}

// const pets = "🐶🐱🐸🐷🐵🐔🦊🦁🐴🐝🦄🐳🦜🦔"
// const flowers = "🌵🌴🌲🍄🌹🌻"
// const ingredients = "🍎🍐🍊🍋🍋‍🟩🍌🍉🍇🍓🫐🍒🍑🥭🍍🥥🥝🍅🥑🥒🌶️🫑🌽🥕🧄🧅🥔🫚🍞🧀🥚🧈🥓🥜"
// const dishes = "🥐🌭🍔🍟🍕🥪🌮🍣🍰🍿"
// const other = "🔥⭐️"
// const emojis = pets + flowers + ingredients + dishes + other
