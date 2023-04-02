package configs

import (
	"os"

	"github.com/cucumberjaye/url-shortener/pkg/flags"
)

var (
	ServerAddress   string
	BaseURL         string
	Scheme          string // http or https
	FileStoragePath string
	SigningKey      string // for token generate
	DataBaseDSN     string
)

const (
	defaultScheme     = "http"
	defaultSigningKey = "qwerty1234"
)

// LoadConfig устанавливает переменные из env
func LoadConfig() {
	flags.InitFlags()

	ServerAddress = lookUpOrSetDefault("SERVER_ADDRESS", flags.ServerAddress)
	Scheme = lookUpOrSetDefault("SCHEME", defaultScheme)
	BaseURL = lookUpOrSetDefault("BASE_URL", flags.BaseURL)
	FileStoragePath = lookUpOrSetDefault("FILE_STORAGE_PATH", flags.FileStoragePath)
	SigningKey = lookUpOrSetDefault("SIGNING_KEY", defaultSigningKey)
	DataBaseDSN = lookUpOrSetDefault("DATABASE_DSN", flags.DataBaseDSN)
}

// lookUpOrSetDefault если нет env значения, устанавливает значение по умолчанию
func lookUpOrSetDefault(name, defaultValue string) string {
	out, ok := os.LookupEnv(name)
	if !ok {
		return defaultValue
	}

	return out
}
