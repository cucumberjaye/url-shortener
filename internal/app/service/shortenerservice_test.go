package service

import (
	"github.com/cucumberjaye/url-shortener/internal/app/repository"
	"github.com/cucumberjaye/url-shortener/internal/app/repository/mocks"
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
			args:    args{shortURL: "test"},
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
	repos := &repository.Repository{Shortener: &mocks.RepositoryMock{}}
	services := NewService(repos)
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
	repos := &repository.Repository{Shortener: &mocks.RepositoryMock{}}
	services := NewService(repos)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := services.ShortingURL(tt.args.fullURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("ShortingURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
