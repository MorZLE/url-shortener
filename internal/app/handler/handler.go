package handler

import (
	"encoding/json"
	gzipmilddle "github.com/MorZLE/url-shortener/internal/app/gzip"
	"github.com/MorZLE/url-shortener/internal/app/logger"
	"github.com/MorZLE/url-shortener/internal/config"
	"github.com/MorZLE/url-shortener/internal/consterr"
	"github.com/MorZLE/url-shortener/internal/constjson"
	"github.com/MorZLE/url-shortener/internal/domains"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
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

	router := gin.Default()
	router.Use(gzipmilddle.GzipMiddleware())
	router.Use(gzip.Gzip(gzip.BestSpeed))
	router.Use(gin.Logger())

	router.POST(`/`, h.URLShortener)
	router.POST(`/api/shorten`, h.JSONURLShort)
	router.GET(`/:id`, h.URLGetID)

	log.Fatal(router.Run(h.cnf.ServerAddr))

	log.Println("Run server ", h.cnf.ServerAddr)
}

func (h *Handler) JSONURLShort(c *gin.Context) {
	var url constjson.URLLong

	//b, err := UseGzip(c.Request.Body, c.Request.Header.Get("Content-Type"))
	//if err != nil {
	//	c.Error(err)
	//	c.AbortWithStatus(http.StatusInternalServerError)
	//
	//	return
	//}

	// b:= io.ByteCloser(c.Request.Body)

	if err := json.NewDecoder(c.Request.Body).Decode(&url); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	longURL := url.URL
	if longURL == "" {
		c.Error(consterr.ErrGetURL)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	shortURL, err := h.logic.URLShorter(longURL)
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	h.ResponseValueJSON(c, constjson.URLShort{Result: shortURL})

}

func (h *Handler) ResponseValueJSON(c *gin.Context, obj constjson.URLShort) {
	resp, err := json.Marshal(&obj)
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.Header("Content-Type", "application/json")
	c.Status(http.StatusCreated)

	c.Writer.Write(resp)

}

func (h *Handler) URLShortener(c *gin.Context) {

	//body, err := UseGzip(c.Request.Body, c.Request.Header.Get("Content-Type"))
	//if err != nil {
	//	c.Error(err)
	//	c.AbortWithStatus(http.StatusInternalServerError)
	//
	//	return
	//}
	body, err := io.ReadAll(c.Request.Body)

	if err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	log.Println("получен URL", string(body))
	shortURL, err := h.logic.URLShorter(string(body))

	if err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.Header("Content-Type", "text/plain")

	c.Status(http.StatusCreated)

	c.Writer.WriteString(shortURL)

}

func (h *Handler) URLGetID(c *gin.Context) {

	url, err := h.logic.URLGetID(c.Param("id"))
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	log.Println("отправлен url:", url)

	c.Header("Location", url)
	c.Status(http.StatusTemporaryRedirect)
}

//
//func (h *Handler) JSONURLGetID(w http.ResponseWriter, r *http.Request) {
//
//	var url constjson.URLLong
//	var buf bytes.Buffer
//
//	_, err := buf.ReadFrom(r.Body)
//	if err != nil {
//		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
//		return
//	}
//
//	if err = json.Unmarshal(buf.Bytes(), &url); err != nil {
//		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
//		return
//	}
//
//	longURL := url.URL
//	shortURL, err := h.logic.URLGetID(longURL[len(longURL)-8:])
//	if err != nil {
//		http.Error(w, "Error shorting URL", http.StatusBadRequest)
//		return
//	}
//	w.Header().Set("Location", shortURL)
//	w.WriteHeader(http.StatusTemporaryRedirect)
//}
