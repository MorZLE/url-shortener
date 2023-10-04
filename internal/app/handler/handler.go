package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/MorZLE/url-shortener/internal/app/logger"
	"github.com/MorZLE/url-shortener/internal/config"
	"github.com/MorZLE/url-shortener/internal/constjson"
	"github.com/MorZLE/url-shortener/internal/domains"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

func NewHandler(lg domains.ServiceInterface, cnf *config.Config) Handler {
	return Handler{logic: lg, cnf: *cnf}
}

type Handler struct {
	logic domains.ServiceInterface
	cnf   config.Config
}

func (h *Handler) RunServer() {
	logger.Initialize()
	router := mux.NewRouter()
	router.Handle(`/`, logger.RequestLogger(h.URLShortener)).Methods(http.MethodPost)
	router.Handle(`/api/shorten,`, logger.RequestLogger(h.JSONURLShort)).Methods(http.MethodPost)
	router.Handle(`/api/shorten,`, logger.RequestLogger(h.JSONURLGetID)).Methods(http.MethodGet)
	router.Handle(`/{id}`, logger.RequestLogger(h.URLGetID)).Methods(http.MethodGet)

	log.Println("Run server ", h.cnf.ServerAddr)

	log.Fatal(http.ListenAndServe(h.cnf.ServerAddr, router))
}

func (h *Handler) JSONURLShort(w http.ResponseWriter, r *http.Request) {
	var url constjson.URLLong
	var buf bytes.Buffer

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(buf.Bytes(), &url); err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	longURL := url.URL
	shortURL, err := h.logic.URLShorter(longURL)
	if err != nil {
		http.Error(w, "Error shorting URL", http.StatusBadRequest)
		return
	}
	h.ResponseValueJSON(w, constjson.URLShort{Result: shortURL})

}
func (h *Handler) ResponseValueJSON(res http.ResponseWriter, obj constjson.URLShort) {
	resp, err := json.Marshal(&obj)
	if err != nil {
		http.Error(res, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	res.Header().Set("Accept", "application/json")
	res.Header().Set("Content-Type", "application/json")

	res.WriteHeader(http.StatusOK)

	_, err = res.Write(resp)
	if err != nil {
		return
	}

}

func (h *Handler) URLShortener(w http.ResponseWriter, r *http.Request) {

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

	_, err = fmt.Fprint(w, shortURL)
	if err != nil {
		return
	}

}

func (h *Handler) URLGetID(w http.ResponseWriter, r *http.Request) {

	url, err := h.logic.URLGetID(mux.Vars(r)["id"])
	if err != nil {
		log.Println("Error getting URL:", err)
		http.Error(w, "Error getting URL", http.StatusBadRequest)
		return
	}

	log.Println("отправлен url:", url)

	w.Header().Set("Location", url)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func (h *Handler) JSONURLGetID(w http.ResponseWriter, r *http.Request) {

	var url constjson.URLLong
	var buf bytes.Buffer

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(buf.Bytes(), &url); err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	longURL := url.URL
	shortURL, err := h.logic.URLGetID(longURL[len(longURL)-8:])
	if err != nil {
		http.Error(w, "Error shorting URL", http.StatusBadRequest)
		return
	}
	w.Header().Set("Location", shortURL)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
