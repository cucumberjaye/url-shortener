package hexshortener

import (
	"github.com/cucumberjaye/url-shortener/internal/app/repository/mocks"
	"github.com/golang/mock/gomock"
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
		wantErr bool
	}{
		{
			name:    "ok",
			args:    args{shortURL: "0"},
			want:    "test.com",
			wantErr: false,
		},
		{
			name:    "fail",
			args:    args{shortURL: "fail"},
			want:    "",
			wantErr: true,
		},
	}
	repos := &mocks.RepositoryMock{}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	rSQL := mocks.NewMockSQLRepository(ctrl)
	services := NewShortenerService(repos, rSQL)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
		wantErr bool
	}{
		{
			name:    "ok",
			args:    args{fullURL: "test.com"},
			wantErr: false,
		},
		{
			name:    "fail",
			args:    args{fullURL: "pop.corn"},
			wantErr: true,
		},
	}
	repos := &mocks.RepositoryMock{}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	rSQL := mocks.NewMockSQLRepository(ctrl)
	services := NewShortenerService(repos, rSQL)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := services.ShortingURL(tt.args.fullURL, "", 0)
			if (err != nil) != tt.wantErr {
				t.Errorf("ShortingURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
