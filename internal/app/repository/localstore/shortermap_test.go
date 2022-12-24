package localstore

import (
	"testing"
)

func TestDatabase_GetURL(t *testing.T) {
	type fields struct {
		Store map[string]string
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
			fields:  fields{Store: map[string]string{"0": "test.com"}},
			args:    args{shortURL: "0"},
			want:    "test.com",
			wantErr: false,
		},
		{
			name:    "error",
			fields:  fields{Store: map[string]string{}},
			args:    args{shortURL: "0"},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &LocalStorage{
				Store: tt.fields.Store,
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
		Store map[string]string
		Exist map[string]struct{}
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
				Store: map[string]string{},
				Exist: map[string]struct{}{},
			},
			args:    args{shortURL: "0", fullURL: "test.com"},
			wantErr: false,
		},
		{
			name: "error",
			fields: fields{
				Store: map[string]string{"0": "test.com"},
				Exist: map[string]struct{}{"test.com": {}},
			},
			args:    args{shortURL: "0", fullURL: "test.com"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &LocalStorage{
				Store: tt.fields.Store,
				Exist: tt.fields.Exist,
			}
			if err := d.SetURL(tt.args.fullURL, tt.args.shortURL); (err != nil) != tt.wantErr {
				t.Errorf("SetURL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
