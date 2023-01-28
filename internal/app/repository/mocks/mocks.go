package mocks

import (
	"errors"
	"github.com/cucumberjaye/url-shortener/models"
)

type RepositoryMock struct {
}

func (r *RepositoryMock) SetURL(fullURL, shortURL string, id int) error {
	if fullURL == "test.com" {
		return nil
	}
	return errors.New("test")
}

func (r *RepositoryMock) GetURL(shortURL string, id int) (string, error) {
	if shortURL == "0" {
		return "test.com", nil
	}
	return "", errors.New("test")
}

func (r *RepositoryMock) GetURLCount() map[int]int {
	return map[int]int{0: 0}
}

func (r *RepositoryMock) GetAllUserURL(id int) []models.URLs {
	return []models.URLs{}
}
