package service

import (
	"github.com/MorZLE/url-shortener/internal/app/storage"
	"math/rand"
	"time"
)

func NewService(s *storage.AppStorage) *AppService {
	return &AppService{storage: *s}
}

type InterfaceAppService interface {
	URLShorter(url string) (string, error)
	URLGetID(url string) (string, error)
	GenerateShortURL() string
}

type AppService struct {
	InterfaceAppService
	storage storage.AppStorage
}

func (s *AppService) URLShorter(url string) (string, error) {
	var shortURL string

	for {
		shortURL := s.GenerateShortURL()
		err := s.storage.Set(shortURL, url)
		if err == nil {
			break
		}
	}
	return shortURL, nil
}

func (s *AppService) GenerateShortURL() string {
	rand.NewSource(time.Now().UnixNano())

	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	result := make([]byte, 8)
	for i := 0; i < 8; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}

	return string(result)

}

func (s *AppService) URLGetID(url string) (string, error) {
	val, err := s.storage.Get(url)
	if err != nil {
		return "", err
	}

	return val, nil

}
