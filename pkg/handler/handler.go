package handler

import (
	"net/http"

	"github.com/MarselBissengaliyev/gochat/pkg/service"
	"golang.org/x/net/websocket"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitHandlers(port string) *http.Server {
	mux := http.NewServeMux()

	mux.Handle("/", websocket.Handler(func(ws *websocket.Conn) {
		h.services.HandleClientActions(ws)
	}))

	srv := http.Server{Addr: ":" + port, Handler: mux}

	return &srv
}
