package configs

import (
	"github.com/cucumberjaye/url-shortener/pkg/flags"
	"os"
)

var (
	ServerAddress   string
	BaseURL         string
	Scheme          string
	FileStoragePath string
)

const (
	defaultScheme = "http"
)

func LoadConfig() {
	flags.InitFlags()

	ServerAddress = lookUpOrSetDefault("SERVER_ADDRESS", flags.ServerAddress)
	Scheme = lookUpOrSetDefault("SCHEME", defaultScheme)
	BaseURL = lookUpOrSetDefault("BASE_URL", flags.BaseURL)
	FileStoragePath = lookUpOrSetDefault("FILE_STORAGE_PATH", flags.FileStoragePath)
}

func lookUpOrSetDefault(name, defaultValue string) string {
	out, ok := os.LookupEnv(name)
	if !ok {
		return defaultValue
	}

	return out
}
