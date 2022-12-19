package main

import (
	"github.com/cucumberjaye/url-shortener/internal/app"
	"log"
	"net/http"
)

func main() {
	db := app.NewDB()
	services := app.NewService(db)
	handlers := app.NewHandler(services)
	http.HandleFunc("/", handlers.Shortener)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
