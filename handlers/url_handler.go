package handlers

import (
	"net/http"
	"urlshortener/config"
	"urlshortener/storage"
	"urlshortener/utils"
)

type URLHandler struct {
	storage storage.URLStorage
	config  *config.Config
}

func NewURLHandler(storage storage.URLStorage, cfg *config.Config) *URLHandler {
	return &URLHandler{
		storage: storage,
		config:  cfg,
	}
}

func (h *URLHandler) ShortenURL(w http.ResponseWriter, r *http.Request) {
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

	shortPath := utils.EncodeToShortPath(longURL)
	shortURL := h.config.BaseURL + "/" + shortPath

	h.storage.SaveURL(shortPath, longURL)

	w.Write([]byte(shortURL))
}

func (h *URLHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	shortPath := r.URL.Path[1:]
	longURL, ok := h.storage.GetURL(shortPath)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Short URL not found"))
		return
	}
	http.Redirect(w, r, longURL, http.StatusFound)
}
