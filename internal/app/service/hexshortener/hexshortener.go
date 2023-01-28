package hexshortener

import (
	"fmt"
	"github.com/cucumberjaye/url-shortener/internal/app/service"
	"github.com/cucumberjaye/url-shortener/models"
	"sync"
	"sync/atomic"
)

type ShortenerService struct {
	repos   service.URLRepository
	counter int64
	mx      sync.Mutex
}

func NewShortenerService(repos service.URLRepository) *ShortenerService {
	return &ShortenerService{
		repos:   repos,
		counter: repos.GetURLCount(),
	}
}

func (s *ShortenerService) ShortingURL(fullURL, baseURL string, id int) (string, error) {
	shortURL := baseURL + fmt.Sprintf("%x", s.counter)
	if err := s.repos.SetURL(fullURL, shortURL, id); err != nil {
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

func (s *ShortenerService) GetAllUserURL(id int) []models.URLs {
	return s.repos.GetAllUserURL(id)
}
