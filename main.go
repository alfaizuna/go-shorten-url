package main

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"log"
	"net/http"
	"os"

	"github.com/redis/go-redis/v9"
)

var (
	urlMap      = make(map[string]string)
	redisClient *redis.Client
	ctx         = context.Background()
)

func initRedis() {
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		return // Redis opsional
	}
	redisClient = redis.NewClient(&redis.Options{
		Addr: addr,
	})
}

func saveURL(shortPath, longURL string) {
	if redisClient != nil {
		redisClient.Set(ctx, shortPath, longURL, 0)
	} else {
		urlMap[shortPath] = longURL
	}
}

func getURL(shortPath string) (string, bool) {
	if redisClient != nil {
		val, err := redisClient.Get(ctx, shortPath).Result()
		if err == nil {
			return val, true
		}
		return "", false
	}
	val, ok := urlMap[shortPath]
	return val, ok
}

func encodeToShortPath(longURL string) string {
	hash := sha256.Sum256([]byte(longURL))
	// Ambil 6 karakter base64 agar pendek
	return base64.URLEncoding.EncodeToString(hash[:])[:6]
}

func shortenURLHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	longURL := r.FormValue("url")
	if longURL == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing url parameter"))
		return
	}

	// Gunakan encoding untuk shortPath
	shortPath := encodeToShortPath(longURL)
	shortURL := "http://localhost:8080/" + shortPath

	saveURL(shortPath, longURL)

	w.Write([]byte(shortURL))
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	shortPath := r.URL.Path[1:]
	longURL, ok := getURL(shortPath)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Short URL not found"))
		return
	}
	// Redirect ke URL asli
	http.Redirect(w, r, longURL, http.StatusFound)
}

func main() {
	initRedis()
	mux := http.NewServeMux()
	mux.HandleFunc("/shorten", shortenURLHandler)
	mux.HandleFunc("/", redirectHandler)
	// Routing akan ditambahkan di sini

	log.Println("Server berjalan di :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
