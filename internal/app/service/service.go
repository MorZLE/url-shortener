package service

import (
	"github.com/MorZLE/url-shortener/internal/app/logger"
	"github.com/MorZLE/url-shortener/internal/config"
	"github.com/MorZLE/url-shortener/internal/domains"
	"github.com/speps/go-hashids"
	"time"
)

func NewService(s domains.StorageInterface, cnf *config.Config) Service {
	return Service{
		Storage: s,
		Cnf:     *cnf,
	}
}

type Service struct {
	Storage domains.StorageInterface
	Cnf     config.Config
}

func (s *Service) URLShorter(url string) (string, error) {
	hd := hashids.NewData()
	hd.MinLength = 6
	h, _ := hashids.NewWithData(hd)
	now := time.Now()
	shortURL, _ := h.Encode([]int{int(now.Unix())})
	err := s.Storage.Set(shortURL, url)
	shortURL = s.Cnf.BaseURL + "/" + shortURL
	if err != nil {
		logger.Error("Ключ short URL занят:", err)
		_, _ = s.URLShorter(url)
		//return "", err
	}
	logger.ShortURL(shortURL)
	return shortURL, nil
}

func (s *Service) URLGetID(url string) (string, error) {
	val, err := s.Storage.Get(url)
	if err != nil {
		return "", err
	}

	return val, nil

}
