package storage

import (
	"errors"
	"sync"
)

type Store interface {
	Save(shortCode, LongURL string)
	Get(shortCode string) (string, error)
}

type InMemoryStorage struct {
	urls map[string]string
	mu   sync.RWMutex
}

var (
	instance *InMemoryStorage
	once     sync.Once
)

func GetInstance() *InMemoryStorage {
	once.Do(func() {
		instance = &InMemoryStorage{
			urls: make(map[string]string),
		}
	})
	return instance
}

func (s *InMemoryStorage) Save(shortCode, LongURL string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.urls[shortCode] = LongURL
}

func (s *InMemoryStorage) Get(shortCode string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if LongURL, ok := s.urls[shortCode]; ok {
		return LongURL, nil
	}

	return "", ErrNotFound
}

var ErrNotFound = errors.New("NotFound")
