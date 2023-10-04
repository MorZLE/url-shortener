package storage

import (
	"github.com/MorZLE/url-shortener/internal/consterr"
)

func NewStorage() Storage {
	return Storage{M: make(map[string]string)}
}

type Storage struct {
	M map[string]string
}

func (s *Storage) Set(key string, value string) error {
	if s.M[key] != "" {
		return consterr.ErrKeyBusy
	}
	s.M[key] = value
	return nil
}

func (s *Storage) Get(key string) (string, error) {
	if s.M[key] != "" {
		return s.M[key], nil
	}
	return "", consterr.ErrNotFound
}
