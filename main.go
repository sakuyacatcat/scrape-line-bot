package main

import (
	"log"
	"net/http"

	"github.com/sakuyacatcat/scrape-line-bot/pkg/handler"
)

func main() {
	handler, err := handler.NewLineBotHandler()
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", handler.ServeHTTP)
	http.ListenAndServe(":8080", nil)
}
