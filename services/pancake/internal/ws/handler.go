package ws

import (
	"net/http"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Connect(w http.ResponseWriter, r *http.Request) {}
