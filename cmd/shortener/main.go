package main

import (
	"github.com/MorZLE/url-shortener/internal/app/handler"
	"github.com/MorZLE/url-shortener/internal/app/service"
	"github.com/MorZLE/url-shortener/internal/app/storage"
)

func main() {

	st := storage.NewStorage()
	lgc := service.NewService(st)
	hdr := handler.NewHandler(lgc)

	hdr.RunServer()

}
