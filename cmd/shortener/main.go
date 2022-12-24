package main

import (
	"github.com/cucumberjaye/url-shortener/internal/app/handler"
	"github.com/cucumberjaye/url-shortener/internal/app/repository/localstore"
	"github.com/cucumberjaye/url-shortener/internal/app/service/hexshortener"
	"log"
	"net/http"
)

func main() {
	db := localstore.NewShortenerDB()
	services := hexshortener.NewShortenerService(db)
	handlers := handler.NewHandler(services)
	log.Fatal(http.ListenAndServe(":8080", handlers.InitRoutes()))
}
