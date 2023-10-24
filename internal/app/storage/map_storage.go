package storage

import (
	"github.com/MorZLE/url-shortener/internal/consts"
	"github.com/MorZLE/url-shortener/internal/models"
)

type Storage struct {
	m  map[string]map[string]string
	wr *Writer
}

func (s *Storage) Ping() error {
	return nil
}

func (s *Storage) Set(id, key, value string) error {
	if s.m[id] == nil {
		s.m[id] = make(map[string]string)
	}

	if s.m[id][key] != "" {
		return consts.ErrAlreadyExists
	}
	if s.wr != nil {
		err := s.wr.WriteURL(&models.URLFile{UserID: id, ShortURL: key, OriginalURL: value})
		if err != nil {
			return err
		}
	}
	s.m[id][key] = value
	return nil
}

func (s *Storage) SetBatch(id string, m map[string]string) error {
	for key, value := range m {
		if s.m[id][key] != "" {
			return consts.ErrAlreadyExists
		}
		if s.wr != nil {
			err := s.wr.WriteURL(&models.URLFile{ShortURL: key, OriginalURL: value})
			if err != nil {
				return err
			}
		}
		s.m[id][key] = value
	}
	return nil
}
func (s *Storage) GetAllURL(id string) (map[string]string, error) {
	if v, ok := s.m[id]; ok {
		return v, nil
	}

	return nil, consts.ErrAlreadyExists
}

func (s *Storage) Get(id string, key string) (string, error) {
	if v, ok := s.m[id][key]; ok {
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
