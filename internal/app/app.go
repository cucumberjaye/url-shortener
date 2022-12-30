package app

import (
	"github.com/cucumberjaye/url-shortener/internal/handler"
	"github.com/cucumberjaye/url-shortener/internal/repository/localstore"
	"github.com/cucumberjaye/url-shortener/internal/service/hexshortener"
	"github.com/cucumberjaye/url-shortener/internal/service/shortenerlogsinfo"
)

type App struct {
	Handlers *handler.Handler
}

func New() *App {
	repos := localstore.NewShortenerDB()
	serviceURL := hexshortener.NewShortenerService(repos)
	logsService := shortenerlogsinfo.NewURLLogsInfo(repos)
	handlers := handler.NewHandler(serviceURL, logsService)
	app := &App{Handlers: handlers}
	return app
}
