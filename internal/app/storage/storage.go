package storage

import (
	"fmt"
	"github.com/MorZLE/url-shortener/internal/config"
	"github.com/MorZLE/url-shortener/internal/domains"
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
