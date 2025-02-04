package handlers

import (
	"breakfast/internal/models"
	"breakfast/internal/services"
	u "breakfast/internal/utilities"
	"strconv"

	"encoding/json"
	"net/http"
)

type YogurtHandler struct {
	Yogurt *services.YogurtService
}

func NewYogurtHandler(service *services.YogurtService) *YogurtHandler {
	return &YogurtHandler{Yogurt: service}
}

func (h *YogurtHandler) Create(w http.ResponseWriter, r *http.Request) {
	userClaims := u.GetUserFromContext(r.Context())
	if userClaims == nil {
		u.Send(w, "Error: user context is empty", nil, http.StatusInternalServerError)
		return
	}

	user_id, err := u.ParseUUID(w, userClaims.ID)
	if err != nil {
		return
	}

	var data models.Yogurt
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		u.Send(w, err.Error(), nil, http.StatusConflict)
		return
	}

	msg, err := models.ValidateModel(data)
	if err != nil {
		u.Send(w, "Invalid request", msg, http.StatusBadRequest)
		return
	}

	err = h.Yogurt.Create(user_id, &data)
	if err != nil {
		u.Send(w, "Could not create task: "+err.Error(), data, http.StatusInternalServerError)
		return
	}

	u.Send(w, "", data, http.StatusCreated)
}

func (h *YogurtHandler) GetByID(w http.ResponseWriter, r *http.Request) {
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

	note, err := h.Yogurt.GetTaskByID(user_id, id)
	if err != nil {
		u.Send(w, "Error retrieving task:"+err.Error(), note, http.StatusInternalServerError)
		return
	}

	u.Send(w, "", note, http.StatusOK)
}

func (h *YogurtHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	userClaims := u.GetUserFromContext(r.Context())
	if userClaims == nil {
		u.Send(w, "Error: user context is empty", nil, http.StatusInternalServerError)
		return
	}

	user_id, err := u.ParseUUID(w, userClaims.ID)
	if err != nil {
		return
	}

	notes, err := h.Yogurt.GetUserTasks(user_id)
	if err != nil {
		u.Send(w, "Error retrieving task:"+err.Error(), notes, http.StatusInternalServerError)
		return
	}

	u.Send(w, "", notes, http.StatusOK)
}

func (h *YogurtHandler) Update(w http.ResponseWriter, r *http.Request) {
	userClaims := u.GetUserFromContext(r.Context())
	if userClaims == nil {
		u.Send(w, "Error: user context is empty", nil, http.StatusInternalServerError)
		return
	}

	user_id, err := u.ParseUUID(w, userClaims.ID)
	if err != nil {
		return
	}

	var data models.Yogurt
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

	data.TaskID = id
	err, note := h.Yogurt.UpdateTask(user_id, &data)
	if err != nil {
		u.Send(w, "Error updating task:"+err.Error(), note, http.StatusInternalServerError)
		return
	}

	u.Send(w, "", note, http.StatusOK)
}

func (h *YogurtHandler) UpdateCompleted(w http.ResponseWriter, r *http.Request) {
	userClaims := u.GetUserFromContext(r.Context())
	if userClaims == nil {
		u.Send(w, "Error: user context is empty", nil, http.StatusInternalServerError)
		return
	}

	user_id, err := u.ParseUUID(w, userClaims.ID)
	if err != nil {
		return
	}

	var data models.Yogurt
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		u.Send(w, err.Error(), nil, http.StatusConflict)
		return
	}

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		u.Send(w, "Error reading the ID requested", nil, http.StatusBadRequest)
		return
	}

	data.TaskID = id
	err, note := h.Yogurt.UpdateCompleted(user_id, &data)
	if err != nil {
		u.Send(w, "Error updating task:"+err.Error(), note, http.StatusInternalServerError)
		return
	}

	u.Send(w, "", note, http.StatusOK)
}

func (h *YogurtHandler) Delete(w http.ResponseWriter, r *http.Request) {
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

	err = h.Yogurt.DeleteTask(user_id, id)
	if err != nil {
		u.Send(w, "Error retrieving task:"+err.Error(), nil, http.StatusInternalServerError)
		return
	}

	u.Send(w, "DELETED", nil, http.StatusOK)
}
