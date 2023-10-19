package logger

import (
	"go.uber.org/zap"
	"log"
	"net/http"
)

// Log будет доступен всему коду как синглтон.

var Log *zap.Logger = zap.NewNop()

// Initialize инициализирует синглтон логера с необходимым уровнем логирования.
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
	Log = zl
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

// RequestLogger — middleware-логер для входящих HTTP-запросов.
func RequestLogger(h http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		loggingResponseWriter := &LoggingResponseWriter{
			ResponseWriter: w,
			Status:         http.StatusOK,
		}

		//start := time.Now()
		h.ServeHTTP(loggingResponseWriter, r)

		//duration := time.Since(start)

		//Log.Info("Request",
		//	zap.String("Request", r.RequestURI),
		//	zap.String("method", r.Method),
		//	zap.String("uri", r.RequestURI),
		//	zap.String("duration", strconv.FormatInt(int64(duration), 10)),
		//)
		//Log.Info("Response",
		//	zap.String("Response", "Response"),
		//	zap.Int("status", loggingResponseWriter.Status),
		//	zap.String("method", r.Method),
		//	zap.Int("content_size", loggingResponseWriter.Size),
		//)
	})
}
func Error(info string, err error) {
	Log.Info(info, zap.Error(err))
}

func ShortURL(url string) {
	Log.Info("CreateShortURL", zap.String("url", url))
}

func Info(info string) {
	Log.Info("INFO", zap.String("info", info))
}
