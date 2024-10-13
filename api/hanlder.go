package api

import (
	"encoding/json"
	"log"
	"net/http"
	"url/shortener"
	"url/storage"
)

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	Shortend string `json:"shortend"`
}

func ShortenURL(store storage.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
			return
		}

		var req ShortenRequest
		json.NewDecoder(r.Body).Decode(&req)

		if req.URL == "" {
			http.Error(w, "URL not found", http.StatusNotFound)
			return
		}

		shortCode := shortener.Shorten(req.URL)
		log.Println(req.URL, "->", shortCode)

		store.Save(shortCode, req.URL)
		json.NewEncoder(w).Encode(ShortenResponse{Shortend: shortCode})
	}
}

func RedirectURL(store storage.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		shortCode := r.URL.Query().Get("url")

		OriginalURL, err := store.Get(shortCode)
		if err != nil {
			http.Error(w, "URL not found", http.StatusNotFound)
			return
		}

		http.Redirect(w, r, OriginalURL, http.StatusFound)
	}
}
