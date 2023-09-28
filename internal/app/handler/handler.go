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

	router.HandleFunc(`/}`, h.URLShortener).Methods(http.MethodPost)
	router.HandleFunc(`/{id}`, h.URLGetID).Methods(http.MethodGet)

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func (h *AppHandler) URLShortener(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

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

	// Echo the URL string in the response
	_, err = fmt.Fprint(w, shortURL)
	if err != nil {
		return
	}

}

func (h *AppHandler) URLGetID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	url, err := h.logic.URLGetID(id)
	if err != nil {
		http.Error(w, "Error getting URL", http.StatusBadRequest)
		return
	}
	w.Header().Set("Location", url)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
