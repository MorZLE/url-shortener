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
	UrlShorter(url string) (string, error)
	UrlGetID(url string) (string, error)
	GenerateShortUrl() string
}

type AppService struct {
	InterfaceAppService
	storage storage.AppStorage
}

func (s *AppService) UrlShorter(url string) (string, error) {
	var shortUrl string

	for {
		shortUrl := s.GenerateShortUrl()
		err := s.storage.Set(shortUrl, url)
		if err == nil {
			break
		}
	}
	return shortUrl, nil
}

func (s *AppService) GenerateShortUrl() string {
	rand.Seed(time.Now().UnixNano())

	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	result := make([]byte, 8)
	for i := 0; i < 8; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}

	return string(result)
}

func (s *AppService) UrlGetID(url string) (string, error) {
	val, err := s.storage.Get(url)
	if err != nil {
		return "", err
	}

	return val, nil

}
