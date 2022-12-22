package mocks

import (
	"errors"
	"github.com/cucumberjaye/url-shortener/internal/app/service"
)

type ServiceMock struct {
	service.Shortener
}

func (m *ServiceMock) ShortingURL(fullURL string) (string, error) {
	if fullURL == "test.com" {
		return "test", nil
	}
	return "", errors.New("test")
}

func (m *ServiceMock) GetFullURL(shortURL string) (string, error) {
	if shortURL == "test" {
		return "test.com", nil
	}
	return "", errors.New("test")
}
