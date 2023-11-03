package main

import (
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/pkg/errors"
	"github.com/sakuyacatcat/scrape-line-bot/pkg/handler"
)

var (
	secret string
	token  string
)

func init() {
	if err := getEnv(); err != nil {
		log.Fatal(err)
	}
}

func getEnv() error {
	s, err := os.LookupEnv("LINE_CHANNEL_SECRET")
	if !err {
		return errors.New("env LINE_CHANNEL_SECRET is not found")
	}

	secret = s

	t, err := os.LookupEnv("LINE_CHANNEL_ACCESS_TOKEN")
	if !err {
		log.Fatal("env LINE_CHANNEL_ACCESS_TOKEN is not found")
	}

	token = t

	return nil
}

func main() {
	bot, err := linebot.New(secret, token)
	if err != nil {
		log.Printf("failed to get line bot client: %v", err)
	}
	handlers := handler.NewEventHandlerContainer(bot)
	lineHandler := handler.NewLineHandler(bot, handlers)

	http.HandleFunc("/", lineHandler.Handle)
	http.ListenAndServe(":8080", nil)
}
