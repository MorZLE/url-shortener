package domains

import (
	"github.com/MorZLE/url-shortener/internal/models"
	"github.com/gin-gonic/gin"
)

//go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name=HandlerInterface
type HandlerInterface interface {
	RunServer()
	URLShortener(c *gin.Context)
	URLGetID(c *gin.Context)
	CheckPing(c *gin.Context) error
	JSONURLShort(c *gin.Context, obj models.URLShort)
}

//go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name=StorageInterface
type StorageInterface interface {
	Set(key string, value string) error
	Get(key string) (string, error)
	Count() int
	Close() error
	Ping() error
	SetBatch(map[string]string) error
	GetDuplicate(longURL string) (string, error)
}

//go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name=ServiceInterface
type ServiceInterface interface {
	URLShorter(url string) (string, error)
	URLGetID(url string) (string, error)
	CheckPing() error
	URLsShorter(url []models.BatchSet) ([]models.BatchGet, error)
}
