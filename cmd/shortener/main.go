package main

import (
	"github.com/MorZLE/url-shortener/internal/app/handler"
	"github.com/MorZLE/url-shortener/internal/app/service"
	"github.com/MorZLE/url-shortener/internal/app/storage"
	"github.com/MorZLE/url-shortener/internal/config"
)

func main() {

	cnf := config.NewConfig()

	st := storage.NewStorage()
	lgc := service.NewService(&st, cnf)
	hdr := handler.NewHandler(&lgc, cnf)

	hdr.RunServer()

}
