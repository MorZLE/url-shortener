package storage

import (
	"github.com/MorZLE/url-shortener/internal/consterr"
)

func NewStorage() AppStorage {
	return AppStorage{M: make(map[string]string)}
}

//go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name=AppStorageInterface
type AppStorageInterface interface {
	Set(key string, value string) error
	Get(key string) (string, error)
}

type AppStorage struct {
	AppStorageInterface
	M map[string]string
}

func (s *AppStorage) Set(key string, value string) error {
	if s.M[key] != "" {
		return consterr.ErrKeyBusy
	}
	s.M[key] = value
	return nil
}

func (s *AppStorage) Get(key string) (string, error) {
	if s.M[key] != "" {
		return s.M[key], nil
	}
	return "", consterr.ErrNotFound
}
