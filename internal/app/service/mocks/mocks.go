package mocks

import (
	"errors"
	"github.com/cucumberjaye/url-shortener/models"
)

type ServiceMock struct {
}

func (m *ServiceMock) ShortingURL(fullURL, baseURL string, id string) (string, error) {
	if fullURL == "test.com" {
		return "0", nil
	}
	return "", errors.New("test")
}

func (m *ServiceMock) GetFullURL(shortURL string) (string, error) {
	if shortURL[len(shortURL)-1] == '0' {
		return "test.com", nil
	}
	return "", errors.New("test")
}

func (m *ServiceMock) GetAllUserURL(id string) ([]models.URLs, error) {
	return []models.URLs{}, nil
}

func (m *ServiceMock) CheckDBConn() error {
	return nil
}

func (m *ServiceMock) BatchSetURL(data []models.BatchInputJSON, baseURL string, id string) ([]models.BatchInputJSON, error) {
	return []models.BatchInputJSON{}, nil
}
