package hexshortener

import (
	"fmt"
	"github.com/cucumberjaye/url-shortener/internal/app/repository"
)

type ShortenerService struct {
	repos   repository.URLRepository
	counter int
}

func NewShortenerService(repos repository.URLRepository) *ShortenerService {
	return &ShortenerService{repos: repos}
}

func (s *ShortenerService) ShortingURL(fullURL string) (string, error) {
	shortURL := fmt.Sprintf("%x", s.counter)
	if err := s.repos.SetURL(fullURL, shortURL); err != nil {
		return "", err
	}
	s.counter++

	return shortURL, nil
}

func (s *ShortenerService) GetFullURL(shortURL string) (string, error) {
	fullURL, err := s.repos.GetURL(shortURL)
	if err != nil {
		return "", err
	}

	return fullURL, err
}
