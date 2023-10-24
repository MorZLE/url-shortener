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
	ServerAddr  string
	BaseURL     string
	Memory      string
	DatabaseDsn string
}

func ParseFlags(p *Config) *Config {

	flag.StringVar(&p.ServerAddr, "a", ":8080", "address and port to run server")
	flag.StringVar(&p.BaseURL, "b", "http://127.0.0.1:8080", "address shortURLer")
	flag.StringVar(&p.Memory, "f", "", "save memory")
	flag.StringVar(&p.DatabaseDsn, "d", "", "database dsn")

	flag.Parse()

	if serverAddr := os.Getenv("SERVER_ADDRESS"); serverAddr != "" {
		p.ServerAddr = serverAddr
	}

	if baseURL := os.Getenv("BASE_URL"); baseURL != "" {
		p.BaseURL = baseURL
	}

	if memory := os.Getenv("FILE_STORAGE_PATH"); memory != "" {
		p.Memory = memory
	}

	if databaseDsn := os.Getenv("DATABASE_DSN"); databaseDsn != "" {
		p.DatabaseDsn = databaseDsn
	}

	log.Println("server", p.ServerAddr)
	log.Println("memory", p.Memory)
	log.Println("database", p.DatabaseDsn)
	return p
}
