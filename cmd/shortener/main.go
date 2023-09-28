package main

import (
	"github.com/MorZLE/url-shortener/internal/app/handler"
	"github.com/MorZLE/url-shortener/internal/app/logic"
	"github.com/MorZLE/url-shortener/internal/app/storage"
	"net/http"
)

func main() {

	st := storage.NewStorage()
	lgc := logic.NewLogic(st)
	hdr := handler.NewHandler(lgc)
	http.ListenAndServe(":8080", hdr)
}
