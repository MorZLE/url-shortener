package storage

import (
	"fmt"
	"github.com/MorZLE/url-shortener/internal/config"
	"github.com/MorZLE/url-shortener/internal/consts"
	"github.com/MorZLE/url-shortener/internal/domains"
	"github.com/MorZLE/url-shortener/internal/models"
)

func NewStorage(cnf *config.Config) (domains.StorageInterface, error) {
	if cnf.DatabaseDsn != "" {
		db, err := NewDB(cnf)
		if err != nil {
			return nil, fmt.Errorf("не удалось создать базу данных %w", err)
		}
		return &db, nil
	}

	if cnf.Memory != "" {
		writer, err := NewWriter(cnf.Memory)
		if err != nil {
			return nil, fmt.Errorf("не удалось создать файл для хранения %w", err)
		}

		reader, err := NewReader(cnf.Memory)
		if err != nil {
			return nil, fmt.Errorf("не удалось создать файл для чтения %w", err)
		}
		m, err := reader.ReadURL()
		if err != nil {
			return nil, fmt.Errorf("не удалось прочитать файл %w", err)
		}
		return &Storage{M: m, Writer: writer}, nil
	}

	return &Storage{M: make(map[string]string)}, nil
}

type Storage struct {
	M      map[string]string
	Writer *Writer
}

func (s *Storage) Ping() error {
	return nil
}

func (s *Storage) Set(key string, value string) error {
	if s.M[key] != "" {
		return consts.ErrKeyBusy
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

func (s *Storage) Get(key string) (string, error) {
	if s.M[key] != "" {
		return s.M[key], nil
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
