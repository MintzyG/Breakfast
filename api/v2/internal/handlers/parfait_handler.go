package handlers

import (
	"breakfast/internal/models"
	"breakfast/internal/services"
	u "breakfast/internal/utilities"
	"strconv"

	"encoding/json"
	"net/http"
)

type ParfaitHandler struct {
	Parfait *services.ParfaitService
}

func NewParfaitHandler(service *services.ParfaitService) *ParfaitHandler {
	return &ParfaitHandler{Parfait: service}
}

func (h *ParfaitHandler) Create(w http.ResponseWriter, r *http.Request) {
	userClaims := u.GetUserFromContext(r.Context())
	if userClaims == nil {
		u.Send(w, "Error: user context is empty", nil, http.StatusInternalServerError)
		return
	}

	user_id, err := u.ParseUUID(w, userClaims.ID)
	if err != nil {
		return
	}

	var data models.ParfaitEvent
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		u.Send(w, err.Error(), nil, http.StatusConflict)
		return
	}

	msg, err := models.ValidateModel(data)
	if err != nil {
		u.Send(w, "Invalid request", msg, http.StatusBadRequest)
		return
	}

	err = h.Parfait.Create(user_id, &data)
	if err != nil {
		u.Send(w, "Could not create event: "+err.Error(), data, http.StatusInternalServerError)
		return
	}

	u.Send(w, "", data, http.StatusCreated)
}

func (h *ParfaitHandler) GetByID(w http.ResponseWriter, r *http.Request) {
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

	event, err := h.Parfait.GetByID(user_id, id)
	if err != nil {
		u.Send(w, "Error retrieving event:"+err.Error(), event, http.StatusInternalServerError)
		return
	}

	u.Send(w, "", event, http.StatusOK)
}

func (h *ParfaitHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	userClaims := u.GetUserFromContext(r.Context())
	if userClaims == nil {
		u.Send(w, "Error: user context is empty", nil, http.StatusInternalServerError)
		return
	}

	user_id, err := u.ParseUUID(w, userClaims.ID)
	if err != nil {
		return
	}

	events, err := h.Parfait.GetAll(user_id)
	if err != nil {
		u.Send(w, "Error retrieving event:"+err.Error(), events, http.StatusInternalServerError)
		return
	}

	u.Send(w, "", events, http.StatusOK)
}

func (h *ParfaitHandler) Update(w http.ResponseWriter, r *http.Request) {
	userClaims := u.GetUserFromContext(r.Context())
	if userClaims == nil {
		u.Send(w, "Error: user context is empty", nil, http.StatusInternalServerError)
		return
	}

	user_id, err := u.ParseUUID(w, userClaims.ID)
	if err != nil {
		return
	}

	var data models.ParfaitEvent
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

	data.EventID = id
	err, event := h.Parfait.Update(user_id, &data)
	if err != nil {
		u.Send(w, "Error updating event:"+err.Error(), event, http.StatusInternalServerError)
		return
	}

	u.Send(w, "", event, http.StatusOK)
}

func (h *ParfaitHandler) Delete(w http.ResponseWriter, r *http.Request) {
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

	err = h.Parfait.Delete(user_id, id)
	if err != nil {
		u.Send(w, "Error retrieving event:"+err.Error(), nil, http.StatusInternalServerError)
		return
	}

	u.Send(w, "DELETED", nil, http.StatusOK)
}

func (h *ParfaitHandler) CreateReminder(w http.ResponseWriter, r *http.Request) {
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

	var data models.ParfaitReminder
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		u.Send(w, err.Error(), nil, http.StatusConflict)
		return
	}

	msg, err := models.ValidateModel(data)
	if err != nil {
		u.Send(w, "Invalid request", msg, http.StatusBadRequest)
		return
	}

	data.EventID = id
	event, err := h.Parfait.CreateReminder(user_id, id, &data)
	if err != nil {
		u.Send(w, "Could not create event: "+err.Error(), event, http.StatusInternalServerError)
		return
	}

	u.Send(w, "", data, http.StatusCreated)
}

func (h *ParfaitHandler) GetReminder(w http.ResponseWriter, r *http.Request) {
	userClaims := u.GetUserFromContext(r.Context())
	if userClaims == nil {
		u.Send(w, "Error: user context is empty", nil, http.StatusInternalServerError)
		return
	}

	userID, err := u.ParseUUID(w, userClaims.ID)
	if err != nil {
		return
	}

	eventIDStr := r.PathValue("id")
	reminderIDStr := r.PathValue("reminder_id")

	eventID, err := strconv.Atoi(eventIDStr)
	if err != nil {
		u.Send(w, "Invalid event ID", nil, http.StatusBadRequest)
		return
	}

	reminderID, err := strconv.Atoi(reminderIDStr)
	if err != nil {
		u.Send(w, "Invalid reminder ID", nil, http.StatusBadRequest)
		return
	}

	reminder, err := h.Parfait.GetReminder(userID, eventID, reminderID)
	if err != nil {
		u.Send(w, "Could not retrieve ParfaitReminder: "+err.Error(), nil, http.StatusNotFound)
		return
	}

	u.Send(w, "", reminder, http.StatusOK)
}

func (h *ParfaitHandler) GetAllReminders(w http.ResponseWriter, r *http.Request) {
	userClaims := u.GetUserFromContext(r.Context())
	if userClaims == nil {
		u.Send(w, "Error: user context is empty", nil, http.StatusInternalServerError)
		return
	}

	userID, err := u.ParseUUID(w, userClaims.ID)
	if err != nil {
		return
	}

	reminders, err := h.Parfait.GetAllReminders(userID)
	if err != nil {
		u.Send(w, "Could not retrieve ParfaitReminder: "+err.Error(), nil, http.StatusNotFound)
		return
	}

	u.Send(w, "", reminders, http.StatusOK)
}

func (h *ParfaitHandler) UpdateReminder(w http.ResponseWriter, r *http.Request) {
	userClaims := u.GetUserFromContext(r.Context())
	if userClaims == nil {
		u.Send(w, "Error: user context is empty", nil, http.StatusInternalServerError)
		return
	}

	userID, err := u.ParseUUID(w, userClaims.ID)
	if err != nil {
		return
	}

	eventIDStr := r.PathValue("id")
	reminderIDStr := r.PathValue("reminder_id")

	eventID, err := strconv.Atoi(eventIDStr)
	if err != nil {
		u.Send(w, "Invalid event ID", nil, http.StatusBadRequest)
		return
	}

	reminderID, err := strconv.Atoi(reminderIDStr)
	if err != nil {
		u.Send(w, "Invalid reminder ID", nil, http.StatusBadRequest)
		return
	}

	var data models.ParfaitReminder
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		u.Send(w, "Invalid request body", nil, http.StatusBadRequest)
		return
	}

	msg, err := models.ValidateModel(data)
	if err != nil {
		u.Send(w, "Invalid request", msg, http.StatusBadRequest)
		return
	}

	updatedReminder, err := h.Parfait.UpdateReminder(userID, eventID, reminderID, &data)
	if err != nil {
		u.Send(w, "Could not update ParfaitReminder: "+err.Error(), nil, http.StatusInternalServerError)
		return
	}

	u.Send(w, "", updatedReminder, http.StatusOK)
}

func (h *ParfaitHandler) DeleteReminder(w http.ResponseWriter, r *http.Request) {
	userClaims := u.GetUserFromContext(r.Context())
	if userClaims == nil {
		u.Send(w, "Error: user context is empty", nil, http.StatusInternalServerError)
		return
	}

	userID, err := u.ParseUUID(w, userClaims.ID)
	if err != nil {
		return
	}

	eventIDStr := r.PathValue("id")
	reminderIDStr := r.PathValue("reminder_id")

	eventID, err := strconv.Atoi(eventIDStr)
	if err != nil {
		u.Send(w, "Invalid event ID", nil, http.StatusBadRequest)
		return
	}

	reminderID, err := strconv.Atoi(reminderIDStr)
	if err != nil {
		u.Send(w, "Invalid reminder ID", nil, http.StatusBadRequest)
		return
	}

	err = h.Parfait.DeleteReminder(userID, eventID, reminderID)
	if err != nil {
		u.Send(w, "Could not delete ParfaitReminder: "+err.Error(), nil, http.StatusInternalServerError)
		return
	}

	u.Send(w, "ParfaitReminder deleted successfully", nil, http.StatusOK)
}
