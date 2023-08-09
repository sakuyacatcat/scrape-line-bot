package handler

import (
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

var (
	secret string
	token  string
)

func init() {
	s, err := os.LookupEnv("LINE_CHANNEL_SECRET")
	if !err {
		log.Fatal("LINE_CHANNEL_SECRET is not set")
	}
	secret = s

	t, err := os.LookupEnv("LINE_CHANNEL_ACCESS_TOKEN")
	if !err {
		log.Fatal("LINE_CHANNEL_ACCESS_TOKEN is not set")
	}
	token = t
}

func LineBotHandler(w http.ResponseWriter, r *http.Request) {
	bot, err := linebot.New(secret, token)
	if err != nil {
		log.Fatal(err)
	}

	events, err := bot.ParseRequest(r)
	if err != nil {
		log.Print(err)
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	}
}
