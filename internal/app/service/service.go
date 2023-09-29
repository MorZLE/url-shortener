package service

import (
	"github.com/MorZLE/url-shortener/internal/app/storage"
	"github.com/MorZLE/url-shortener/internal/config"
	"math/rand"
	"time"
)

func NewService(s storage.AppStorageInterface, cnf *config.Config) AppService {
	return AppService{Storage: s, Cnf: *cnf}
}

//go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name=InterfaceAppService
type InterfaceAppService interface {
	URLShorter(url string) (string, error)
	URLGetID(url string) (string, error)
	GenerateShortURL() string
}

type AppService struct {
	InterfaceAppService
	Storage storage.AppStorageInterface
	Cnf     config.Config
}

func (s *AppService) URLShorter(url string) (string, error) {
	for {
		shortURL := s.GenerateShortURL()
		shortURL = s.Cnf.BaseURL + "/" + shortURL
		err := s.Storage.Set(shortURL, url)
		if err == nil {
			return shortURL, nil
		}
	}
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
	val, err := s.Storage.Get(url)
	if err != nil {
		return "", err
	}

	return val, nil

}
