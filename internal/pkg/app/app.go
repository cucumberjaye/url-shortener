package app

import (
	"fmt"
	"github.com/cucumberjaye/url-shortener/configs"
	"github.com/cucumberjaye/url-shortener/internal/app/handler"
	"github.com/cucumberjaye/url-shortener/internal/app/repository/localstore"
	ps "github.com/cucumberjaye/url-shortener/internal/app/repository/postrgresdb"
	"github.com/cucumberjaye/url-shortener/internal/app/service"
	"github.com/cucumberjaye/url-shortener/internal/app/service/auth"
	"github.com/cucumberjaye/url-shortener/internal/app/service/hexshortener"
	"github.com/cucumberjaye/url-shortener/internal/app/service/shortenerlogsinfo"
	"github.com/cucumberjaye/url-shortener/pkg/postgres"
	"github.com/go-chi/chi"
	"net/http"
)

type App struct {
	mux *chi.Mux
}

func New() (*App, error) {
	configs.LoadConfig()

	var repos service.URLLogs
	var err error

	if configs.DataBaseDSN != "" {
		pSQL, err := postgres.New()
		if err != nil {
			return nil, err
		}
		repos = ps.NewSQLStore(pSQL)
	} else {
		repos, err = localstore.NewShortenerDB(configs.FileStoragePath)
		if err != nil {
			return nil, err
		}
	}
	serviceURL, err := hexshortener.NewShortenerService(repos)
	if err != nil {
		return nil, err
	}
	logsService := shortenerlogsinfo.NewURLLogsInfo(repos)
	authService := auth.New()

	handlers := handler.NewHandler(serviceURL, logsService, authService)

	mux := chi.NewMux()
	mux.Mount("/", handlers.InitRoutes())

	app := &App{mux: mux}

	return app, nil
}

func (a *App) Run() error {
	fmt.Println("server running")

	return http.ListenAndServe(configs.ServerAddress, a.mux)
}
