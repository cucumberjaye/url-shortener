package localstore

import (
	"github.com/cucumberjaye/url-shortener/internal/app/repository"
	"github.com/stretchr/testify/require"
	"os"
	"sync"
	"testing"
)

func TestDatabase_GetURL(t *testing.T) {
	type fields struct {
		Store map[string]map[string]string
	}
	type args struct {
		shortURL string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "ok",
			fields:  fields{Store: map[string]map[string]string{"1": {"0": "test.com"}}},
			args:    args{shortURL: "0"},
			want:    "test.com",
			wantErr: false,
		},
		{
			name:    "error",
			fields:  fields{Store: map[string]map[string]string{}},
			args:    args{shortURL: "0"},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &LocalStorage{
				users: repository.DB{
					Store: tt.fields.Store,
					Exist: map[string]map[string]int{"1": {}},
				},
				keeper: nil,
				mx:     sync.Mutex{},
			}
			got, err := d.GetURL(tt.args.shortURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetURL() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabase_SetURL(t *testing.T) {
	type fields struct {
		Store map[string]map[string]string
		Exist map[string]map[string]int
	}
	type args struct {
		fullURL  string
		shortURL string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				Store: map[string]map[string]string{},
				Exist: map[string]map[string]int{},
			},
			args:    args{shortURL: "0", fullURL: "test.com"},
			wantErr: false,
		},
		{
			name: "error",
			fields: fields{
				Store: map[string]map[string]string{"1": {"0": "test.com"}},
				Exist: map[string]map[string]int{"1": {"test.com": 0}},
			},
			args:    args{shortURL: "0", fullURL: "test.com"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, err := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
			require.NoError(t, err)
			defer os.Remove("test.txt")
			defer file.Close()
			d := &LocalStorage{
				users: repository.DB{
					Store: tt.fields.Store,
					Exist: tt.fields.Exist,
				},
				keeper: nil,
			}
			if _, err := d.SetURL(tt.args.fullURL, tt.args.shortURL, "1"); (err != nil) != tt.wantErr {
				t.Errorf("SetURL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
