package configs

import (
	"os"

	"github.com/cucumberjaye/url-shortener/pkg/flags"
)

// переменные для считывания env значений
var (
	ServerAddress   string // адрес сервера
	BaseURL         string // базовый адрес результирующего сокращённого URL
	Scheme          string // http или https
	FileStoragePath string // путь к файлу для хранения данных
	SigningKey      string // для генерации токена авторизации
	DataBaseDSN     string // для подключения к postgreSQL
)

// значения пол умолчанию
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
