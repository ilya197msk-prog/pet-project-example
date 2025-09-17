package main

import (
	"log"
	"net/http"
	"your_username/go-url-shortener/internal/app/handler"
	"your_username/go-url-shortener/internal/app/repository"
)

func main() {
	// Репозиторий (in-memory)
	repo := repository.NewInMemoryRepository()
	
	// Обработчик
	h := handler.NewHandler(repo)
	
	// Роутинг
	http.HandleFunc("/shorten", h.ShortenHandler)
	http.HandleFunc("/", h.RedirectHandler)

	log.Println("Server is running on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
