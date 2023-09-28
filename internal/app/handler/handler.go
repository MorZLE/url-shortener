package handler

import (
	"fmt"
	"github.com/MorZLE/url-shortener/internal/app/service"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

func NewHandler(lg *service.AppService) *AppHandler {
	return &AppHandler{logic: lg}
}

type InterfaceAppHandler interface {
	RunServer()
	URLShortener(w http.ResponseWriter, r *http.Request)
	URLGetID(w http.ResponseWriter, r *http.Request)
}

type AppHandler struct {
	InterfaceAppHandler
	logic *service.AppService
}

func (h *AppHandler) RunServer() {

	router := mux.NewRouter()

	router.HandleFunc(`/`, h.URLShortener).Methods(http.MethodPost)
	router.HandleFunc(`/`, h.URLGetID).Methods(http.MethodGet)

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func (h *AppHandler) URLShortener(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	log.Println("Получен url:", string(body))
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(r.Body)

	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	shortURL, err := h.logic.URLShorter(string(body))

	if err != nil {
		http.Error(w, "Error shorting URL", http.StatusBadRequest)
		return
	}
	// Set the response content type
	w.Header().Set("Content-Type", "text/plain")

	w.WriteHeader(http.StatusCreated)
	log.Println("Created short URL:", shortURL)

	_, err = fmt.Fprint(w, shortURL)

	if err != nil {
		return
	}

}

func (h *AppHandler) URLGetID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	log.Println("id:", id)

	url, err := h.logic.URLGetID(id)
	if err != nil {
		log.Println("Error getting URL:", err)
		http.Error(w, "Error getting URL", http.StatusBadRequest)
		return
	}

	log.Println("отправлен url:", url)

	w.Header().Set("Location", url)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
