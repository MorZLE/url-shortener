package handler

import (
	"bytes"
	gzclient "compress/gzip"
	"encoding/json"
	gzipmilddle "github.com/MorZLE/url-shortener/internal/app/gzip"
	"github.com/MorZLE/url-shortener/internal/app/service"
	"github.com/MorZLE/url-shortener/internal/app/storage"
	"github.com/MorZLE/url-shortener/internal/config"
	"github.com/MorZLE/url-shortener/internal/domains/mocks"
	"github.com/MorZLE/url-shortener/internal/models"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAppHandler_URLGetID(t *testing.T) {

	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	cnf := config.Config{
		BaseURL:    "http://127.0.0.1:8080",
		ServerAddr: "http://127.0.0.1:8080",
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
	}{
		{
			name: "Test case 1",
			args: args{
				w: &httptest.ResponseRecorder{},
				r: httptest.NewRequest(http.MethodGet, "http://127.0.0.1:8080/AWcwasd", nil),
			},

			wantStatus: http.StatusTemporaryRedirect,
		},
		{
			name: "Test case 2",

			args: args{
				w: &httptest.ResponseRecorder{},
				r: httptest.NewRequest(http.MethodGet, "http://127.0.0.1:8080/wadaw", nil),
			},
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gin.Default()

			st, err := storage.NewStorage(&cnf)
			st.Set("35fd02bf-0801-4041-b541-a4f2677e6df9", "AWcwasd", "http://127.0.0.1:8080/site.com")

			if err != nil {
				log.Fatal(err)
			}
			ser := service.NewService(st, &cnf)

			h := &Handler{
				logic: &ser,
				cnf:   cnf,
			}

			r.GET(`/:id`, h.URLGetID)
			tt.args.r.AddCookie(&http.Cookie{
				Name:  "auth",
				Value: "35fd02bf-0801-4041-b541-a4f2677e6df9",
				Path:  "/",
			})
			r.ServeHTTP(tt.args.w, tt.args.r)
			assert.Equal(t, tt.args.w.Code, tt.wantStatus)
		})
	}
}

func TestAppHandler_URLShortener(t *testing.T) {

	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	cnf := config.Config{
		BaseURL:    "http://127.0.0.1:8080",
		ServerAddr: "http://127.0.0.1:8080",
	}

	tests := []struct {
		name       string
		args       args
		wantStatus int
	}{
		{
			name: "Test case 1",

			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPost, "http://localhost:8080/", strings.NewReader("https://practicum.yandex.ru/")),
			},

			wantStatus: http.StatusCreated,
		},
		{
			name: "Test case 2",

			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPost, "http://localhost:8080/", strings.NewReader(" -3040svgbfb-0o-ow4gm'xmfbzdbdbzdb")),
			},

			wantStatus: http.StatusCreated,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gin.Default()

			st, err := storage.NewStorage(&cnf)

			if err != nil {
				log.Fatal(err)
			}
			ser := service.NewService(st, &cnf)

			h := &Handler{
				logic: &ser,
				cnf:   cnf,
			}

			r.POST(`/`, h.URLShortener)
			r.ServeHTTP(tt.args.w, tt.args.r)
			assert.Equal(t, tt.wantStatus, tt.args.w.Code)
		})
	}
}

func TestHandler_JSONURLShort(t *testing.T) {

	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	cnf := config.Config{
		BaseURL:    "http://127.0.0.1:8080",
		ServerAddr: "http://127.0.0.1:8080",
	}
	t1, _ := json.Marshal(&models.URLLong{URL: "https://practicum.yandex.ru/"})
	t2, _ := json.Marshal(&models.URLLong{URL: "https://vk.com/"})
	t3, _ := json.Marshal(&models.URLShort{Result: "https://vk.com/"})

	tests := []struct {
		name string

		args       args
		wantStatus int
	}{
		{
			name: "Test case 1",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/shorten", bytes.NewBuffer(t1)),
			},

			wantStatus: http.StatusCreated,
		},
		{
			name: "Test case 2",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/shorten", bytes.NewBuffer(t2)),
			},

			wantStatus: http.StatusCreated,
		},
		{
			name: "Test case 3",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/shorten", bytes.NewBuffer(t3)),
			},

			wantStatus: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gin.Default()

			st, err := storage.NewStorage(&cnf)

			if err != nil {
				log.Fatal(err)
			}
			ser := service.NewService(st, &cnf)

			h := &Handler{
				logic: &ser,
				cnf:   cnf,
			}

			r.POST(`/api/shorten`, h.JSONURLShort)
			tt.args.r.Header.Set("Content-Type", "application/json")
			tt.args.r.Header.Set("Accept", "application/json")

			tt.args.w.Header().Set("Content-Type", "application/json")
			r.ServeHTTP(tt.args.w, tt.args.r)
			assert.Equal(t, tt.args.w.Code, tt.wantStatus)
		})
	}
}

func TestHandler_JSONURLShortGzip(t *testing.T) {

	type mckL func(r *mocks.Service)
	type mckS func(r *mocks.Storage)

	type fields struct {
		mckL mckL
		mckS mckS
		cnf  config.Config
	}

	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}

	cnf := config.Config{
		BaseURL:    "http://127.0.0.1:8080",
		ServerAddr: "http://127.0.0.1:8080",
	}

	var buf1 bytes.Buffer
	t1 := `{"URL": "https://practicum.yandex.ru/"}`
	gz1 := gzclient.NewWriter(&buf1)
	_, _ = gz1.Write([]byte(t1))
	_ = gz1.Close()

	w1 := cnf.BaseURL + "/qwd3212d"

	tests := []struct {
		name         string
		fields       fields
		args         args
		wantStatus   int
		wantShortURL string
	}{
		{
			name: "Test case 1",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/shorten", bytes.NewBuffer(buf1.Bytes())),
			},
			fields: fields{
				mckL: func(r *mocks.Service) {
					r.On("URLShorter", "35fd02bf-0801-4041-b541-a4f2677e6df9", "https://practicum.yandex.ru/").Return(w1, nil)
					r.On("GenerateCookie").Return("35fd02bf-0801-4041-b541-a4f2677e6df9", nil)
				},
				mckS: func(r *mocks.Storage) {
					r.On("Set", "/qwd3212d").Return(nil)
				},
				cnf: cnf,
			},
			wantStatus:   http.StatusCreated,
			wantShortURL: w1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			r := gin.Default()
			m := mocks.NewService(t)

			tt.fields.mckL(m)

			h := &Handler{
				logic: m,
				cnf:   cnf,
			}
			tt.args.r.Header.Set("Accept-Encoding", "gzip")
			tt.args.r.Header.Set("Content-Encoding", "gzip")
			tt.args.r.Header.Set("Content-Type", "application/json")

			r.Use(gzipmilddle.GzipMiddleware())
			r.Use(gzip.Gzip(gzip.BestSpeed))

			r.POST(`/api/shorten`, h.JSONURLShort)

			r.ServeHTTP(tt.args.w, tt.args.r)

			body, err := io.ReadAll(tt.args.w.Body)
			if err != nil {
				log.Fatal(err)
			}
			reader, err := gzclient.NewReader(bytes.NewReader(body))
			if err != nil {
				log.Fatal(err)
			}

			reader.Close()

			var url models.URLShort

			if err := json.NewDecoder(reader).Decode(&url); err != nil {
				log.Fatal(err)
				return
			}

			assert.Equal(t, url.Result, tt.wantShortURL)
			assert.Equal(t, tt.args.w.Code, tt.wantStatus)
		})
	}
}
