package storage

import (
	"fmt"
	"github.com/MorZLE/url-shortener/internal/config"
	"github.com/MorZLE/url-shortener/internal/domains"
)

func NewStorage(cnf *config.Config) (domains.Storage, error) {
	if cnf.DatabaseDsn != "" {
		db, err := NewDB(cnf)
		if err != nil {
			return nil, fmt.Errorf("Don`t create db %w", err)
		}
		return &db, nil
	}

	if cnf.Memory != "" {
		writer, err := NewWriter(cnf.Memory)
		if err != nil {
			return nil, fmt.Errorf("Don`t create file %w", err)
		}

		reader, err := NewReader(cnf.Memory)
		if err != nil {
			return nil, fmt.Errorf("Don`t create file for read %w", err)
		}
		m, err := reader.ReadURL()
		if err != nil {
			return nil, fmt.Errorf("Don`t read file %w", err)
		}
		return &Storage{m: m, wr: writer}, nil
	}

	return &Storage{m: make(map[string]string)}, nil
}
