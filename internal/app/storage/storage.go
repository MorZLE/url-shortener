package storage

import (
	"github.com/MorZLE/url-shortener/internal/config"
	"github.com/MorZLE/url-shortener/internal/consterr"
	"github.com/MorZLE/url-shortener/internal/constjson"
	"log"
)

func NewStorage(cnf *config.Config) Storage {
	if cnf.Memory == "" {
		writer, err := NewWriter(cnf.Memory)
		if err != nil {
			log.Fatal("Не удалось создать файл для хранения")
		}

		reader, err := NewReader(cnf.Memory)
		if err != nil {
			log.Fatal("Не удалось создать файл для хранения")
		}
		m, err := reader.ReadURL()
		if err != nil {
			log.Fatal("Не удалось прочитать файл")
		}
		return Storage{M: m, Writer: writer}
	}
	return Storage{M: make(map[string]string)}
}

type Storage struct {
	M      map[string]string
	count  int
	Writer *Writer
}

func (s *Storage) Set(key string, value string) error {
	if s.M[key] != "" {
		return consterr.ErrKeyBusy
	}
	if s.Writer != nil {
		err := s.Writer.WriteURL(&constjson.URLFile{ShortURL: key, OriginalURL: value})
		if err != nil {
			return err
		}
	}
	s.M[key] = value
	s.count++
	return nil
}

func (s *Storage) Get(key string) (string, error) {
	if s.M[key] != "" {
		return s.M[key], nil
	}
	return "", consterr.ErrNotFound
}

func (s *Storage) Count() int {
	return s.count
}
