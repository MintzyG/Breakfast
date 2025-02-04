package handlers

import (
	"breakfast/internal/services"
	u "breakfast/internal/utilities"
	"net/http"
)

type DataHandler struct {
	DataService *services.DataService
}

func NewDataHandler(service *services.DataService) *DataHandler {
	return &DataHandler{DataService: service}
}

func (h *DataHandler) HelloMe(w http.ResponseWriter, r *http.Request) {
	userClaims := u.GetUserFromContext(r.Context())
	if userClaims == nil {
		u.Send(w, "Error: user context is empty", nil, http.StatusInternalServerError)
		return
	}

	id, err := u.ParseUUID(w, userClaims.ID)
	if err != nil {
		return
	}

	name, err := h.DataService.Me(id)
	if err != nil {
		u.Send(w, "User not found or malformed", nil, http.StatusUnauthorized)
		return
	}

	u.Send(w, "Access Granted!", map[string]string{"greet": "Hello" + name + "!", "userID": userClaims.ID}, http.StatusOK)
}
