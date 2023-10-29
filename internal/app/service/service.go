package service

import (
	"fmt"
	"github.com/MorZLE/url-shortener/internal/app/logger"
	"github.com/MorZLE/url-shortener/internal/config"
	"github.com/MorZLE/url-shortener/internal/consts"
	"github.com/MorZLE/url-shortener/internal/domains"
	"github.com/MorZLE/url-shortener/internal/models"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/speps/go-hashids"
	"sync/atomic"
)

func NewService(s domains.Storage, cnf *config.Config) Service {

	c := atomic.Uint64{}
	num, err := s.Count()
	if err != nil {
		logger.Error("Error Count:", err)
	}
	c.Add(uint64(num + 1))
	return Service{
		storage:      s,
		cnf:          *cnf,
		countStorage: &c,
	}
}

type Service struct {
	storage      domains.Storage
	cnf          config.Config
	countStorage *atomic.Uint64
}

func (s *Service) URLsShorter(id string, data []models.BatchSet) ([]models.BatchGet, error) {
	var shUrls []models.BatchGet
	shURStorage := make(map[string]string)
	for _, url := range data {
		if url.OriginalURL == "" {
			continue
		}
		shortURL, err := s.storage.GetDuplicate(url.OriginalURL)
		if err != nil {
			shortURL, err = s.Generate(int(s.countStorage.Load()))
			s.countStorage.Add(1)
			if err != nil {
				logger.Error("error Generate:", err)
				return nil, err
			}
			shURStorage[shortURL] = url.OriginalURL
		}

		shortURL = s.cnf.BaseURL + "/" + shortURL
		shUrls = append(shUrls, models.BatchGet{
			CorrelationID: url.CorrelationID,
			ShortURL:      shortURL,
		})
	}

	err := s.storage.SetBatch(id, shURStorage)
	if err != nil {
		logger.Error("key already exists:", err)
		return nil, err
	}

	return shUrls, nil
}

func (s *Service) URLShorter(id string, url string) (string, error) {
	shortURL, err := s.Generate(int(s.countStorage.Load()))
	s.countStorage.Add(1)
	if err != nil {
		logger.Error("error Generate:", err)
		return "", err
	}
	err = s.storage.Set(id, shortURL, url)
	if err != nil {
		if errors.Is(err, consts.ErrDuplicateURL) {
			shortURL, err = s.storage.GetDuplicate(url)
			if err != nil {
				logger.Error("error GetDuplicate:", err)
				return "", err
			}
			shortURL = s.cnf.BaseURL + "/" + shortURL
			return shortURL, consts.ErrDuplicateURL
		}
		return "", err
	}
	shortURL = s.cnf.BaseURL + "/" + shortURL
	logger.ShortURL(shortURL)
	return shortURL, nil
}

func (s *Service) Generate(num int) (string, error) {
	hd := hashids.NewData()
	h, err := hashids.NewWithData(hd)
	if err != nil {
		logger.Error("Error NewWithData:", err)
		return "", err
	}
	shortURL, err := h.Encode([]int{num})
	if err != nil {
		logger.Error("Error Encode:", err)
		return "", err
	}
	return shortURL, nil
}

func (s *Service) URLGetID(url string) (string, error) {
	val, err := s.storage.Get(url)
	if err != nil {
		logger.Error("error Get:", err)
		return "", err
	}

	return val, nil

}

func (s *Service) GetAllURLUsers(id string) ([]models.AllURLs, error) {
	var resurls []models.AllURLs
	urls, err := s.storage.GetAllURL(id)
	if err != nil {
		if errors.Is(err, consts.ErrDuplicateURL) {
			return nil, consts.ErrDuplicateURL
		}
		return nil, fmt.Errorf("error GetAllURL: %s", err)
	}
	for k, v := range urls {
		shortURL := s.cnf.BaseURL + "/" + k
		resurls = append(resurls, models.AllURLs{
			ShortURL:    shortURL,
			OriginalURL: v,
		})
	}
	return resurls, nil

}

func (s *Service) URLDelete(id string, url []string) {
	err := s.storage.UpdateDelete(id, url)
	if err != nil {
		logger.Error("error UpdateDelete:", err)
	}
}

func (s *Service) CheckPing() error {
	return s.storage.Ping()
}

func (s *Service) GenerateCookie() string {
	return uuid.New().String()
}
