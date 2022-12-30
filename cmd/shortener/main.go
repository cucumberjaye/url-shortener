package main

import (
	app_ "github.com/cucumberjaye/url-shortener/internal/app"
	"github.com/cucumberjaye/url-shortener/pkg/logger"
	"net/http"
)

func main() {
	logger.New()
	app := app_.New()
	logger.ErrorLogger.Fatal(http.ListenAndServe(":8080", app.Handlers.InitRoutes()))
}
