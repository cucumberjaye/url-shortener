package main

import (
	"github.com/cucumberjaye/url-shortener/configs"
	app_ "github.com/cucumberjaye/url-shortener/internal/pkg/app"
	"github.com/cucumberjaye/url-shortener/pkg/logger"
	"net/http"
)

func main() {
	logger.New()
	app, err := app_.New()
	if err != nil {
		logger.ErrorLogger.Fatal(err)
	}
	logger.ErrorLogger.Fatal(http.ListenAndServe(configs.ServerAddress, app.Handlers.InitRoutes()))
}
