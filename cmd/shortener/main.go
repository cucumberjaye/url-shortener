package main

import (
	"fmt"

	"github.com/cucumberjaye/url-shortener/internal/pkg/app"
	"github.com/cucumberjaye/url-shortener/pkg/logger"
	"golang.org/x/sync/errgroup"
)

var (
	buildVersion string = "N/A"
	buildDate    string = "N/A"
	buildCommit  string = "N/A"
)

func main() {
	fmt.Printf("Build version: %s\n", buildVersion)
	fmt.Printf("Build date: %s\n", buildDate)
	fmt.Printf("Build commit: %s\n", buildCommit)

	logger.New()

	shortener, err := app.New()
	if err != nil {
		logger.ErrorLogger.Fatal(err)
	}

	g := new(errgroup.Group)

	g.Go(shortener.Run)
	g.Go(shortener.GRPCRun)

	if err := g.Wait(); err != nil {
		logger.ErrorLogger.Fatal(err)
	}
}
