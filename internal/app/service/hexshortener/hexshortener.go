package hexshortener

import (
	"fmt"
	"github.com/cucumberjaye/url-shortener/internal/app/service"
	"sync/atomic"
)

type ShortenerService struct {
	repos   service.URLRepository
	counter int64
}

func NewShortenerService(repos service.URLRepository) *ShortenerService {
	return &ShortenerService{repos: repos}
}

func (s *ShortenerService) ShortingURL(fullURL string) (string, error) {
	shortURL := fmt.Sprintf("%x", s.counter)
	if err := s.repos.SetURL(fullURL, shortURL); err != nil {
		return "", err
	}
	atomic.AddInt64(&s.counter, 1)

	return shortURL, nil
}

func (s *ShortenerService) GetFullURL(shortURL string) (string, error) {
	fullURL, err := s.repos.GetURL(shortURL)
	if err != nil {
		return "", err
	}

	return fullURL, err
}
