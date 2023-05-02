package app

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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
	"golang.org/x/crypto/acme/autocert"
)

// Структура для запуска приложения
type App struct {
	mux *chi.Mux
}

// создаем Арр
func New() (*App, error) {
	err := configs.LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("load config failed with error: %w", err)
	}

	var keeper repository.Keeper

	if configs.DataBaseDSN != "" {
		var pSQL *sql.DB
		pSQL, err = postgres.New()
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

	var srv *http.Server

	sigint := make(chan os.Signal, 1)
	idleConnsClosed := make(chan struct{})
	signal.Notify(sigint, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	shutdown := func(srv *http.Server) {
		<-sigint
		if err := srv.Shutdown(context.Background()); err != nil {
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}

	if configs.EnableHTTPS {
		manager := &autocert.Manager{
			Cache:  autocert.DirCache("cert"),
			Prompt: autocert.AcceptTOS,
		}

		srv = &http.Server{
			Addr:      configs.ServerAddress,
			Handler:   a.mux,
			TLSConfig: manager.TLSConfig(),
		}
		go shutdown(srv)

		if err := srv.ListenAndServeTLS("", ""); err != http.ErrServerClosed {
			return fmt.Errorf("listen server failed with error: %w", err)
		}
	} else {
		srv = &http.Server{
			Addr:    configs.ServerAddress,
			Handler: a.mux,
		}

		go shutdown(srv)

		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			return fmt.Errorf("listen server failed with error: %w", err)
		}
	}
	<-idleConnsClosed
	return nil
}
