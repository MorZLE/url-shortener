package storage

import (
	"github.com/MorZLE/url-shortener/internal/consts"
	"github.com/MorZLE/url-shortener/internal/models"
)

type Storage struct {
	m  map[string]string
	wr *Writer
}

func (s *Storage) Ping() error {
	return nil
}

func (s *Storage) Set(key string, value string) error {
	if s.m[key] != "" {
		return consts.ErrAlreadyExists
	}
	if s.wr != nil {
		err := s.wr.WriteURL(&models.URLFile{ShortURL: key, OriginalURL: value})
		if err != nil {
			return err
		}
	}
	s.m[key] = value
	return nil
}

func (s *Storage) SetBatch(m map[string]string) error {
	for key, value := range m {
		if s.m[key] != "" {
			return consts.ErrAlreadyExists
		}
		if s.wr != nil {
			err := s.wr.WriteURL(&models.URLFile{ShortURL: key, OriginalURL: value})
			if err != nil {
				return err
			}
		}
		s.m[key] = value
	}
	return nil
}

func (s *Storage) Get(key string) (string, error) {
	if v, ok := s.m[key]; ok {
		return v, nil
	}
	return "", consts.ErrNotFound
}

func (s *Storage) Count() (int, error) {
	return len(s.m), nil
}

func (s *Storage) Close() error {
	if s.wr != nil {
		return s.wr.Close()
	}
	return nil
}

func (s *Storage) GetDuplicate(longURL string) (string, error) {
	return "", nil
}
