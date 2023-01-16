package configs

import "os"

var (
	ServerAddress string
	BaseURL       string
	Scheme        string
)

const (
	defaultServerAddress = "localhost:8080"
	defaultScheme        = "http"
)

func LoadConfig() {
	var ok bool

	ServerAddress, ok = os.LookupEnv("SERVER_ADDRESS")
	if !ok {
		ServerAddress = defaultServerAddress
	}

	BaseURL = os.Getenv("BASE_URL")

	Scheme, ok = os.LookupEnv("SCHEME")
	if !ok {
		Scheme = defaultScheme
	}
}
