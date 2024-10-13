package main

import (
	"log"
	"net/http"
	"url/api"
	"url/storage"
)

func main() {
	store := storage.GetInstance()

	http.HandleFunc("/shorten", api.ShortenURL(store))
	http.HandleFunc("/redirect", api.RedirectURL(store))

	log.Fatal(http.ListenAndServe(":1000", nil))
}
