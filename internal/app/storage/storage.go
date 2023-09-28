package storage

import (
	"github.com/MorZLE/url-shortener/internal/constErr"
)

func NewStorage() *AppStorage {
	return &AppStorage{m: make(map[string]string)}
}

type AppStorageInterface interface {
	Set(key string, value string) error
	Get(key string) (string, error)
}

type AppStorage struct {
	AppStorageInterface
	m map[string]string
}

func (s *AppStorage) Set(key string, value string) error {
	if s.m[key] != "" {
		return constErr.ErrKeyBusy
	}
	s.m[key] = value
	return nil
}

func (s *AppStorage) Get(key string) (string, error) {
	if s.m[key] != "" {
		return s.m[key], nil
	}
	return "", constErr.ErrNotFound
}
