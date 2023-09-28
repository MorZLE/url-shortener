package handler

import (
	"fmt"
	"github.com/MorZLE/url-shortener/internal/app/service"
	"github.com/MorZLE/url-shortener/internal/config"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

func NewHandler(lg service.InterfaceAppService, cnf *config.Config) AppHandler {
	return AppHandler{logic: lg, cnf: *cnf}
}

//go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name=InterfaceAppHandler
type InterfaceAppHandler interface {
	RunServer()
	URLShortener(w http.ResponseWriter, r *http.Request)
	URLGetID(w http.ResponseWriter, r *http.Request)
}

type AppHandler struct {
	InterfaceAppHandler
	logic service.InterfaceAppService
	cnf   config.Config
}

func (h *AppHandler) RunServer() {

	router := mux.NewRouter()

	router.HandleFunc(`/`, h.URLShortener).Methods(http.MethodPost)
	router.HandleFunc(`/{id}`, h.URLGetID).Methods(http.MethodGet)

	log.Println("Run server ", h.cnf.FlagAddrReq)
	log.Fatal(http.ListenAndServe(h.cnf.FlagAddrReq, router))
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

	w.Header().Set("Content-Type", "text/plain")

	w.WriteHeader(http.StatusCreated)
	log.Println("Created short URL:", shortURL)

	_, err = fmt.Fprint(w, shortURL)

	if err != nil {
		return
	}

}

func (h *AppHandler) URLGetID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	uri := fmt.Sprintf("http://%s/%s", h.cnf.FlagAddrReq, id)

	log.Println("uri:", uri)

	url, err := h.logic.URLGetID(uri)
	if err != nil {
		log.Println("Error getting URL:", err)
		http.Error(w, "Error getting URL", http.StatusBadRequest)
		return
	}

	log.Println("отправлен url:", url)

	w.Header().Set("Location", url)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
