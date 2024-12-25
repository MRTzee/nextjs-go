package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mrtzee/nextjs-go/internal/config"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "text/plain")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	log.Printf("Received %s request from %s", r.Method, r.RemoteAddr)
	w.Write([]byte("Hello, world!"))
}

func main() {
	http.HandleFunc("/", handler)
	viperConfig := config.LoadConfig()
	log := config.NewLogger(viperConfig)
	db := config.NewDatabase(viperConfig, log)
	app := config.NewFiber(viperConfig)

	cfg := &config.AppConfig{
		DB:     db,
		App:    app,
		Log:    log,
		Config: viperConfig,
	}

	cfg.Run()

	webPort := viperConfig.GetInt("APP_PORT")
	err := app.Listen(fmt.Sprintf(":%d", webPort))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
