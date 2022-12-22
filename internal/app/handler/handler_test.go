package handler

import (
	"bytes"
	"github.com/cucumberjaye/url-shortener/internal/app/service"
	"github.com/cucumberjaye/url-shortener/internal/app/service/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_Shortener(t *testing.T) {
	type want struct {
		code     int
		response string
	}
	tests := []struct {
		name    string
		handler *Handler
		method  string
		way     string
		body    io.Reader
		want    want
	}{
		{
			name:   "ok_get",
			method: http.MethodGet,
			body:   nil,
			way:    "/test",
			want: want{
				code: 307,
			},
		},
		{
			name:   "ok_post",
			method: http.MethodPost,
			body:   bytes.NewBufferString("test.com"),
			way:    "/",
			want: want{
				code:     201,
				response: "test",
			},
		},
		{
			name:   "fail_get_400",
			method: http.MethodGet,
			body:   nil,
			way:    "/",
			want: want{
				code: 400,
			},
		},
		{
			name:   "fail_get_500",
			method: http.MethodGet,
			body:   nil,
			way:    "/none",
			want: want{
				code: 500,
			},
		},
		{
			name:   "ok_post",
			method: http.MethodPost,
			body:   bytes.NewBufferString("test.com"),
			way:    "/",
			want: want{
				code:     201,
				response: "test",
			},
		},
		{
			name:   "fail_post_400",
			method: http.MethodPost,
			body:   bytes.NewBufferString(""),
			way:    "/",
			want: want{
				code: 400,
			},
		},
		{
			name:   "fail_post_500",
			method: http.MethodPost,
			body:   bytes.NewBufferString("none"),
			way:    "/",
			want: want{
				code: 500,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			services := &service.Service{Shortener: &mocks.ServiceMock{}}
			tt.handler = NewHandler(services)
			request := httptest.NewRequest(tt.method, tt.way, tt.body)
			w := httptest.NewRecorder()
			h := http.HandlerFunc(tt.handler.Shortener)
			h.ServeHTTP(w, request)
			res := w.Result()
			assert.Equal(t, res.StatusCode, tt.want.code)
			defer res.Body.Close()
			resBody, err := io.ReadAll(res.Body)
			require.NoError(t, err)
			if tt.method == http.MethodPost && tt.want.code == 201 {
				assert.Equal(t, string(resBody), tt.want.response)
			}
		})
	}
}
