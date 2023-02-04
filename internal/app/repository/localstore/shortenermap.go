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

func (d *LocalStorage) SetURL(fullURL, shortURL string, id int) (string, error) {
	d.mx.Lock()
	defer d.mx.Unlock()
	if _, ok := d.users.Store[id]; ok {
		if _, ok := d.users.Exist[id][fullURL]; ok {
			for short, full := range d.users.Store[id] {
				if full == fullURL {
					return short, errors.New("url already exist")
				}
			}

		}
		d.users.Exist[id][fullURL] = 0
		d.users.Store[id][shortURL] = fullURL
	} else {
		d.users.Store[id] = map[string]string{shortURL: fullURL}
		d.users.Exist[id] = map[string]int{fullURL: 0}
	}

	if d.fileStore != nil {
		if err := d.fileStore.encoder.Encode(&d.users); err != nil {
			return "", err
		}
	}

	return "", nil
}

func (d *LocalStorage) GetURL(shortURL string) (string, error) {
	var url string
	d.mx.Lock()
	defer d.mx.Unlock()
	for id, s := range d.users.Store {
		for k, v := range s {
			if k == shortURL {
				url = v
				d.users.Exist[id][url]++
				if d.fileStore != nil {
					if err := d.fileStore.encoder.Encode(&d.users); err != nil {
						return url, err
					}
				}
				return url, nil
			}
		}
	}

	return url, errors.New("url is not exist")
}

func (d *LocalStorage) GetRequestCount(shortURL string) (int, error) {
	d.mx.Lock()
	defer d.mx.Unlock()

	for id, s := range d.users.Store {
		for k, v := range s {
			if k == shortURL {
				return d.users.Exist[id][v], nil
			}
		}
	}
	return 0, errors.New("url is not exist")
}

func (d *LocalStorage) GetURLCount() (int64, error) {
	var out int

	d.mx.Lock()
	defer d.mx.Unlock()
	for _, v := range d.users.Store {
		out += len(v)
		if out > 0 {
			out++
		}
	}

	return int64(out), nil
}

func (d *LocalStorage) GetAllUserURL(id int) ([]models.URLs, error) {
	var out = []models.URLs{}

	for k, v := range d.users.Store[id] {
		out = append(out, models.URLs{
			ShortURL:    k,
			OriginalURL: v,
		})
	}

	return out, nil
}

func (d *LocalStorage) CheckDBConn() error {
	return nil
}

func (d *LocalStorage) BatchSetURL(data []models.BatchInputJSON, shortURL []string, id int) ([]models.BatchInputJSON, error) {
	d.mx.Lock()
	defer d.mx.Unlock()
	for i := 0; i < len(data); i++ {
		if _, ok := d.users.Store[id]; ok {
			if _, ok := d.users.Exist[id][data[i].OriginalURL]; ok {
				return nil, errors.New("url already exist")
			}
			d.users.Exist[id][data[i].OriginalURL] = 0
			d.users.Store[id][shortURL[i]] = data[i].OriginalURL
		} else {
			d.users.Store[id] = map[string]string{shortURL[i]: data[i].OriginalURL}
			d.users.Exist[id] = map[string]int{data[i].OriginalURL: 0}
		}

		data[i].OriginalURL = shortURL[i]

		if d.fileStore != nil {
			if err := d.fileStore.encoder.Encode(&d.users); err != nil {
				return nil, err
			}
		}
	}

	return data, nil
}
