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
	handleEvent(*linebot.Event) error
}

type lineHandler struct {
	bot      LineBotClient
	handlers *eventHandlerContainer
}

type messageEventHandler struct {
	bot LineBotClient
}

type eventHandlerContainer struct {
	handlers map[linebot.EventType]EventHandler
}

func NewLineHandler(bot LineBotClient, handlers *eventHandlerContainer) LineHandler {
	return &lineHandler{bot, handlers}
}

func NewEventHandlerContainer(bot LineBotClient) *eventHandlerContainer {
	return &eventHandlerContainer{
		handlers: map[linebot.EventType]EventHandler{
			linebot.EventTypeMessage: &messageEventHandler{bot},
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
		handler := h.handlers.getHandler(event.Type)
		if err := handler.handleEvent(event); err != nil {
			log.Printf("handleEvent failed: %v", err)
		}
	}
}

func (c *eventHandlerContainer) getHandler(eventType linebot.EventType) EventHandler {
	return c.handlers[eventType]
}

func (h *messageEventHandler) handleEvent(event *linebot.Event) error {
	switch message := event.Message.(type) {
	case *linebot.TextMessage:
		if _, err := h.bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
			return fmt.Errorf("ReplyMessage failed: %v", err)
		}
	}
	return nil
}
