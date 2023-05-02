package configs

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/cucumberjaye/url-shortener/pkg/flags"
)

// переменные для считывания env значений
var (
	ServerAddress   string          // адрес сервера
	BaseURL         string          // базовый адрес результирующего сокращённого URL
	Scheme          string = "http" // http или https
	FileStoragePath string          // путь к файлу для хранения данных
	SigningKey      string          // для генерации токена авторизации
	DataBaseDSN     string          // для подключения к postgreSQL
	EnableHTTPS     bool   = false  // https
	Config          string          // json file
)

// значения пол умолчанию
const (
	defaultScheme     = "http"
	defaultSigningKey = "qwerty1234"
)

// для парсинга значений из файла конфигураций
type configJSON struct {
	ServerAddress   string `json:"server_address"`
	BaseUrl         string `json:"base_url"`
	FileStoragePath string `json:"file_storage_path"`
	DatabaseDsn     string `json:"database_dsn"`
	EnableHTTPS     bool   `json:"enable_https"`
}

// LoadConfig устанавливает переменные из env
func LoadConfig() error {
	flags.InitFlags()

	ServerAddress = lookUpOrSetDefault("SERVER_ADDRESS", flags.ServerAddress)
	BaseURL = lookUpOrSetDefault("BASE_URL", flags.BaseURL)
	FileStoragePath = lookUpOrSetDefault("FILE_STORAGE_PATH", flags.FileStoragePath)
	SigningKey = lookUpOrSetDefault("SIGNING_KEY", defaultSigningKey)
	DataBaseDSN = lookUpOrSetDefault("DATABASE_DSN", flags.DataBaseDSN)
	Config = os.Getenv("CONFIG")

	env, ok := os.LookupEnv("ENABLE_HTTPS")
	if ok && env != "false" {
		EnableHTTPS = true
		Scheme = "https"
	}

	return readConfigFile()
}

// lookUpOrSetDefault если нет env значения, устанавливает значение по умолчанию
func lookUpOrSetDefault(name, defaultValue string) string {
	out, ok := os.LookupEnv(name)
	if !ok {
		return defaultValue
	}

	return out
}

// читаем и подставляем данные из файла конфигураций, наименьший приоритет
func readConfigFile() error {
	if len(Config) != 0 {
		data, err := os.ReadFile(Config)
		if err != nil {
			return fmt.Errorf("read config file failed with error: %w", err)
		}

		var cfg configJSON
		err = json.Unmarshal(data, &cfg)
		if err != nil {
			return fmt.Errorf("json file parse failed with error: %w", err)
		}

		if len(ServerAddress) == 0 {
			ServerAddress = cfg.ServerAddress
		}
		if len(BaseURL) == 0 {
			BaseURL = cfg.BaseUrl
		}
		if len(FileStoragePath) == 0 {
			FileStoragePath = cfg.FileStoragePath
		}
		if len(DataBaseDSN) == 0 {
			DataBaseDSN = cfg.DatabaseDsn
		}
		if EnableHTTPS {
			EnableHTTPS = cfg.EnableHTTPS
		}
	}

	return nil
}
