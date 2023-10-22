package logger

import (
	"go.uber.org/zap"
	"log"
	"net/http"
)

var mylog *zap.Logger = zap.NewNop()

func Initialize() {
	// преобразуем текстовый уровень логирования в zap.AtomicLevel
	lvl, err := zap.ParseAtomicLevel("info")
	if err != nil {
		log.Fatal(err)
	}
	// создаём новую конфигурацию логера
	cfg := zap.NewProductionConfig()
	// устанавливаем уровень
	cfg.Level = lvl
	// создаём логер на основе конфигурации
	zl, err := cfg.Build()
	if err != nil {
		log.Fatal(err)
	}
	// устанавливаем синглтон
	mylog = zl
}

type LoggingResponseWriter struct {
	http.ResponseWriter
	Status int
	Size   int
}

func (w *LoggingResponseWriter) WriteHeader(status int) {
	w.Status = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *LoggingResponseWriter) Write(b []byte) (int, error) {
	size, err := w.ResponseWriter.Write(b)
	w.Size += size
	return size, err
}

func Error(info string, err error) {
	mylog.Info(info, zap.Error(err))
}

func ShortURL(url string) {
	mylog.Info("CreateShortURL", zap.String("url", url))

}

func Info(info string) {
	mylog.Info("INFO", zap.String("info", info))
}
