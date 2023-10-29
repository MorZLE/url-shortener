package main

import (
	"github.com/MorZLE/url-shortener/internal/app/handler"
	"github.com/MorZLE/url-shortener/internal/app/logger"
	"github.com/MorZLE/url-shortener/internal/app/service"
	"github.com/MorZLE/url-shortener/internal/app/storage"
	"github.com/MorZLE/url-shortener/internal/config"
	"log"
)

func main() {

	logger.Initialize()
	cnf := config.NewConfig()

	st, err := storage.NewStorage(cnf)
	if err != nil {
		log.Fatal(err)
	}
	lgc := service.NewService(st, cnf)
	hdr := handler.NewHandler(&lgc, cnf)

	hdr.RunServer()
	defer st.Close()
}
