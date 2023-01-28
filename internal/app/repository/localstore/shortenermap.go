package localstore

import (
	"encoding/json"
	"errors"
	"github.com/cucumberjaye/url-shortener/models"
	"os"
	"sync"
)

type db struct {
	Store map[int]map[string]string `json:"store"`
	Exist map[int]map[string]int    `json:"exist"`
}

type fileStore struct {
	fileStore *os.File
	encoder   *json.Encoder
}

type LocalStorage struct {
	users     db
	fileStore *fileStore
	mx        sync.Mutex
}

func NewShortenerDB(filename string) (*LocalStorage, error) {
	var fs *fileStore

	users := db{
		Store: map[int]map[string]string{},
		Exist: map[int]map[string]int{},
	}

	if filename != "" {
		file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
		if err != nil {
			return nil, err
		}

		dec := json.NewDecoder(file)

		for dec.More() {
			err = dec.Decode(&users)
			if err != nil {
				return nil, err
			}
		}

		fs = &fileStore{
			fileStore: file,
			encoder:   json.NewEncoder(file),
		}
	}

	return &LocalStorage{
		users:     users,
		fileStore: fs,
		mx:        sync.Mutex{},
	}, nil
}

func (d *LocalStorage) SetURL(fullURL, shortURL string, id int) error {
	d.mx.Lock()
	defer d.mx.Unlock()
	if _, ok := d.users.Store[id]; ok {
		if _, ok := d.users.Exist[id][fullURL]; ok {
			return errors.New("url already exist")
		}
		d.users.Exist[id][fullURL] = 0
		d.users.Store[id][shortURL] = fullURL
	} else {
		d.users.Store[id] = map[string]string{shortURL: fullURL}
		d.users.Exist[id] = map[string]int{fullURL: 0}
	}

	if d.fileStore != nil {
		if err := d.fileStore.encoder.Encode(&d.users); err != nil {
			return err
		}
	}

	return nil
}

func (d *LocalStorage) GetURL(shortURL string, id int) (string, error) {
	var url string
	d.mx.Lock()
	defer d.mx.Unlock()
	if u, ok := d.users.Store[id]; !ok {
		return "", errors.New("url is not exist")
	} else {
		if url, ok = u[shortURL]; !ok {
			return "", errors.New("url is not exist")
		}
	}
	d.users.Exist[id][url]++
	if d.fileStore != nil {
		if err := d.fileStore.encoder.Encode(&d.users); err != nil {
			return url, err
		}
	}

	return url, nil
}

func (d *LocalStorage) GetRequestCount(shortURL string, id int) (int, error) {
	d.mx.Lock()
	defer d.mx.Unlock()
	if url, ok := d.users.Store[id][shortURL]; !ok {
		return 0, errors.New("url is not exist")
	} else {
		return d.users.Exist[id][url], nil
	}
}

func (d *LocalStorage) GetURLCount() map[int]int {
	var out = map[int]int{}

	d.mx.Lock()
	defer d.mx.Unlock()
	for k, v := range d.users.Store {
		out[k] = len(v)
		if out[k] > 0 {
			out[k]++
		}
	}

	return out
}

func (d *LocalStorage) GetAllUserURL(id int) []models.URLs {
	var out = []models.URLs{}

	for k, v := range d.users.Store[id] {
		out = append(out, models.URLs{
			ShortURL:    k,
			OriginalURL: v,
		})
	}

	return out
}
