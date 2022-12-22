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
		name   string
		method string
		way    string
		body   io.Reader
		want   want
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
			name:   "fail_get_405",
			method: http.MethodGet,
			body:   nil,
			way:    "/",
			want: want{
				code: 405,
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

	services := &service.Service{Shortener: &mocks.ServiceMock{}}
	handlers := NewHandler(services)

	r := handlers.InitRoutes()
	ts := httptest.NewServer(r)
	defer ts.Close()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(tt.method, ts.URL+tt.way, tt.body)
			request.RequestURI = ""

			resp, err := http.DefaultClient.Do(request)
			require.NoError(t, err)

			defer resp.Body.Close()
			if tt.method == http.MethodPost && tt.want.code == 201 {
				assert.Equal(t, tt.want.code, resp.StatusCode, ts.URL+tt.way)
				resBody, err := io.ReadAll(resp.Body)
				require.NoError(t, err)
				assert.Equal(t, tt.want.response, string(resBody))
			}
		})
	}
}
