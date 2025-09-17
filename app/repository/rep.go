package repository

import (
	"sync"
	"./go-url-shortener/internal/app/model"
)

type Repository interface {
	Save(url *model.URL) error
	Find(shortURL string) (*model.URL, error)
}

type inMemoryRepository struct {
	urls map[string]string
	mu   sync.RWMutex
}

func NewInMemoryRepository() Repository {
	return &inMemoryRepository{
		urls: make(map[string]string),
	}
}

func (r *inMemoryRepository) Save(url *model.URL) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.urls[url.ShortURL] = url.OriginalURL
	return nil
}

func (r *inMemoryRepository) Find(shortURL string) (*model.URL, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	originalURL, found := r.urls[shortURL]
	if !found {
		return nil, nil //err
	}
	return &model.URL{ShortURL: shortURL, OriginalURL: originalURL}, nil
}
