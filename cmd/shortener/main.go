package main

import (
	"github.com/cucumberjaye/url-shortener/internal/app/handler"
	"github.com/cucumberjaye/url-shortener/internal/app/repository"
	"github.com/cucumberjaye/url-shortener/internal/app/service"
	"log"
	"net/http"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	http.HandleFunc("/", handlers.Shortener)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
