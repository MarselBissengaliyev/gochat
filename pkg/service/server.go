package service

import (
	"net/http"

	"github.com/MarselBissengaliyev/gochat/pkg/model"
	"golang.org/x/net/websocket"
)

type ServerService struct {
	mux *http.ServeMux
}

func NewServer(mux *http.ServeMux) *ServerService {
	return &ServerService{mux: mux}
}

func (s *ServerService) WebSocketHandler(h *HubService) websocket.Handler {
	return func(ws *websocket.Conn) {
		go h.Run()

		h.AddClientChan <- ws

		for {
			var m model.Message
			err := websocket.JSON.Receive(ws, &m)
			if err != nil {
				h.BroadcastChan <- m
				h.RemoveClient(ws)
				return
			}
			h.BroadcastChan <- m
		}
	}
}
