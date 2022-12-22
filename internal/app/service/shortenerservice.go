package service

import (
	"github.com/cucumberjaye/url-shortener/internal/app/repository"
	"math/rand"
)

type ShortenerService struct {
	repos *repository.Repository
}

func NewShortenerService(repos *repository.Repository) *ShortenerService {
	return &ShortenerService{repos: repos}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func (s *ShortenerService) ShortingURL(fullURL string) (string, error) {
	shortURL := shorting()
	if err := s.repos.Shortener.SetURL(fullURL, shortURL); err != nil {
		return "", err
	}

	return "http://localhost:8080/" + shortURL, nil
}

func (s *ShortenerService) GetFullURL(shortURL string) (string, error) {
	fullURL, err := s.repos.Shortener.GetURL(shortURL)
	if err != nil {
		return "", err
	}

	return fullURL, err
}

func shorting() string {
	b := make([]byte, 5)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
