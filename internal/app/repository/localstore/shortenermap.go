package localstore

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sync"
)

type db struct {
	Store map[string]string `json:"store"`
	Exist map[string]int    `json:"exist"`
}

type fileStore struct {
	fileStore *os.File
	encoder   *json.Encoder
}

type LocalStorage struct {
	db        db
	fileStore *fileStore
	mx        sync.Mutex
}

func NewShortenerDB(filename string) (*LocalStorage, error) {
	var fs *fileStore

	maps := db{
		Store: make(map[string]string),
		Exist: make(map[string]int),
	}

	if filename != "" {
		file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
		if err != nil {
			return nil, err
		}

		dec := json.NewDecoder(file)

		for dec.More() {
			err = dec.Decode(&maps)
			if err != nil {
				return nil, err
			}
		}

		delete(maps.Store, "")

		fs = &fileStore{
			fileStore: file,
			encoder:   json.NewEncoder(file),
		}
	}

	return &LocalStorage{
		db:        maps,
		fileStore: fs,
		mx:        sync.Mutex{},
	}, nil
}

func (d *LocalStorage) SetURL(fullURL, shortURL string) error {
	d.mx.Lock()
	defer d.mx.Unlock()
	fmt.Println(d.db)
	if _, ok := d.db.Exist[fullURL]; ok {
		return errors.New("url already exist")
	}
	d.db.Exist[fullURL] = 0
	d.db.Store[shortURL] = fullURL

	if d.fileStore != nil {
		maps := db{
			Store: map[string]string{shortURL: fullURL},
			Exist: map[string]int{fullURL: 0},
		}

		if err := d.fileStore.encoder.Encode(&maps); err != nil {
			return err
		}
	}

	return nil
}

func (d *LocalStorage) GetURL(shortURL string) (string, error) {
	d.mx.Lock()
	defer d.mx.Unlock()
	if url, ok := d.db.Store[shortURL]; !ok {
		return "", errors.New("url is not exist")
	} else {
		d.db.Exist[url]++

		if d.fileStore != nil {
			if err := d.fileStore.encoder.Encode(db{
				Store: map[string]string{"": ""},
				Exist: map[string]int{url: d.db.Exist[url]},
			}); err != nil {
				return url, err
			}
		}
		return url, nil
	}
}

func (d *LocalStorage) GetRequestCount(shortURL string) (int, error) {
	d.mx.Lock()
	defer d.mx.Unlock()
	if url, ok := d.db.Store[shortURL]; !ok {
		return 0, errors.New("url is not exist")
	} else {
		return d.db.Exist[url], nil
	}
}

func (d *LocalStorage) GetURLCount() int64 {
	tmp := len(d.db.Store)
	if tmp > 0 {
		return int64(tmp + 1)
	}
	return int64(tmp)
}
