//go:generate mockgen -source=lineHandler.go -destination=../../mock/mock_lineHandler.go -package=mock
package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"
)

type LineHandler interface {
	Handle(http.ResponseWriter, *http.Request)
}

type LineBotClient interface {
	ParseRequest(*http.Request) ([]*linebot.Event, error)
	ReplyMessage(string, ...linebot.SendingMessage) *linebot.ReplyMessageCall
}

type EventHandler interface {
	HandleEvent(*linebot.Event) error
}

type lineHandler struct {
	bot      LineBotClient
	handlers *EventHandlerContainer
}

type MessageEventHandler struct {
	bot LineBotClient
}

type EventHandlerContainer struct {
	handlers map[linebot.EventType]EventHandler
}

func NewLineHandler(bot LineBotClient, handlers *EventHandlerContainer) LineHandler {
	return &lineHandler{bot, handlers}
}

func NewEventHandlerContainer(bot LineBotClient) *EventHandlerContainer {
	return &EventHandlerContainer{
		handlers: map[linebot.EventType]EventHandler{
			linebot.EventTypeMessage: &MessageEventHandler{bot},
		},
	}
}

func (h *lineHandler) Handle(w http.ResponseWriter, r *http.Request) {
	events, err := h.bot.ParseRequest(r)
	if err != nil {
		log.Print(err)
		return
	}

	for _, event := range events {
		handler := h.handlers.GetHandler(event.Type)
		if err := handler.HandleEvent(event); err != nil {
			log.Printf("handleEvent failed: %v", err)
		}
	}
}

func (c *EventHandlerContainer) GetHandler(eventType linebot.EventType) EventHandler {
	return c.handlers[eventType]
}

func (h *MessageEventHandler) HandleEvent(event *linebot.Event) error {
	switch message := event.Message.(type) {
	case *linebot.TextMessage:
		if _, err := h.bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
			return fmt.Errorf("ReplyMessage failed: %v", err)
		}
	}
	return nil
}
