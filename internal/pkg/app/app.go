package app

import (
	"fmt"
	"github.com/cucumberjaye/url-shortener/configs"
	"github.com/cucumberjaye/url-shortener/internal/app/handler"
	"github.com/cucumberjaye/url-shortener/internal/app/repository/localstore"
	"github.com/cucumberjaye/url-shortener/internal/app/service/hexshortener"
	"github.com/cucumberjaye/url-shortener/internal/app/service/shortenerlogsinfo"
	"github.com/go-chi/chi"
	"net/http"
)

type App struct {
	mux *chi.Mux
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

	mux := chi.NewMux()
	mux.Mount("/", handlers.InitRoutes())

	app := &App{mux: mux}

	return app, nil
}

func (a *App) Run() error {
	fmt.Println("server running")

	return http.ListenAndServe(configs.ServerAddress, a.mux)
}
