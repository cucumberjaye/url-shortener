package app

import (
	"fmt"
	"github.com/cucumberjaye/url-shortener/configs"
	"github.com/cucumberjaye/url-shortener/internal/app/handler"
	"github.com/cucumberjaye/url-shortener/internal/app/repository"
	"github.com/cucumberjaye/url-shortener/internal/app/repository/filestore"
	"github.com/cucumberjaye/url-shortener/internal/app/repository/localstore"
	ps "github.com/cucumberjaye/url-shortener/internal/app/repository/postrgresdb"
	"github.com/cucumberjaye/url-shortener/internal/app/service/deleter"
	"github.com/cucumberjaye/url-shortener/internal/app/service/hexshortener"
	"github.com/cucumberjaye/url-shortener/internal/app/service/shortenerlogsinfo"
	"github.com/cucumberjaye/url-shortener/models"
	"github.com/cucumberjaye/url-shortener/pkg/postgres"
	"github.com/go-chi/chi"
	"net/http"
)

type App struct {
	mux *chi.Mux
}

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

	ch := make(chan []models.DeleteData)

	deleterService := deleter.New(repos, ch)
	go deleterService.Deleting()

	handlers := handler.NewHandler(serviceURL, logsService, ch)

	mux := chi.NewMux()
	mux.Mount("/", handlers.InitRoutes())

	app := &App{
		mux: mux,
	}

	return app, nil
}

func (a *App) Run() error {
	fmt.Println("server running")

	return http.ListenAndServe(configs.ServerAddress, a.mux)
}
