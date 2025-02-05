package handlers

import (
	"breakfast/internal/models"
	"breakfast/internal/services"
	u "breakfast/internal/utilities"
	"strconv"

	"encoding/json"
	"net/http"
)

type ToastHandler struct {
	Toast *services.ToastService
}

func NewToastHandler(service *services.ToastService) *ToastHandler {
	return &ToastHandler{Toast: service}
}

func (h *ToastHandler) Create(w http.ResponseWriter, r *http.Request) {
	userClaims := u.GetUserFromContext(r.Context())
	if userClaims == nil {
		u.Send(w, "Error: user context is empty", nil, http.StatusInternalServerError)
		return
	}

	user_id, err := u.ParseUUID(w, userClaims.ID)
	if err != nil {
		return
	}

	var data models.ToastSession
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		u.Send(w, err.Error(), nil, http.StatusConflict)
		return
	}

	msg, err := models.ValidateModel(data)
	if err != nil {
		u.Send(w, "Invalid request", msg, http.StatusBadRequest)
		return
	}

	err = h.Toast.Create(user_id, &data)
	if err != nil {
		u.Send(w, "Could not create session: "+err.Error(), data, http.StatusInternalServerError)
		return
	}

	u.Send(w, "", data, http.StatusCreated)
}

func (h *ToastHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	userClaims := u.GetUserFromContext(r.Context())
	if userClaims == nil {
		u.Send(w, "Error: user context is empty", nil, http.StatusInternalServerError)
		return
	}

	user_id, err := u.ParseUUID(w, userClaims.ID)
	if err != nil {
		return
	}

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		u.Send(w, "Error reading the ID requested", nil, http.StatusBadRequest)
		return
	}

	session, err := h.Toast.GetByID(user_id, id)
	if err != nil {
		u.Send(w, "Error retrieving session:"+err.Error(), session, http.StatusInternalServerError)
		return
	}

	u.Send(w, "", session, http.StatusOK)
}

func (h *ToastHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	userClaims := u.GetUserFromContext(r.Context())
	if userClaims == nil {
		u.Send(w, "Error: user context is empty", nil, http.StatusInternalServerError)
		return
	}

	user_id, err := u.ParseUUID(w, userClaims.ID)
	if err != nil {
		return
	}

	sessions, err := h.Toast.GetAll(user_id)
	if err != nil {
		u.Send(w, "Error retrieving session:"+err.Error(), sessions, http.StatusInternalServerError)
		return
	}

	u.Send(w, "", sessions, http.StatusOK)
}

func (h *ToastHandler) Update(w http.ResponseWriter, r *http.Request) {
	userClaims := u.GetUserFromContext(r.Context())
	if userClaims == nil {
		u.Send(w, "Error: user context is empty", nil, http.StatusInternalServerError)
		return
	}

	user_id, err := u.ParseUUID(w, userClaims.ID)
	if err != nil {
		return
	}

	var data models.ToastSession
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		u.Send(w, err.Error(), nil, http.StatusConflict)
		return
	}

	msg, err := models.ValidateModel(data)
	if err != nil {
		u.Send(w, "Invalid request", msg, http.StatusBadRequest)
		return
	}

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		u.Send(w, "Error reading the ID requested", nil, http.StatusBadRequest)
		return
	}

	data.SessionID = id
	err, session := h.Toast.Update(user_id, &data)
	if err != nil {
		u.Send(w, "Error updating session:"+err.Error(), session, http.StatusInternalServerError)
		return
	}

	u.Send(w, "", session, http.StatusOK)
}

func (h *ToastHandler) Delete(w http.ResponseWriter, r *http.Request) {
	userClaims := u.GetUserFromContext(r.Context())
	if userClaims == nil {
		u.Send(w, "Error: user context is empty", nil, http.StatusInternalServerError)
		return
	}

	user_id, err := u.ParseUUID(w, userClaims.ID)
	if err != nil {
		return
	}

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		u.Send(w, "Error reading the ID requested", nil, http.StatusBadRequest)
		return
	}

	err = h.Toast.Delete(user_id, id)
	if err != nil {
		u.Send(w, "Error retrieving session:"+err.Error(), nil, http.StatusInternalServerError)
		return
	}

	u.Send(w, "DELETED", nil, http.StatusOK)
}
