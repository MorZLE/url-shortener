package gzip

import (
	"compress/gzip"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strings"
)

// GzipMiddleware is a strongly typed function that adds gzip compression to the request and response bodies.
func GzipMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check the Content-Encoding header to determine the encoding
		contentEncoding := c.GetHeader("Content-Encoding")
		if strings.Contains(contentEncoding, "gzip") {
			// Создаем gzip.Reader для чтения сжатых данных
			reader, err := gzip.NewReader(c.Request.Body)
			if err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
			defer reader.Close()

			// Заменяем исходное тело запроса на распакованные данные
			c.Request.Body = http.MaxBytesReader(c.Writer, reader, c.Request.ContentLength)
			c.Request.Header.Del("Content-Encoding")
			c.Request.Header.Del("Content-Length")
			c.Request.Header.Del("Content-Type")
		}

		c.Next()
	}
}

// gzipWriter является оберткой над http.ResponseWriter для сжатия данных
type gzipWriter struct {
	gin.ResponseWriter
	writer io.Writer
}

func (g *gzipWriter) Write(data []byte) (int, error) {
	return g.writer.Write(data)
}
