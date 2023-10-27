package service

import (
	"net/http"

	"github.com/MarselBissengaliyev/gochat/pkg/model"
	"golang.org/x/net/websocket"
)

type Service struct {
	Hub
	Server
}

func NewService(mux *http.ServeMux) *Service {
	return &Service{
		Hub:    NewHub(),
		Server: NewServer(mux),
	}
}

type Hub interface {
	Run()
	AddClient(conn *websocket.Conn)
	RemoveClient(conn *websocket.Conn)
	BroadcastMessage(m model.Message)
	HandleClientActions(ws *websocket.Conn)
}

type Server interface {
	WebSocketHandler(h *HubService) websocket.Handler
}
