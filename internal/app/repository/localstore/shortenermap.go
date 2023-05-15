package localstore

import (
	"errors"
	"fmt"
	"sync"

	"github.com/cucumberjaye/url-shortener/internal/app/repository"
	"github.com/cucumberjaye/url-shortener/models"
)

// Структура для хранения в памяти программы
type LocalStorage struct {
	users  repository.DB
	keeper repository.Keeper
	mx     sync.Mutex
}

// Создаем локальное хранилище
func NewShortenerDB(keeper repository.Keeper) (*LocalStorage, error) {
	var users = repository.DB{
		Store: map[string]map[string]string{},
		Exist: map[string]map[string]int{},
	}
	var err error

	if keeper != nil {
		users, err = keeper.GetAllData()
		if err != nil {
			return nil, err
		}
	}

	return &LocalStorage{
		users:  users,
		keeper: keeper,
		mx:     sync.Mutex{},
	}, nil
}

// Записываем короткую ссылку в хранилище и в файл или базу данных
func (d *LocalStorage) SetURL(fullURL, shortURL string, id string) (string, error) {
	d.mx.Lock()
	defer d.mx.Unlock()
	if _, ok := d.users.Store[id]; ok {
		if _, ok := d.users.Exist[id][fullURL]; ok {
			for short, full := range d.users.Store[id] {
				if full == fullURL {
					return short, errors.New("url already exists")
				}
			}

		}
		d.users.Exist[id][fullURL] = 0
		d.users.Store[id][shortURL] = fullURL
	} else {
		d.users.Store[id] = map[string]string{shortURL: fullURL}
		d.users.Exist[id] = map[string]int{fullURL: 0}
	}

	if d.keeper != nil {
		user := repository.DB{
			Store: map[string]map[string]string{id: {shortURL: fullURL}},
			Exist: map[string]map[string]int{id: {fullURL: 0}},
		}
		if err := d.keeper.Set(user); err != nil {
			return "", err
		}
	}

	return "", nil
}

// Получаем полную ссылку по сокращенной из хранилища
func (d *LocalStorage) GetURL(shortURL string) (string, error) {
	var url string
	d.mx.Lock()
	defer d.mx.Unlock()
	for id, s := range d.users.Store {
		for k, v := range s {
			if k == shortURL {
				if d.users.Exist[id][v] == -1 {
					return "", errors.New("URL was deleted")
				}
				url = v
				d.users.Exist[id][url]++
				if d.keeper != nil {
					user := repository.DB{
						Store: nil,
						Exist: map[string]map[string]int{id: {url: d.users.Exist[id][url]}},
					}
					if err := d.keeper.Set(user); err != nil {
						return url, err
					}
				}
				return url, nil
			}
		}
	}

	return url, errors.New("url is not exists")
}

// Получаем количество получений по некоторой сокращенной ссылке
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
	return 0, errors.New("url is not exists")
}

// Получаем количество хранящихся ссылок
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

// Получаем все ссылки пользователя
func (d *LocalStorage) GetAllUserURL(id string) ([]models.URLs, error) {
	var out = []models.URLs{}

	for k, v := range d.users.Store[id] {
		out = append(out, models.URLs{
			ShortURL:    k,
			OriginalURL: v,
		})
	}

	return out, nil
}

// Сохраняем ссылки пачкой
func (d *LocalStorage) BatchSetURL(data []models.BatchInputJSON, shortURL []string, id string) ([]models.BatchInputJSON, error) {
	d.mx.Lock()
	defer d.mx.Unlock()
	for i := 0; i < len(data); i++ {
		if _, ok := d.users.Store[id]; ok {
			if _, ok := d.users.Exist[id][data[i].OriginalURL]; ok {
				return nil, errors.New("url already exists")
			}
			d.users.Exist[id][data[i].OriginalURL] = 0
			d.users.Store[id][shortURL[i]] = data[i].OriginalURL
		} else {
			d.users.Store[id] = map[string]string{shortURL[i]: data[i].OriginalURL}
			d.users.Exist[id] = map[string]int{data[i].OriginalURL: 0}
		}

		data[i].OriginalURL = shortURL[i]

		if d.keeper != nil {
			user := repository.DB{
				Store: map[string]map[string]string{id: {shortURL[i]: data[i].OriginalURL}},
				Exist: map[string]map[string]int{id: {data[i].OriginalURL: 0}},
			}
			if err := d.keeper.Set(user); err != nil {
				return nil, err
			}
		}
	}

	return data, nil
}

// Удаляем ссылки пачкой
func (d *LocalStorage) BatchDeleteURL(short, id string) error {
	d.mx.Lock()
	defer d.mx.Unlock()
	if _, ok := d.users.Store[id]; ok {
		if full, ok := d.users.Store[id][short]; ok {
			if d.users.Exist[id][full] == -1 {
				return nil
			}
			d.users.Exist[id][full] = -1
			if d.keeper != nil {
				user := repository.DB{
					Store: nil,
					Exist: map[string]map[string]int{id: {full: -1}},
				}
				if err := d.keeper.Set(user); err != nil {
					return err
				}
			}
		}
	} else {
		return fmt.Errorf("url %s does not exist", short)
	}
	return nil
}

// Проверяем постоянное хранилище на работоспособность
func (d *LocalStorage) CheckStorage() error {
	return d.keeper.CheckKeeper()
}

// получаем количество пользователей и ссылок
func (d *LocalStorage) GetStats() (models.Stats, error) {
	var urlsCount, usersCount int

	for _, user := range d.users.Exist {
		usersCount++
		for range user {
			urlsCount++
		}
	}

	return models.Stats{
		URLs:  urlsCount,
		Users: usersCount,
	}, nil
}
