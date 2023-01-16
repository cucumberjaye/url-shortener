package configs

import "os"

var (
	ServerAddress string
	BaseURL       string
	Protocol      string
)

const (
	defaultServerAddress = "localhost:8080"
	defaultProtocol      = "http"
)

func LoadConfig() {
	var ok bool

	ServerAddress, ok = os.LookupEnv("SERVER_ADDRESS")
	if !ok {
		ServerAddress = defaultServerAddress
	}

	BaseURL = os.Getenv("BASE_URL")

	Protocol, ok = os.LookupEnv("PROTOCOL")
	if !ok {
		Protocol = defaultProtocol
	}
}
