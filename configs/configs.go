package configs

import "os"

var (
	ServerAddress   string
	BaseURL         string
	Scheme          string
	FileStoragePath string
)

const (
	defaultServerAddress = "localhost:8080"
	defaultScheme        = "http"
)

func LoadConfig() {
	ServerAddress = lookUpOrSetDefault("SERVER_ADDRESS", defaultServerAddress)
	Scheme = lookUpOrSetDefault("SCHEME", defaultScheme)
	BaseURL = os.Getenv("BASE_URL")
	FileStoragePath = os.Getenv("FILE_STORAGE_PATH")
}

func lookUpOrSetDefault(name, defaultValue string) string {
	out, ok := os.LookupEnv(name)
	if !ok {
		return defaultValue
	}

	return out
}
