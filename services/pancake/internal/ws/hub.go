package ws

type message struct {
	sender  *Client
	docID   string
	payload []byte
}

type Hub struct {
	rooms      map[string]map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan message
}

func NewHub() *Hub {
	return &Hub{
		rooms:      make(map[string]map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan message),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case c := <-h.register:
			if h.rooms[c.docID] == nil {
				h.rooms[c.docID] = make(map[*Client]bool)
			}
			h.rooms[c.docID][c] = true

		case c := <-h.unregister:
			if room, ok := h.rooms[c.docID]; ok {
				if _, ok := room[c]; ok {
					delete(room, c)
					close(c.send)
					if len(room) == 0 {
						delete(h.rooms, c.docID)
					}
				}
			}

		case msg := <-h.broadcast:
			for c := range h.rooms[msg.docID] {
				if c == msg.sender {
					continue
				}
				select {
				case c.send <- msg.payload:
				default:
					delete(h.rooms[msg.docID], c)
					close(c.send)
				}
			}
		}
	}
}
