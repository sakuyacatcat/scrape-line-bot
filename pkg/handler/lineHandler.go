//go:generate mockgen -source=lineHandler.go -destination=../../mock/mock_lineHandler.go -package=mock
package handler

import (
	"log"
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/sakuyacatcat/scrape-line-bot/pkg/controller"
)

type LineHandler interface {
	Handle(http.ResponseWriter, *http.Request)
}

type lineHandler struct {
	client     *linebot.Client
	controller *controller.CoatController
}

func NewLineHandler(cl *linebot.Client, cn *controller.CoatController) *lineHandler {
	return &lineHandler{
		client:     cl,
		controller: cn,
	}
}

func (h *lineHandler) Handle(w http.ResponseWriter, r *http.Request) {
	log.Println("handle start")
	events, err := h.client.ParseRequest(r)
	if err != nil {
		log.Print(err)
		return
	}

	for _, event := range events {
		switch message := event.Message.(type) {
		case *linebot.TextMessage:
			response, err := h.controller.GetCoat(message.Text)
			if err != nil {
				log.Printf("failed to get coat: %v", err)
				return
			}

			if _, err := h.client.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(response)).Do(); err != nil {
				log.Printf("ReplyMessage failed: %v", err)
			}
		}
	}
	log.Println("handle finished")
}
