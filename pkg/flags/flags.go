package flags

import (
	"flag"
)

var (
	ServerAddress   string // адрес сервера
	BaseURL         string // базовый адрес результирующего сокращённого URL
	FileStoragePath string // путь к файлу для хранения данных
	DataBaseDSN     string // для подключения к postgreSQL
)

const (
	defaultServerAddress = "localhost:8080"
)

// Обрабатывает флаги при запуске
func InitFlags() {
	flag.StringVar(&ServerAddress, "a", defaultServerAddress, "server address to listen on")
	flag.StringVar(&BaseURL, "b", "", "base address for get method result")
	flag.StringVar(&FileStoragePath, "f", "", "file storage path")
	flag.StringVar(&DataBaseDSN, "d", "", "database URL")

	flag.Parse()
}
