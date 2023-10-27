package handler

import (
	"github.com/MarselBissengaliyev/gochat/pkg/model"
	"github.com/MarselBissengaliyev/gochat/pkg/service"
	"golang.org/x/net/websocket"
)

func (h *Handler) Handle(hub *service.HubService) websocket.Handler {
	
	return func(ws *websocket.Conn) {
		go hub.Run()

		hub.AddClientChan <- ws

		for {
			var m model.Message
			err := websocket.JSON.Receive(ws, &m)
			if err != nil {
				hub.BroadcastChan <- m
				hub.RemoveClient(ws)
				return
			}
			hub.BroadcastChan <- m
		}
	}
}
