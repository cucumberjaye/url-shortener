package handler

import (
	"bytes"
	"github.com/cucumberjaye/url-shortener/configs"
	mocks2 "github.com/cucumberjaye/url-shortener/internal/app/service/mocks"
	"github.com/cucumberjaye/url-shortener/models"
	"github.com/cucumberjaye/url-shortener/pkg/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	configs.LoadConfig()
	code := m.Run()
	os.Exit(code)
}

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
			way:    "/0",
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
				response: "0",
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
			way:    "/error",
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
			body:   bytes.NewBufferString("error"),
			way:    "/",
			want: want{
				code: 500,
			},
		},
	}

	logger.New()
	logger.Discard()
	URLServices := &mocks2.ServiceMock{}
	logsServices := &mocks2.LogsMock{}
	ch := make(chan models.DeleteData)
	handlers := NewHandler(URLServices, logsServices, ch)

	r := handlers.InitRoutes()
	ts := httptest.NewServer(r)
	defer ts.Close()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(tt.method, ts.URL+tt.way, tt.body)
			request.RequestURI = ""

			http.DefaultClient.CheckRedirect = func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			}

			resp, err := http.DefaultClient.Do(request)
			require.NoError(t, err)

			assert.Equal(t, tt.want.code, resp.StatusCode, ts.URL+tt.way)

			defer resp.Body.Close()
			if tt.method == http.MethodPost && tt.want.code == 201 {
				resBody, err := io.ReadAll(resp.Body)
				require.NoError(t, err)
				assert.Equal(t, tt.want.response, string(resBody))
			}
		})
	}
}

func TestHandler_JSONShortener(t *testing.T) {
	type want struct {
		code     int
		response string
	}
	tests := []struct {
		name string
		body io.Reader
		want want
	}{
		{
			name: "ok",
			body: bytes.NewBufferString("{\"url\":\"test.com\"}"),
			want: want{
				code:     201,
				response: "{\"result\":\"0\"}\n",
			},
		},
		{
			name: "fail_post_500_empty",
			body: bytes.NewBufferString(""),
			want: want{
				code: 500,
			},
		},
		{
			name: "fail_post_500",
			body: bytes.NewBufferString("{\"url\":\"error\"}"),
			want: want{
				code: 500,
			},
		},
	}

	logger.New()
	logger.Discard()
	URLServices := &mocks2.ServiceMock{}
	logsServices := &mocks2.LogsMock{}
	ch := make(chan models.DeleteData)
	handlers := NewHandler(URLServices, logsServices, ch)

	r := handlers.InitRoutes()
	ts := httptest.NewServer(r)
	defer ts.Close()

	way := "/api/shorten"
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, ts.URL+way, tt.body)
			request.RequestURI = ""

			resp, err := http.DefaultClient.Do(request)
			require.NoError(t, err)

			assert.Equal(t, tt.want.code, resp.StatusCode, ts.URL+way)

			defer resp.Body.Close()
			if tt.want.code == 201 {
				resBody, err := io.ReadAll(resp.Body)
				require.NoError(t, err)
				assert.Equal(t, tt.want.response, string(resBody))
			}
		})
	}
}
