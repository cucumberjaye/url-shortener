package app

import "math/rand"

type Service struct {
	db *Database
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func NewService(db *Database) *Service {
	return &Service{db: db}
}

func (s *Service) ShortingUrl(fullUrl string) (string, error) {
	shortUrl := shorting()
	if err := s.db.SetUrl(fullUrl, shortUrl); err != nil {
		return "", err
	}

	return shortUrl, nil
}

func (s *Service) GetFullUrl(shortUrl string) (string, error) {
	fullUrl, err := s.db.GetUrl(shortUrl)
	if err != nil {
		return "", err
	}

	return fullUrl, err
}

func shorting() string {
	b := make([]byte, 5)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
