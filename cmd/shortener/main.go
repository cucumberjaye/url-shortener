package main

import (
	"github.com/cucumberjaye/url-shortener/internal/pkg/app"
	"github.com/cucumberjaye/url-shortener/pkg/logger"
)

func main() {
	logger.New()

	a, err := app.New()
	if err != nil {
		logger.ErrorLogger.Fatal(err)
	}

	err = a.Run()
	if err != nil {
		logger.ErrorLogger.Fatal(err)
	}
}
