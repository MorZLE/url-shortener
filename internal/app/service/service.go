package service

import (
	"errors"
	"github.com/MorZLE/url-shortener/internal/app/logger"
	"github.com/MorZLE/url-shortener/internal/config"
	"github.com/MorZLE/url-shortener/internal/consts"
	"github.com/MorZLE/url-shortener/internal/domains"
	"github.com/MorZLE/url-shortener/internal/models"
	"github.com/speps/go-hashids"
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

func (s *Service) URLsShorter(data []models.BatchSet) ([]models.BatchGet, error) {

	var shUrls []models.BatchGet

	shURStorage := make(map[string]string)
	for i, url := range data {
		if url.OriginalURL == "" {
			continue
		}
		ln := s.Storage.Count() + i + 1
		hd := hashids.NewData()
		h, err := hashids.NewWithData(hd)
		if err != nil {
			logger.Error("Ошибка NewWithData:", err)
			return nil, err
		}
		shortURL, err := h.Encode([]int{ln})
		if err != nil {
			logger.Error("Ошибка Encode:", err)
			return nil, err
		}
		shURStorage[shortURL] = url.OriginalURL
		shortURL = s.Cnf.BaseURL + "/" + shortURL
		shUrls = append(shUrls, models.BatchGet{
			CorrelationID: url.CorrelationID,
			ShortURL:      shortURL,
		})
	}

	err := s.Storage.SetBatch(shURStorage)
	if err != nil {
		logger.Error("Ключ short URL занят:", err)
		return nil, err
	}

	return shUrls, nil
}

func (s *Service) URLShorter(url string) (string, error) {
	hd := hashids.NewData()
	h, err := hashids.NewWithData(hd)
	if err != nil {
		logger.Error("Ошибка NewWithData:", err)
		return "", err
	}
	shortURL, err := h.Encode([]int{s.Storage.Count()})
	if err != nil {
		logger.Error("Ошибка Encode:", err)
		return "", err
	}
	dubleurl, err := s.Storage.Set(shortURL, url)
	if err != nil {
		if errors.Is(err, consts.ErrDuplicateURL) {
			return dubleurl, consts.ErrDuplicateURL
		}
		if errors.Is(err, consts.ErrKeyBusy) {
			logger.Error("Ключ short URL занят:", err)
			return "", err
		}
	}
	shortURL = s.Cnf.BaseURL + "/" + shortURL
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

func (s *Service) CheckPing() error {
	return s.Storage.Ping()
}
