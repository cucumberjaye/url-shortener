package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cucumberjaye/url-shortener/configs"
	"github.com/cucumberjaye/url-shortener/internal/app/handler"
	"github.com/cucumberjaye/url-shortener/internal/app/repository"
	"github.com/cucumberjaye/url-shortener/internal/app/repository/filestore"
	"github.com/cucumberjaye/url-shortener/internal/app/repository/localstore"
	ps "github.com/cucumberjaye/url-shortener/internal/app/repository/postrgresdb"
	"github.com/cucumberjaye/url-shortener/internal/app/service/hexshortener"
	"github.com/cucumberjaye/url-shortener/internal/app/service/shortenerlogsinfo"
	"github.com/cucumberjaye/url-shortener/internal/app/worker"
	"github.com/cucumberjaye/url-shortener/models"
	"github.com/cucumberjaye/url-shortener/pkg/postgres"
	"github.com/go-chi/chi"
)

// Структура для запуска приложения
type App struct {
	mux *chi.Mux
}

// создаем Арр
func New() (*App, error) {
	configs.LoadConfig()

	var keeper repository.Keeper
	var err error

	if configs.DataBaseDSN != "" {
		pSQL, err := postgres.New()
		if err != nil {
			return nil, err
		}
		keeper = ps.NewSQLStore(pSQL)
	} else if configs.FileStoragePath != "" {
		keeper, err = filestore.New(configs.FileStoragePath)
		if err != nil {
			return nil, err
		}
	}

	repos, err := localstore.NewShortenerDB(keeper)
	if err != nil {
		return nil, err
	}

	serviceURL, err := hexshortener.NewShortenerService(repos)
	if err != nil {
		return nil, err
	}
	logsService := shortenerlogsinfo.NewURLLogsInfo(repos)

	ch := make(chan models.DeleteData)

	workers := worker.New(repos, ch)
	workers.Start(context.Background())

	handlers := handler.NewHandler(serviceURL, logsService, ch)

	mux := chi.NewMux()
	mux.Mount("/", handlers.InitRoutes())

	app := &App{
		mux: mux,
	}

	return app, nil
}

// запускем сервер
func (a *App) Run() error {
	fmt.Println("server running")

	return http.ListenAndServe(configs.ServerAddress, a.mux)
}
