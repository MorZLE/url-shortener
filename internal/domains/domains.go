package domains

import "net/http"

//go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name=HandlerInterface
type HandlerInterface interface {
	RunServer()
	URLShortener(w http.ResponseWriter, r *http.Request)
	URLGetID(w http.ResponseWriter, r *http.Request)
	JSONURLShort(w http.ResponseWriter, r *http.Request)
}

//go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name=StorageInterface
type StorageInterface interface {
	Set(key string, value string) error
	Get(key string) (string, error)
}

//go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name=ServiceInterface
type ServiceInterface interface {
	URLShorter(url string) (string, error)
	URLGetID(url string) (string, error)
}
