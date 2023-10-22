package storage

import (
	"github.com/MorZLE/url-shortener/internal/consts"
	"github.com/MorZLE/url-shortener/internal/models"
)

type Storage struct {
	M      map[string]string
	Writer *Writer
}

func (s *Storage) Ping() error {
	return nil
}

func (s *Storage) Set(key string, value string) error {
	if s.M[key] != "" {
		return consts.ErrAlreadyExists
	}
	if s.Writer != nil {
		err := s.Writer.WriteURL(&models.URLFile{ShortURL: key, OriginalURL: value})
		if err != nil {
			return err
		}
	}
	s.M[key] = value
	return nil
}

func (s *Storage) SetBatch(m map[string]string) error {
	for key, value := range m {
		if s.M[key] != "" {
			return consts.ErrAlreadyExists
		}
		if s.Writer != nil {
			err := s.Writer.WriteURL(&models.URLFile{ShortURL: key, OriginalURL: value})
			if err != nil {
				return err
			}
		}
		s.M[key] = value
	}
	return nil
}

func (s *Storage) Get(key string) (string, error) {
	if v, ok := s.M[key]; ok {
		return v, nil
	}
	return "", consts.ErrNotFound
}

func (s *Storage) Count() int {
	return len(s.M)
}

func (s *Storage) Close() error {
	if s.Writer != nil {
		return s.Writer.Close()
	}
	return nil
}

func (s *Storage) GetDuplicate(longURL string) (string, error) {
	return "", nil
}
