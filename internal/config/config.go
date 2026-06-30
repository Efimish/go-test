package config

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Host      string
	Port      uint16
	HostPort  string
	PublicURL string
	JWTSecret string
}

func Load() Config {
	godotenv.Load()
	config := Config{}

	config.Host = getEnvFallback("HOST", "192.168.10.20")
	config.Port = getEnvFallbackUint16("PORT", 3000)
	config.HostPort = net.JoinHostPort(config.Host, strconv.FormatUint(uint64(config.Port), 10))
	config.PublicURL = fmt.Sprintf("http://%s", config.HostPort)
	config.JWTSecret = getEnvRequired("JWT_SECRET")

	return config
}

func getEnvRequired(key string) string {
	value, found := os.LookupEnv(key)
	if !found {
		log.Panicf("Missing required environment variable: %s", key)
	}
	return value
}

func getEnvFallback(key, fallback string) string {
	value, found := os.LookupEnv(key)
	if !found {
		return fallback
	}
	return value
}

func getEnvFallbackUint16(key string, fallback uint16) uint16 {
	str, found := os.LookupEnv(key)
	if !found {
		return fallback
	}
	value, err := strconv.ParseUint(str, 10, 16)
	if err != nil {
		log.Panicf("Cannot convert environment variable to uint16: %s", key)
	}
	return uint16(value)
}
