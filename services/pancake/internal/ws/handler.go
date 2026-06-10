package ws

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type Handler struct {
	hub *Hub
}

func NewHandler(hub *Hub) *Handler {
	return &Handler{hub: hub}
}

func (h *Handler) Connect(w http.ResponseWriter, r *http.Request) {
	docID := chi.URLParam(r, "docId")

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("[ws] upgrade error: %v", err)
		return
	}

	client := &Client{
		hub:   h.hub,
		conn:  conn,
		docID: docID,
		send:  make(chan []byte, 256),
	}

	h.hub.register <- client

	go client.writePump()
	go client.readPump()
}
