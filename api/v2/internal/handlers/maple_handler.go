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

	note, err := h.Maple.GetByID(user_id, id)
	if err != nil {
		u.Send(w, "Error retrieving habit:"+err.Error(), note, http.StatusInternalServerError)
		return
	}

	u.Send(w, "", note, http.StatusOK)
}

func (h *MapleHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	userClaims := u.GetUserFromContext(r.Context())
	if userClaims == nil {
		u.Send(w, "Error: user context is empty", nil, http.StatusInternalServerError)
		return
	}

	user_id, err := u.ParseUUID(w, userClaims.ID)
	if err != nil {
		return
	}

	notes, err := h.Maple.GetAll(user_id)
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
	err, note := h.Maple.Update(user_id, &data)
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

	err = h.Maple.Delete(user_id, id)
	if err != nil {
		u.Send(w, "Error retrieving habit:"+err.Error(), nil, http.StatusInternalServerError)
		return
	}

	u.Send(w, "DELETED", nil, http.StatusOK)
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

func (h *MapleHandler) GetDay(w http.ResponseWriter, r *http.Request) {
	userClaims := u.GetUserFromContext(r.Context())
	if userClaims == nil {
		u.Send(w, "Error: user context is empty", nil, http.StatusInternalServerError)
		return
	}

	userID, err := u.ParseUUID(w, userClaims.ID)
	if err != nil {
		return
	}

	habitIDStr := r.PathValue("id")
	dayIDStr := r.PathValue("day_id")

	habitID, err := strconv.Atoi(habitIDStr)
	if err != nil {
		u.Send(w, "Invalid habit ID", nil, http.StatusBadRequest)
		return
	}

	dayID, err := strconv.Atoi(dayIDStr)
	if err != nil {
		u.Send(w, "Invalid day ID", nil, http.StatusBadRequest)
		return
	}

	mapleDay, err := h.Maple.GetDay(userID, habitID, dayID)
	if err != nil {
		u.Send(w, "Could not retrieve MapleDay: "+err.Error(), nil, http.StatusNotFound)
		return
	}

	u.Send(w, "", mapleDay, http.StatusOK)
}

func (h *MapleHandler) UpdateDay(w http.ResponseWriter, r *http.Request) {
	userClaims := u.GetUserFromContext(r.Context())
	if userClaims == nil {
		u.Send(w, "Error: user context is empty", nil, http.StatusInternalServerError)
		return
	}

	userID, err := u.ParseUUID(w, userClaims.ID)
	if err != nil {
		return
	}

	habitIDStr := r.PathValue("id")
	dayIDStr := r.PathValue("day_id")

	habitID, err := strconv.Atoi(habitIDStr)
	if err != nil {
		u.Send(w, "Invalid habit ID", nil, http.StatusBadRequest)
		return
	}

	dayID, err := strconv.Atoi(dayIDStr)
	if err != nil {
		u.Send(w, "Invalid day ID", nil, http.StatusBadRequest)
		return
	}

	var data models.MapleDay
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		u.Send(w, "Invalid request body", nil, http.StatusBadRequest)
		return
	}

	msg, err := models.ValidateModel(data)
	if err != nil {
		u.Send(w, "Invalid request", msg, http.StatusBadRequest)
		return
	}

	updatedDay, err := h.Maple.UpdateDay(userID, habitID, dayID, &data)
	if err != nil {
		u.Send(w, "Could not update MapleDay: "+err.Error(), nil, http.StatusInternalServerError)
		return
	}

	u.Send(w, "", updatedDay, http.StatusOK)
}

func (h *MapleHandler) DeleteDay(w http.ResponseWriter, r *http.Request) {
	userClaims := u.GetUserFromContext(r.Context())
	if userClaims == nil {
		u.Send(w, "Error: user context is empty", nil, http.StatusInternalServerError)
		return
	}

	userID, err := u.ParseUUID(w, userClaims.ID)
	if err != nil {
		return
	}

	habitIDStr := r.PathValue("id")
	dayIDStr := r.PathValue("day_id")

	habitID, err := strconv.Atoi(habitIDStr)
	if err != nil {
		u.Send(w, "Invalid habit ID", nil, http.StatusBadRequest)
		return
	}

	dayID, err := strconv.Atoi(dayIDStr)
	if err != nil {
		u.Send(w, "Invalid day ID", nil, http.StatusBadRequest)
		return
	}

	err = h.Maple.DeleteDay(userID, habitID, dayID)
	if err != nil {
		u.Send(w, "Could not delete MapleDay: "+err.Error(), nil, http.StatusInternalServerError)
		return
	}

	u.Send(w, "MapleDay deleted successfully", nil, http.StatusOK)
}
