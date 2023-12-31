package storage

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/MorZLE/url-shortener/internal/models"
	"io"
	"os"
	"path/filepath"
)

type Writer struct {
	file    *os.File
	encoder *json.Encoder
}

func NewWriter(fileName string) (*Writer, error) {
	err := os.MkdirAll(filepath.Dir(fileName), os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("can't create directory: %w", err)
	}
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, fmt.Errorf("can't open directory: %w", err)
	}

	return &Writer{
		file:    file,
		encoder: json.NewEncoder(file),
	}, nil
}

func (p *Writer) WriteURL(url *models.URLFile) error {
	data, err := json.Marshal(&url)
	if err != nil {
		return err
	}
	// добавляем перенос строки
	data = append(data, '\n')

	_, err = p.file.Write(data)
	return err
}

func (p *Writer) Close() error {
	return p.file.Close()
}

type Reader struct {
	file   *os.File // файл для чтения
	reader *bufio.Reader
}

func NewReader(filename string) (*Reader, error) {
	err := os.MkdirAll(filepath.Dir(filename), os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("can't create directory: %w", err)
	}

	file, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, fmt.Errorf("can't open directory: %w", err)
	}

	return &Reader{
		file: file,
		// создаём новый Reader
		reader: bufio.NewReader(file),
	}, nil
}

func (c *Reader) Close() error {
	// закрываем файл
	return c.file.Close()
}

func (c *Reader) ReadURL() (map[string]map[string]string, error) {
	m := make(map[string]map[string]string)
	defer c.Close()

	for {
		var url models.URLFile
		nextData, err := c.reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("can't read file: %w", err)
		}

		err = json.Unmarshal(nextData, &url)
		if err != nil {
			return nil, err
		}
		m[url.UserID][url.ShortURL] = url.OriginalURL
	}

	return m, nil
}
