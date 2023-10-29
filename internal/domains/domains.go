package domains

import (
	"github.com/MorZLE/url-shortener/internal/models"
	"github.com/gin-gonic/gin"
)

//go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name=Handler
type Handler interface {
	RunServer()
	URLShortener(c *gin.Context)
	URLGetID(c *gin.Context)
	CheckPing(c *gin.Context) error
	JSONURLShort(c *gin.Context, obj models.URLShort)
	JSONURLShortBatch(c *gin.Context)
	URLGetCookie(c *gin.Context)
	Cookie(c *gin.Context) string
	URLDelete(c *gin.Context)
}

//go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name=Storage
type Storage interface {
	Set(id, key, value string) error
	Get(key string) (string, error)
	Count() (int, error)
	Close() error
	Ping() error
	SetBatch(string, map[string]string) error
	GetDuplicate(longURL string) (string, error)
	GetAllURL(id string) (map[string]string, error)
	UpdateDelete(id string, key []string) error
}

//go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name=Service
type Service interface {
	URLShorter(id, url string) (string, error)
	URLGetID(url string) (string, error)
	CheckPing() error
	URLsShorter(id string, url []models.BatchSet) ([]models.BatchGet, error)
	Generate(num int) (string, error)
	GetAllURLUsers(id string) ([]models.AllURLs, error)
	GenerateCookie() string
	URLDelete(id string, urls []string)
}
