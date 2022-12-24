package service

type URLService interface {
	ShortingURL(fullURL string) (string, error)
	GetFullURL(shortURL string) (string, error)
}
