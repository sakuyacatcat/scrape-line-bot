//go:generate mockgen -source=lineHandler.go -destination=../../mock/mock_lineHandler.go -package=mock
package handler

import (
	"log"
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"
)

type LineHandler interface {
	Handle(http.ResponseWriter, *http.Request)
}

type lineHandler struct {
	client *linebot.Client
}

func NewLineHandler(c *linebot.Client) *lineHandler {
	return &lineHandler{
		client: c,
	}
}

func (h *lineHandler) Handle(w http.ResponseWriter, r *http.Request) {
	events, err := h.client.ParseRequest(r)
	log.Println(events)
	if err != nil {
		log.Print(err)
		return
	}

	for _, event := range events {
		switch message := event.Message.(type) {
		case *linebot.TextMessage:
			if _, err := h.client.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
				log.Printf("ReplyMessage failed: %v", err)
			}
		}
	}
}
