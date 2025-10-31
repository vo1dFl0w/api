package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	HTTPAddr string
	DatabaseURL  string
}

// Init config ...
func InitConfig() *Config {
	if err := godotenv.Load("config.env"); err != nil {
		log.Fatal("config not found")
	}

	addr := os.Getenv("HTTPAddr")
	if addr == "" {
		log.Fatal("http address not set")
	}

	databaseURL := os.Getenv("databaseurl")
	if databaseURL == "" {
		log.Fatal("wrong database url")
	}

	return &Config{
		HTTPAddr: addr,
		DatabaseURL: databaseURL,
	}
}
