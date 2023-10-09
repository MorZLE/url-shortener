package gzip

import (
	"compress/gzip"
	"net/http"
)

// compressWriter реализует интерфейс http.ResponseWriter и позволяет прозрачно для сервера
// сжимать передаваемые данные и выставлять правильные HTTP-заголовки
type compressWriter struct {
	w  http.ResponseWriter
	zw *gzip.Writer
}

func newCompressWriter(w http.ResponseWriter) *compressWriter {
	return &compressWriter{
		w:  w,
		zw: gzip.NewWriter(w),
	}
}

func (c *compressWriter) Header() http.Header {
	return c.w.Header()
}

func (c *compressWriter) Write(p []byte) (int, error) {
	return c.zw.Write(p)
}

func (c *compressWriter) WriteHeader(statusCode int) {
	if statusCode < 300 {
		c.w.Header().Set("Content-Encoding", "gzip")
	}
	c.w.WriteHeader(statusCode)
}

// Close закрывает gzip.Writer и досылает все данные из буфера.
func (c *compressWriter) Close() error {
	return c.zw.Close()
}

//func GzipMiddleware(h http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		originalWriter := w
//
//		if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
//			compressedWriter := newCompressWriter(w)
//			compressedWriter.Header().Set("Content-Encoding", "gzip")
//			originalWriter = compressedWriter
//			defer compressedWriter.Close()
//		}
//
//		if strings.Contains(r.Header.Get("Content-Encoding"), "gzip") {
//			gzipReader, err := gzip.NewReader(r.Body)
//			if err != nil {
//				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
//				return
//			}
//			defer gzipReader.Close()
//
//			r.Body = gzipReader
//		}
//
//		h.ServeHTTP(originalWriter, r)
//	})
//}
