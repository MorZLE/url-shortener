package config

import (
	"flag"
	"os"
)

func NewConfig() *Config {
	cnf := &Config{}
	return ParseFlags(cnf)
}

type Config struct {
	ServerAddr string
	BaseURL    string
}

func ParseFlags(p *Config) *Config {

	flag.StringVar(&p.ServerAddr, "a", ":8080", "address and port to run server")
	flag.StringVar(&p.BaseURL, "b", "http://127.0.0.1:8080", "address shortURLer")

	flag.Parse()

	if serverAddr := os.Getenv("SERVER_ADDRESS"); serverAddr != "" {
		p.ServerAddr = serverAddr
	}

	if baseURL := os.Getenv("BASE_URL"); baseURL != "" {
		p.BaseURL = baseURL
	}

	return p
}
