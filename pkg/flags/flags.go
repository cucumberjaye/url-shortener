package flags

import (
	"flag"
)

var (
	ServerAddress   string
	BaseURL         string
	FileStoragePath string
	DataBaseDSN     string
)

const (
	defaultServerAddress = "localhost:8080"
)

func InitFlags() {
	flag.StringVar(&ServerAddress, "a", defaultServerAddress, "server address to listen on")
	flag.StringVar(&BaseURL, "b", "", "base address for get method result")
	flag.StringVar(&FileStoragePath, "f", "", "file storage path")
	flag.StringVar(&DataBaseDSN, "d", "", "database URL")

	flag.Parse()
}
