package main

import (
	"log"
	"net/http"
	"./go-url-shortener/internal/app/handler"
	"./go-url-shortener/internal/app/repository"
)

func main() {
	repo := repository.NewInMemoryRepository()
	
	h := handler.NewHandler(repo)
	
	http.HandleFunc("/shorten", h.ShortenHandler)
	http.HandleFunc("/", h.RedirectHandler)

	log.Println("Server is running on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
