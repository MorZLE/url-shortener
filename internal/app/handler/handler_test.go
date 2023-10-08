package handler

import (
	"bytes"
	"encoding/json"
	"github.com/MorZLE/url-shortener/internal/app/service"
	"github.com/MorZLE/url-shortener/internal/app/storage"
	"github.com/MorZLE/url-shortener/internal/config"
	"github.com/MorZLE/url-shortener/internal/constjson"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAppHandler_URLGetID(t *testing.T) {

	type field struct {
		logic service.Service
	}

	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request

		//	m mockFn
	}
	cnf := config.Config{
		BaseURL:    "http://127.0.0.1:8080",
		ServerAddr: "http://127.0.0.1:8080",
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
				r: httptest.NewRequest(http.MethodGet, "http://127.0.0.1:8080/AWcwasd", nil),
			},
			field: field{
				logic: service.Service{
					Storage: &storage.Storage{
						M: map[string]string{
							"AWcwasd": "http://127.0.0.1:8080/site.com",
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
				r: httptest.NewRequest(http.MethodGet, "http://127.0.0.1:8080/wadaw", nil),
			},
			field: field{
				logic: service.Service{
					Storage: &storage.Storage{
						M: map[string]string{
							"http://127.0.0.1/sefsfvce": "http://127.0.0.1/site.com",
						},
					},
					Cnf: cnf,
				},
			},
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "Test case 3",

			args: args{
				w: &httptest.ResponseRecorder{},
				r: httptest.NewRequest(http.MethodGet, "http://127.0.0.1:8080/gr43ge34g34t3g345g34g", nil),
			},
			field: field{
				logic: service.Service{
					Storage: &storage.Storage{
						M: map[string]string{
							"http://127.0.0.1:8080/sefsfvce": "http://127.0.0.1:8080/site.com",
						},
					},
					Cnf: cnf,
				},
			},
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := mux.NewRouter()

			h := &Handler{
				logic: &tt.field.logic,
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
		logic service.Service
	}

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
		field      field
	}{
		{
			name: "Test case 1",

			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPost, "http://localhost:8080/", strings.NewReader("https://practicum.yandex.ru/")),
			},
			field: field{
				logic: service.Service{
					Storage: &storage.Storage{
						M: map[string]string{},
					},
					Cnf: cnf,
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
				logic: service.Service{
					Storage: &storage.Storage{
						M: map[string]string{},
					},
					Cnf: cnf,
				},
			},
			wantStatus: http.StatusCreated,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := mux.NewRouter()

			h := &Handler{
				logic: &tt.field.logic,
				cnf:   cnf,
			}
			r.HandleFunc(`/`, h.URLShortener).Methods(http.MethodPost)
			r.ServeHTTP(tt.args.w, tt.args.r)
			assert.Equal(t, tt.args.w.Code, tt.wantStatus)
		})
	}
}

func TestHandler_JSONURLShort(t *testing.T) {
	type fields struct {
		logic service.Service
	}

	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	cnf := config.Config{
		BaseURL:    "http://127.0.0.1:8080",
		ServerAddr: "http://127.0.0.1:8080",
	}
	t1, _ := json.Marshal(&constjson.URLLong{URL: "https://practicum.yandex.ru/"})
	t2, _ := json.Marshal(&constjson.URLLong{URL: "https://vk.com/"})
	t3, _ := json.Marshal(&constjson.URLShort{Result: "https://vk.com/"})

	tests := []struct {
		name       string
		fields     fields
		args       args
		wantStatus int
	}{
		{
			name: "Test case 1",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/shorten", bytes.NewBuffer(t1)),
			},
			fields: fields{
				logic: service.Service{
					Storage: &storage.Storage{
						M: map[string]string{},
					},
					Cnf: cnf,
				},
			},
			wantStatus: http.StatusCreated,
		},
		{
			name: "Test case 2",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/shorten", bytes.NewBuffer(t2)),
			},
			fields: fields{
				logic: service.Service{
					Storage: &storage.Storage{
						M: map[string]string{},
					},
					Cnf: cnf,
				},
			},
			wantStatus: http.StatusCreated,
		},
		{
			name: "Test case 3",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/shorten", bytes.NewBuffer(t3)),
			},
			fields: fields{
				logic: service.Service{
					Storage: &storage.Storage{
						M: map[string]string{},
					},
					Cnf: cnf,
				},
			},
			wantStatus: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := mux.NewRouter()

			h := &Handler{
				logic: &tt.fields.logic,
				cnf:   cnf,
			}
			r.HandleFunc(`/api/shorten`, h.JSONURLShort).Methods(http.MethodPost)

			tt.args.w.Header().Set("Content-Type", "application/json")
			r.ServeHTTP(tt.args.w, tt.args.r)
			assert.Equal(t, tt.args.w.Code, tt.wantStatus)
		})
	}
}

//func TestHandler_JSONURLShortGzip(t *testing.T) {
//
//	type mckL func(r *mocks.ServiceInterface)
//	type mckS func(r *mocks.StorageInterface)
//
//	type fields struct {
//		mckL mckL
//		mckS mckS
//		Cnf  config.Config
//	}
//
//	type args struct {
//		w *httptest.ResponseRecorder
//		r *http.Request
//	}
//
//	cnf := config.Config{
//		BaseURL:    "http://127.0.0.1:8080",
//		ServerAddr: "http://127.0.0.1:8080",
//	}
//
//	var buf bytes.Buffer
//	t1, _ := json.Marshal(&constjson.URLLong{URL: "https://practicum.yandex.ru/"})
//	gz := gzip.NewWriter(&buf)
//	_, _ = gz.Write(t1)
//	_ = gz.Close()
//
//	w1 := cnf.BaseURL + "/qwd3212d"
//
//	tests := []struct {
//		name         string
//		fields       fields
//		args         args
//		wantStatus   int
//		wantShortURL string
//	}{
//		{
//			name: "Test case 1",
//			args: args{
//				w: httptest.NewRecorder(),
//				r: httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/shorten", bytes.NewBuffer(t1)),
//			},
//			fields: fields{
//				mckL: func(r *mocks.ServiceInterface) {
//					r.On("URLShorter", "https://practicum.yandex.ru/").Return(w1, nil)
//				},
//				mckS: func(r *mocks.StorageInterface) {
//					r.On("Set", "/qwd3212d").Return(nil)
//				},
//				Cnf: cnf,
//			},
//			wantStatus:   http.StatusCreated,
//			wantShortURL: w1,
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//
//			r := mux.NewRouter()
//			m := mocks.NewServiceInterface(t)
//			tt.fields.mckL(m)
//
//			h := &Handler{
//				logic: m,
//				cnf:   cnf,
//			}
//
//			tt.args.r.Header.Set("Content-Encoding", "gzip")
//			tt.args.r.Header.Set("Accept-Encoding", "application/json")
//
//			r.HandleFunc(`/api/shorten`, awd.GzipMiddleware(h.JSONURLShort)).Methods(http.MethodPost)
//			r.ServeHTTP(tt.args.w, tt.args.r)
//
//			body, err := io.ReadAll(tt.args.w.Body)
//			if err != nil {
//				log.Fatal(err)
//			}
//
//			reader, err := gzip.NewReader(bytes.NewReader(body))
//			if err != nil {
//				log.Fatal(err)
//			}
//
//			decompressedData, err := io.ReadAll(reader)
//			if err != nil {
//				log.Fatal(err)
//			}
//
//			reader.Close()
//			assert.Equal(t, decompressedData, tt.wantShortURL)
//			assert.Equal(t, tt.args.w.Code, tt.wantStatus)
//		})
//	}
//}
