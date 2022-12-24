package localstore

import (
	"errors"
	"sync"
)

type LocalStorage struct {
	Store map[string]string
	Exist map[string]struct{}
	mx    sync.Mutex
}

func NewShortenerDB() *LocalStorage {
	return &LocalStorage{
		Store: make(map[string]string),
		Exist: make(map[string]struct{}),
	}
}

func (d *LocalStorage) SetURL(fullURL, shortURL string) error {
	d.mx.Lock()
	defer d.mx.Unlock()
	if _, ok := d.Exist[fullURL]; ok {
		return errors.New("url already exist")
	}
	d.Exist[fullURL] = struct{}{}
	d.Store[shortURL] = fullURL

	return nil
}

func (d *LocalStorage) GetURL(shortURL string) (string, error) {
	d.mx.Lock()
	defer d.mx.Unlock()
	if url, ok := d.Store[shortURL]; !ok {
		return "", errors.New("url is not exist")
	} else {
		return url, nil
	}
}
