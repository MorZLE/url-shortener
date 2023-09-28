package handler

import (
	"github.com/MorZLE/url-shortener/internal/app/service"
	"github.com/MorZLE/url-shortener/internal/app/storage"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAppHandler_URLGetID(t *testing.T) {
	//type mockFn func(r *mocks.InterfaceAppService)

	type field struct {
		logic service.InterfaceAppService
	}

	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request

		//	m mockFn
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
			}
			r.HandleFunc(`/{id}`, h.URLGetID).Methods(http.MethodGet)
			r.ServeHTTP(tt.args.w, tt.args.r)
			assert.Equal(t, tt.args.w.Code, tt.wantStatus)
		})
	}
}

func TestAppHandler_URLShortener(t *testing.T) {
	type fields struct {
		InterfaceAppHandler InterfaceAppHandler
		logic               *service.AppService
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &AppHandler{
				InterfaceAppHandler: tt.fields.InterfaceAppHandler,
				logic:               tt.fields.logic,
			}
			h.URLShortener(tt.args.w, tt.args.r)
		})
	}
}
