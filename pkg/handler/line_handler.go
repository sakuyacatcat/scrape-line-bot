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

type HttpRequestHandler http.Handler

type lineBotHandler struct {
	bot *linebot.Client
}

func NewLineBotHandler() (HttpRequestHandler, error) {
	bot, err := linebot.New(secret, token)
	if err != nil {
		return nil, err
	}
	return &lineBotHandler{bot}, nil
}

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

func (h *lineBotHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	events, err := h.bot.ParseRequest(r)
	if err != nil {
		log.Print(err)
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				if _, err = h.bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	}
}
