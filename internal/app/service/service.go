package service

type URLRepository interface {
	SetURL(fullURL, shortURL string) error
	GetURL(shortURL string) (string, error)
}
