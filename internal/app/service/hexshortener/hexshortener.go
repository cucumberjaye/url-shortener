package hexshortener

import (
	"fmt"
	"github.com/cucumberjaye/url-shortener/internal/app/service"
	"github.com/cucumberjaye/url-shortener/models"
	"sync/atomic"
)

const workers = 5

type ShortenerService struct {
	repos   service.URLRepository
	counter int64
}

func NewShortenerService(repos service.URLRepository) (*ShortenerService, error) {
	c, err := repos.GetURLCount()
	if err != nil {
		return nil, err
	}
	return &ShortenerService{
		repos:   repos,
		counter: c,
	}, nil
}

func (s *ShortenerService) ShortingURL(fullURL, baseURL string, id string) (string, error) {
	shortURL := baseURL + fmt.Sprintf("%x", s.counter)
	if short, err := s.repos.SetURL(fullURL, shortURL, id); err != nil {
		if err.Error() == "url already exists" {
			return short, err
		}
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

func (s *ShortenerService) GetAllUserURL(id string) ([]models.URLs, error) {
	return s.repos.GetAllUserURL(id)
}

func (s *ShortenerService) CheckDBConn() error {
	return s.repos.CheckStorage()
}

func (s *ShortenerService) BatchSetURL(data []models.BatchInputJSON, baseURL string, id string) ([]models.BatchInputJSON, error) {
	var shortURL = []string{}
	for i := 0; i < len(data); i++ {
		shortURL = append(shortURL, baseURL+fmt.Sprintf("%x", s.counter))
		atomic.AddInt64(&s.counter, 1)
	}
	return s.repos.BatchSetURL(data, shortURL, id)
}
