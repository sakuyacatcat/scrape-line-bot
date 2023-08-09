package main

import (
	"net/http"

	"github.com/sakuyacatcat/scrape-line-bot/pkg/handler"
)

func main() {
	http.HandleFunc("/", handler.HelloHandler)
	http.ListenAndServe(":8080", nil)
}
