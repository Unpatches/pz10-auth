package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"

	router "example.com/pz10-auth/internal/http"
	"example.com/pz10-auth/internal/platform/config"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found, using OS env only")
	}

	cfg := config.Load()
	mux := router.Build(cfg) // см. следующий шаг
	log.Println("listening on", cfg.Port)
	log.Fatal(http.ListenAndServe(cfg.Port, mux))
}
