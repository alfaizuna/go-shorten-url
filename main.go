package main

import (
	"fmt"
	"log"
	"net/http"
	"urlshortener/config"
	"urlshortener/handlers"
	"urlshortener/storage"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Inisialisasi storage (Redis jika tersedia, jika tidak pakai Map)
	var urlStorage storage.URLStorage

	redisStorage := storage.NewRedisStorage(cfg)
	if redisStorage != nil {
		urlStorage = redisStorage
		log.Println("Menggunakan Redis storage")
	} else {
		urlStorage = storage.NewMapStorage()
		log.Println("Menggunakan Map storage")
	}

	// Inisialisasi handler
	urlHandler := handlers.NewURLHandler(urlStorage, cfg)

	// Setup routing
	mux := http.NewServeMux()
	mux.HandleFunc("/shorten", urlHandler.ShortenURL)
	mux.HandleFunc("/", urlHandler.Redirect)

	serverAddr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("Server berjalan di %s", serverAddr)
	if err := http.ListenAndServe(serverAddr, mux); err != nil {
		log.Fatal(err)
	}
}
