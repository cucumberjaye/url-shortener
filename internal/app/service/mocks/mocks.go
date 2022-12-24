package mocks

import (
	"errors"
)

type ServiceMock struct {
}

func (m *ServiceMock) ShortingURL(fullURL string) (string, error) {
	if fullURL == "test.com" {
		return "0", nil
	}
	return "", errors.New("test")
}

func (m *ServiceMock) GetFullURL(shortURL string) (string, error) {
	if shortURL == "0" {
		return "test.com", nil
	}
	return "", errors.New("test")
}
