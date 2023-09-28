package handler

import (
	"github.com/MorZLE/url-shortener/internal/app/service"
	"github.com/MorZLE/url-shortener/internal/app/storage"
	"github.com/MorZLE/url-shortener/internal/config"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAppHandler_URLGetID(t *testing.T) {

	type field struct {
		logic service.InterfaceAppService
	}

	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request

		//	m mockFn
	}
	cnf := config.Config{
		FlagAddrShortener: "localhost:8080",
		FlagAddrReq:       "localhost:8080",
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
		field      field
	}{
		{
			name: "Test case 1",

			args: args{
				w: &httptest.ResponseRecorder{},
				r: httptest.NewRequest(http.MethodGet, "http://localhost:8080/AWcwasd", nil),
			},
			field: field{
				logic: &service.AppService{
					Storage: &storage.AppStorage{
						M: map[string]string{
							"http://localhost:8080/AWcwasd": "http://localhost:8080/site.com",
						},
					},
					Cnf: cnf,
				},
			},
			wantStatus: http.StatusTemporaryRedirect,
		},
		{
			name: "Test case 2",

			args: args{
				w: &httptest.ResponseRecorder{},
				r: httptest.NewRequest(http.MethodGet, "http://localhost:8080/wadaw", nil),
			},
			field: field{
				logic: &service.AppService{
					Storage: &storage.AppStorage{
						M: map[string]string{
							"http://localhost:8080/sefsfvce": "http://localhost:8080/site.com",
						},
					},
					Cnf: config.Config{
						FlagAddrShortener: "localhost:8080",
						FlagAddrReq:       "localhost:8080",
					},
				},
			},
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "Test case 3",

			args: args{
				w: &httptest.ResponseRecorder{},
				r: httptest.NewRequest(http.MethodGet, "http://localhost:8080/gr43ge34g34t3g345g34g", nil),
			},
			field: field{
				logic: &service.AppService{
					Storage: &storage.AppStorage{
						M: map[string]string{
							"http://localhost:8080/sefsfvce": "http://localhost:8080/site.com",
						},
					},
					Cnf: config.Config{
						FlagAddrShortener: "localhost:8080",
						FlagAddrReq:       "localhost:8080",
					},
				},
			},
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := mux.NewRouter()

			h := &AppHandler{
				logic: tt.field.logic,
				cnf:   cnf,
			}
			r.HandleFunc(`/{id}`, h.URLGetID).Methods(http.MethodGet)
			r.ServeHTTP(tt.args.w, tt.args.r)
			assert.Equal(t, tt.args.w.Code, tt.wantStatus)
		})
	}
}

func TestAppHandler_URLShortener(t *testing.T) {

	type field struct {
		logic service.InterfaceAppService
	}

	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
		field      field
	}{
		{
			name: "Test case 1",

			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPost, "http://localhost:8080/", strings.NewReader("https://practicum.yandex.ru/")),
			},
			field: field{
				logic: &service.AppService{
					Storage: &storage.AppStorage{
						M: map[string]string{},
					},
					Cnf: config.Config{
						FlagAddrShortener: "localhost:8080",
						FlagAddrReq:       "localhost:8080",
					},
				},
			},
			wantStatus: http.StatusCreated,
		},
		{
			name: "Test case 2",

			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPost, "http://localhost:8080/", strings.NewReader(" -3040svgbfb-0o-ow4gm'xmfbzdbdbzdb")),
			},
			field: field{
				logic: &service.AppService{
					Storage: &storage.AppStorage{
						M: map[string]string{},
					},
					Cnf: config.Config{
						FlagAddrShortener: "localhost:8080",
						FlagAddrReq:       "localhost:8080",
					},
				},
			},
			wantStatus: http.StatusCreated,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := mux.NewRouter()

			h := &AppHandler{
				logic: tt.field.logic,
			}
			r.HandleFunc(`/`, h.URLShortener).Methods(http.MethodPost)
			r.ServeHTTP(tt.args.w, tt.args.r)
			assert.Equal(t, tt.args.w.Code, tt.wantStatus)
		})
	}
}
