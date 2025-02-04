package handlers

import (
	"breakfast/internal/models"
	"breakfast/internal/services"
	u "breakfast/internal/utilities"
	"strconv"

	"encoding/json"
	"net/http"
)

type MapleHandler struct {
	Maple *services.MapleService
}

func NewMapleHandler(service *services.MapleService) *MapleHandler {
	return &MapleHandler{Maple: service}
}

func (h *MapleHandler) Create(w http.ResponseWriter, r *http.Request) {
	userClaims := u.GetUserFromContext(r.Context())
	if userClaims == nil {
		u.Send(w, "Error: user context is empty", nil, http.StatusInternalServerError)
		return
	}

	user_id, err := u.ParseUUID(w, userClaims.ID)
	if err != nil {
		return
	}

	var data models.Maple
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		u.Send(w, err.Error(), nil, http.StatusConflict)
		return
	}

	msg, err := models.ValidateModel(data)
	if err != nil {
		u.Send(w, "Invalid request", msg, http.StatusBadRequest)
		return
	}

	err = h.Maple.Create(user_id, &data)
	if err != nil {
		u.Send(w, "Could not create habit: "+err.Error(), data, http.StatusInternalServerError)
		return
	}

	u.Send(w, "", data, http.StatusCreated)
}

func (h *MapleHandler) CreateDay(w http.ResponseWriter, r *http.Request) {
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

	var data models.MapleDay
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		u.Send(w, err.Error(), nil, http.StatusConflict)
		return
	}

	msg, err := models.ValidateModel(data)
	if err != nil {
		u.Send(w, "Invalid request", msg, http.StatusBadRequest)
		return
	}

	data.HabitID = id
	habit, err := h.Maple.CreateDay(user_id, id, &data)
	if err != nil {
		u.Send(w, "Could not create habit: "+err.Error(), habit, http.StatusInternalServerError)
		return
	}

	u.Send(w, "", data, http.StatusCreated)
}

func (h *MapleHandler) GetByID(w http.ResponseWriter, r *http.Request) {
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

	note, err := h.Maple.GetHabitByID(user_id, id)
	if err != nil {
		u.Send(w, "Error retrieving habit:"+err.Error(), note, http.StatusInternalServerError)
		return
	}

	u.Send(w, "", note, http.StatusOK)
}

func (h *MapleHandler) GetHabits(w http.ResponseWriter, r *http.Request) {
	userClaims := u.GetUserFromContext(r.Context())
	if userClaims == nil {
		u.Send(w, "Error: user context is empty", nil, http.StatusInternalServerError)
		return
	}

	user_id, err := u.ParseUUID(w, userClaims.ID)
	if err != nil {
		return
	}

	notes, err := h.Maple.GetUserHabits(user_id)
	if err != nil {
		u.Send(w, "Error retrieving habit:"+err.Error(), notes, http.StatusInternalServerError)
		return
	}

	u.Send(w, "", notes, http.StatusOK)
}

func (h *MapleHandler) Update(w http.ResponseWriter, r *http.Request) {
	userClaims := u.GetUserFromContext(r.Context())
	if userClaims == nil {
		u.Send(w, "Error: user context is empty", nil, http.StatusInternalServerError)
		return
	}

	user_id, err := u.ParseUUID(w, userClaims.ID)
	if err != nil {
		return
	}

	var data models.Maple
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

	data.HabitID = id
	err, note := h.Maple.UpdateHabit(user_id, &data)
	if err != nil {
		u.Send(w, "Error updating habit:"+err.Error(), note, http.StatusInternalServerError)
		return
	}

	u.Send(w, "", note, http.StatusOK)
}

func (h *MapleHandler) Delete(w http.ResponseWriter, r *http.Request) {
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

	err = h.Maple.DeleteHabit(user_id, id)
	if err != nil {
		u.Send(w, "Error retrieving habit:"+err.Error(), nil, http.StatusInternalServerError)
		return
	}

	u.Send(w, "DELETED", nil, http.StatusOK)
}
