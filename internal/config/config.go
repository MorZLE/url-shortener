package config

import (
	"flag"
	"log"
	"os"
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

const BaseFile = "tmp/short-url-db.json"

func ParseFlags(p *Config) *Config {

	flag.StringVar(&p.ServerAddr, "a", ":8080", "address and port to run server")
	flag.StringVar(&p.BaseURL, "b", "http://127.0.0.1:8080", "address shortURLer")
	flag.StringVar(&p.Memory, "f", BaseFile, "save memory")

	flag.Parse()

	if serverAddr := os.Getenv("SERVER_ADDRESS"); serverAddr != "" {
		p.ServerAddr = serverAddr
	}

	if baseURL := os.Getenv("BASE_URL"); baseURL != "" {
		p.BaseURL = baseURL
	}
	if memory := os.Getenv("FILE_STORAGE_PATH "); memory != "" {
		p.Memory = memory
	}
	log.Println("memory", p.Memory)
	return p
}
