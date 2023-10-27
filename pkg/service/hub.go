package service

import (
	"fmt"

	"github.com/MarselBissengaliyev/gochat/pkg/model"
	"golang.org/x/net/websocket"
)

type HubService struct {
	clients          map[string]*websocket.Conn
	AddClientChan    chan *websocket.Conn
	removeClientChan chan *websocket.Conn
	BroadcastChan    chan model.Message
}

func NewHub() *HubService {
	return &HubService{
		clients:          make(map[string]*websocket.Conn),
		AddClientChan:    make(chan *websocket.Conn),
		removeClientChan: make(chan *websocket.Conn),
		BroadcastChan:    make(chan model.Message),
	}
}

func (h *HubService) Run() {
	for {
		select {
		case conn := <-h.AddClientChan:
			h.AddClient(conn)
		case conn := <-h.removeClientChan:
			h.RemoveClient(conn)
		case m := <-h.BroadcastChan:
			h.BroadcastMessage(m)
		}
	}
}

func (h *HubService) AddClient(conn *websocket.Conn) {
	h.clients[conn.RemoteAddr().String()] = conn
}

func (h *HubService) RemoveClient(conn *websocket.Conn) {
	delete(h.clients, conn.LocalAddr().String())
}

func (h *HubService) BroadcastMessage(m model.Message) {
	for _, conn := range h.clients {
		err := websocket.JSON.Send(conn, m)

		if err != nil {
			fmt.Println("Error broadcasting message: ", err)
			return
		}
	}
}

func (h *HubService) HandleClientActions(ws *websocket.Conn) {
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
