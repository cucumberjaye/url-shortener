package main

import (
	"github.com/cucumberjaye/url-shortener/internal/pkg/app"
	"github.com/cucumberjaye/url-shortener/pkg/logger"
)

func main() {
	logger.New()

	shortener, err := app.New()
	if err != nil {
		logger.ErrorLogger.Fatal(err)
	}

	err = shortener.Run()
	if err != nil {
		logger.ErrorLogger.Fatal(err)
	}
}
