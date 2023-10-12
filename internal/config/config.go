package config

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func NewConfig() *Config {
	cnf := &Config{}
	return ParseFlags(cnf)
}

type Config struct {
	ServerAddr string
	BaseURL    string
	Memory     string
}

const BaseFile = "short-url-db.json"

func ParseFlags(p *Config) *Config {

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Ошибка при получении рабочей директории:", err)
	}
	// Получаем абсолютный путь до директории проекта
	projectDir := filepath.Dir(wd)

	fmt.Println("Директория проекта:", projectDir)

	flag.StringVar(&p.ServerAddr, "a", ":8080", "address and port to run server")
	flag.StringVar(&p.BaseURL, "b", "http://127.0.0.1:8080", "address shortURLer")
	flag.StringVar(&p.Memory, "f", "sfse", "save memory")

	flag.Parse()

	if serverAddr := os.Getenv("SERVER_ADDRESS"); serverAddr != "" {
		p.ServerAddr = serverAddr
	}

	if baseURL := os.Getenv("BASE_URL"); baseURL != "" {
		p.BaseURL = baseURL
	}
	if memory := os.Getenv("FILE_STORAGE_PATH "); memory != "" {
		p.Memory = memory[1:]
		return p
	}
	if p.Memory == "sfse" {
		p.Memory = ""
	} else {
		p.Memory = p.Memory[1:]
	}
	log.Println("memory", p.Memory)
	return p
}
