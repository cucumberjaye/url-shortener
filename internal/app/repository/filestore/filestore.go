package filestore

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/cucumberjaye/url-shortener/internal/app/repository"
)

// Структура для записи и чтения в файл в формате JSON
type FileStore struct {
	fileStore *os.File
	encoder   *json.Encoder
}

// Создаем FileStore
func New(filename string) (*FileStore, error) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		return nil, err
	}

	return &FileStore{
		fileStore: file,
		encoder:   json.NewEncoder(file),
	}, nil
}

// Проверяет, что возможно взаимодействовать с файлом
func (k *FileStore) CheckKeeper() error {
	if k.fileStore == nil {
		return errors.New("file does not exist")
	}

	return nil
}

// Получаем содержимое файла и возвращаем
func (k *FileStore) GetAllData() (repository.DB, error) {
	var users = repository.DB{
		Store: map[string]map[string]string{},
		Exist: map[string]map[string]int{},
	}
	dec := json.NewDecoder(k.fileStore)

	for dec.More() {
		var tmp repository.DB
		err := dec.Decode(&tmp)
		if err != nil {
			return users, err
		}
		for key, val := range tmp.Store {
			for short, full := range val {
				if _, ok := users.Store[key]; ok {
					users.Store[key][short] = full
				} else {
					users.Store[key] = map[string]string{short: full}
				}
			}
		}
		for key, val := range tmp.Exist {
			for full, count := range val {
				if _, ok := users.Exist[key]; ok {
					users.Exist[key][full] = count
				} else {
					users.Exist[key] = map[string]int{full: count}
				}
			}
		}
	}

	return users, nil
}

// Записываем в файл
func (k *FileStore) Set(users repository.DB) error {
	if err := k.encoder.Encode(&users); err != nil {
		return err
	}

	return nil
}
