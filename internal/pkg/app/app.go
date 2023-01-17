package app

import (
	"github.com/cucumberjaye/url-shortener/configs"
	"github.com/cucumberjaye/url-shortener/internal/app/handler"
	"github.com/cucumberjaye/url-shortener/internal/app/repository/localstore"
	"github.com/cucumberjaye/url-shortener/internal/app/service/hexshortener"
	"github.com/cucumberjaye/url-shortener/internal/app/service/shortenerlogsinfo"
)

type App struct {
	Handlers *handler.Handler
}

func New() (*App, error) {
	configs.LoadConfig()
	repos, err := localstore.NewShortenerDB(configs.FileStoragePath)
	if err != nil {
		return nil, err
	}
	serviceURL := hexshortener.NewShortenerService(repos)
	logsService := shortenerlogsinfo.NewURLLogsInfo(repos)
	handlers := handler.NewHandler(serviceURL, logsService)
	app := &App{Handlers: handlers}
	return app, nil
}
