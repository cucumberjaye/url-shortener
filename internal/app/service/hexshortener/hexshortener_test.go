package hexshortener

import (
	"errors"
	"fmt"
	"github.com/cucumberjaye/url-shortener/internal/app/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestShortenerService_GetFullURL(t *testing.T) {
	type args struct {
		shortURL string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		err     error
		wantErr bool
	}{
		{
			name:    "ok",
			args:    args{shortURL: "0"},
			want:    "test.com",
			err:     nil,
			wantErr: false,
		},
		{
			name:    "fail",
			args:    args{shortURL: "fail"},
			want:    "",
			err:     errors.New("test"),
			wantErr: true,
		},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ping := mocks.NewMockSQLRepository(ctrl)
	pgs := mocks.NewMockURLRepository(ctrl)
	pgs.EXPECT().GetURLCount().Return(int64(0), nil)
	services, err := NewShortenerService(pgs, ping)
	require.NoError(t, err)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pgs.EXPECT().GetURL(tt.args.shortURL).Return(tt.want, tt.err)
			got, err := services.GetFullURL(tt.args.shortURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFullURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetFullURL() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShortenerService_ShortingURL(t *testing.T) {
	type args struct {
		fullURL string
	}
	tests := []struct {
		name    string
		args    args
		want    error
		wantErr bool
	}{
		{
			name:    "ok",
			args:    args{fullURL: "test.com"},
			want:    nil,
			wantErr: false,
		},
		{
			name:    "fail",
			args:    args{fullURL: "pop.corn"},
			want:    errors.New("test"),
			wantErr: true,
		},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ping := mocks.NewMockSQLRepository(ctrl)
	pgs := mocks.NewMockURLRepository(ctrl)
	pgs.EXPECT().GetURLCount().Return(int64(0), nil)
	services, err := NewShortenerService(pgs, ping)
	require.NoError(t, err)
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			pgs.EXPECT().SetURL(tt.args.fullURL, fmt.Sprintf("%d", i), 0).Return(tt.want)
			_, err = services.ShortingURL(tt.args.fullURL, "", 0)
			if (err != nil) != tt.wantErr {
				t.Errorf("ShortingURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
