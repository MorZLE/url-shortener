package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	gzipmilddle "github.com/MorZLE/url-shortener/internal/app/gzip"
	"github.com/MorZLE/url-shortener/internal/app/logger"
	"github.com/MorZLE/url-shortener/internal/config"
	"github.com/MorZLE/url-shortener/internal/consts"
	"github.com/MorZLE/url-shortener/internal/domains"
	"github.com/MorZLE/url-shortener/internal/models"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
)

func NewHandler(lg domains.Service, cnf *config.Config) Handler {
	return Handler{logic: lg, cnf: *cnf}
}

type Handler struct {
	logic domains.Service
	cnf   config.Config
}

func (h *Handler) RunServer() {

	router := gin.Default()
	router.Use(gzipmilddle.GzipMiddleware())
	router.Use(gzip.Gzip(gzip.BestSpeed))
	router.Use(gin.Logger())

	router.POST(`/`, h.URLShortener)
	router.POST(`/api/shorten`, h.JSONURLShort)
	router.POST(`/api/shorten/batch`, h.JSONURLShortBatch)
	router.GET(`/:id`, h.URLGetID)
	router.GET(`/ping`, h.CheckPing)

	log.Fatal(router.Run(h.cnf.ServerAddr))

}

func (h *Handler) JSONURLShortBatch(c *gin.Context) {
	var url []models.BatchSet

	if err := c.Bind(&url); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	res, err := h.logic.URLsShorter(url)
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusCreated, res)
	//resp, err := json.Marshal(&res)
	//if err != nil {
	//	c.Error(err)
	//	c.AbortWithStatus(http.StatusNotFound)
	//	return
	//}
	//
	//c.Header("Content-Type", "application/json")
	//c.Status(http.StatusCreated)
	//
	//c.Writer.Write(resp)
}

func (h *Handler) JSONURLShort(c *gin.Context) {
	var url models.URLLong

	if err := json.NewDecoder(c.Request.Body).Decode(&url); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	longURL := url.URL
	if longURL == "" {
		c.Error(consts.ErrGetURL)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	shortURL, err := h.logic.URLShorter(longURL)
	if err != nil {
		if errors.Is(err, consts.ErrDuplicateURL) {
			c.Status(http.StatusConflict)
		} else {
			logger.Error("Неожиданная ошибка", err)
			c.Error(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
	} else {
		c.Status(http.StatusCreated)
	}

	res := models.URLShort{
		Result: shortURL,
	}

	c.JSON(http.StatusCreated, res)
	//resp, err := json.Marshal(&res)
	//if err != nil {
	//	c.Error(err)
	//	c.AbortWithStatus(http.StatusNotFound)
	//	return
	//}
	//
	//c.Header("Content-Type", "application/json")
	//
	//c.Writer.Write(resp)

}

func (h *Handler) URLShortener(c *gin.Context) {

	body, err := io.ReadAll(c.Request.Body)

	if err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	logger.Info(fmt.Sprintf("получен URL %s", string(body)))
	shortURL, err := h.logic.URLShorter(string(body))

	if err != nil {
		if errors.Is(err, consts.ErrDuplicateURL) {
			c.Status(http.StatusConflict)
		} else {
			c.Error(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
	} else {
		c.Status(http.StatusCreated)
	}

	c.Header("Content-Type", "text/plain")

	c.Writer.WriteString(shortURL)
	logger.Info(fmt.Sprintf("отправлен URL %s", shortURL))
}

func (h *Handler) URLGetID(c *gin.Context) {

	url, err := h.logic.URLGetID(c.Param("id"))
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	logger.Info(fmt.Sprintf("отправлен url: %s", url))

	c.Header("Location", url)
	c.Status(http.StatusTemporaryRedirect)
}

func (h *Handler) CheckPing(c *gin.Context) {
	if err := h.logic.CheckPing(); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

//
//func (h *Handler) Cookie(c *gin.Context) string {
//	cookie, err := c.Cookie("auth")
//	if err != nil {
//		c.SetCookie("auth", h.logic.GenCookie(), 3600, "/", "", false, true)
//
//	}
//	return cookie
//}
