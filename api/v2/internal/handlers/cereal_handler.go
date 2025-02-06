package handlers

import (
	"breakfast/internal/models"
	"breakfast/internal/services"
	u "breakfast/internal/utilities"

	"encoding/json"
	"net/http"
	"strconv"
)

type CerealHandler struct {
	Cereal *services.CerealService
}

func NewCerealHandler(service *services.CerealService) *CerealHandler {
	return &CerealHandler{Cereal: service}
}

func (h *CerealHandler) Create(w http.ResponseWriter, r *http.Request) {
	userClaims := u.GetUserFromContext(r.Context())
	if userClaims == nil {
		u.Send(w, "Error: user context is empty", nil, http.StatusInternalServerError)
		return
	}

	user_id, err := u.ParseUUID(w, userClaims.ID)
	if err != nil {
		return
	}

	var data models.CerealDay
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		u.Send(w, err.Error(), nil, http.StatusConflict)
		return
	}

	msg, err := models.ValidateModel(data)
	if err != nil {
		u.Send(w, "Invalid request", msg, http.StatusBadRequest)
		return
	}

	err = h.Cereal.Create(user_id, &data)
	if err != nil {
		u.Send(w, "Could not create day: "+err.Error(), data, http.StatusInternalServerError)
		return
	}

	u.Send(w, "", data, http.StatusCreated)
}

func (h *CerealHandler) GetByID(w http.ResponseWriter, r *http.Request) {
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

	day, err := h.Cereal.GetByID(user_id, id)
	if err != nil {
		u.Send(w, "Error retrieving day:"+err.Error(), day, http.StatusInternalServerError)
		return
	}

	u.Send(w, "", day, http.StatusOK)
}

func (h *CerealHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	userClaims := u.GetUserFromContext(r.Context())
	if userClaims == nil {
		u.Send(w, "Error: user context is empty", nil, http.StatusInternalServerError)
		return
	}

	user_id, err := u.ParseUUID(w, userClaims.ID)
	if err != nil {
		return
	}

	dateStr := r.URL.Query().Get("date")
	if dateStr != "" {
		day, err := h.Cereal.GetByDate(user_id, dateStr)
		if err != nil {
			if err.Error() == "No routine created on this day" {
				u.Send(w, "NO-ROUTINE", nil, http.StatusNotFound)
				return
			} else {
				u.Send(w, "Error retrieving day: "+err.Error(), nil, http.StatusInternalServerError)
				return
			}
		}
		u.Send(w, "", day, http.StatusOK)
	} else {
		days, err := h.Cereal.GetAll(user_id)
		if err != nil {
			u.Send(w, "Error retrieving days: "+err.Error(), nil, http.StatusInternalServerError)
			return
		}
		u.Send(w, "", days, http.StatusOK)
	}
}

func (h *CerealHandler) Update(w http.ResponseWriter, r *http.Request) {
	userClaims := u.GetUserFromContext(r.Context())
	if userClaims == nil {
		u.Send(w, "Error: user context is empty", nil, http.StatusInternalServerError)
		return
	}

	user_id, err := u.ParseUUID(w, userClaims.ID)
	if err != nil {
		return
	}

	var data models.CerealDay
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

	data.DayID = id
	err, day := h.Cereal.Update(user_id, &data)
	if err != nil {
		u.Send(w, "Error updating day:"+err.Error(), day, http.StatusInternalServerError)
		return
	}

	u.Send(w, "", day, http.StatusOK)
}

func (h *CerealHandler) Delete(w http.ResponseWriter, r *http.Request) {
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

	err = h.Cereal.Delete(user_id, id)
	if err != nil {
		u.Send(w, "Error retrieving day:"+err.Error(), nil, http.StatusInternalServerError)
		return
	}

	u.Send(w, "DELETED", nil, http.StatusOK)
}

func (h *CerealHandler) CreateActivity(w http.ResponseWriter, r *http.Request) {
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

	var data models.CerealActivity
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		u.Send(w, err.Error(), nil, http.StatusConflict)
		return
	}

	msg, err := models.ValidateModel(data)
	if err != nil {
		u.Send(w, "Invalid request", msg, http.StatusBadRequest)
		return
	}

  activity, err := h.Cereal.CreateActivity(user_id, id, &data)
	if err != nil {
		u.Send(w, "Could not create day: "+err.Error(), activity, http.StatusInternalServerError)
		return
	}

	u.Send(w, "", data, http.StatusCreated)
}

func (h *CerealHandler) GetActivity(w http.ResponseWriter, r *http.Request) {
	userClaims := u.GetUserFromContext(r.Context())
	if userClaims == nil {
		u.Send(w, "Error: user context is empty", nil, http.StatusInternalServerError)
		return
	}

	userID, err := u.ParseUUID(w, userClaims.ID)
	if err != nil {
		return
	}

	dayID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		u.Send(w, "Invalid day ID", nil, http.StatusBadRequest)
		return
	}

	activityID, err := strconv.Atoi(r.PathValue("activity_id"))
	if err != nil {
		u.Send(w, "Invalid activity ID", nil, http.StatusBadRequest)
		return
	}

	activity, err := h.Cereal.GetActivity(userID, dayID, activityID)
	if err != nil {
		u.Send(w, "Activity not found or unauthorized: "+err.Error(), nil, http.StatusNotFound)
		return
	}

	u.Send(w, "", activity, http.StatusOK)
}

func (h *CerealHandler) UpdateActivity(w http.ResponseWriter, r *http.Request) {
	userClaims := u.GetUserFromContext(r.Context())
	if userClaims == nil {
		u.Send(w, "Error: user context is empty", nil, http.StatusInternalServerError)
		return
	}

	userID, err := u.ParseUUID(w, userClaims.ID)
	if err != nil {
		return
	}

	dayIDStr := r.PathValue("id")
	activityIDStr := r.PathValue("activity_id")

	dayID, err := strconv.Atoi(dayIDStr)
	if err != nil {
		u.Send(w, "Invalid day ID", nil, http.StatusBadRequest)
		return
	}

	activityID, err := strconv.Atoi(activityIDStr)
	if err != nil {
		u.Send(w, "Invalid activity ID", nil, http.StatusBadRequest)
		return
	}

	var data models.CerealActivity
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		u.Send(w, "Invalid request body", nil, http.StatusBadRequest)
		return
	}

	msg, err := models.ValidateModel(data)
	if err != nil {
		u.Send(w, "Invalid request", msg, http.StatusBadRequest)
		return
	}

	updatedActivity, err := h.Cereal.UpdateActivity(userID, dayID, activityID, &data)
	if err != nil {
		u.Send(w, "Could not update activity: "+err.Error(), nil, http.StatusInternalServerError)
		return
	}

	u.Send(w, "", updatedActivity, http.StatusOK)
}

func (h *CerealHandler) DeleteActivity(w http.ResponseWriter, r *http.Request) {
	userClaims := u.GetUserFromContext(r.Context())
	if userClaims == nil {
		u.Send(w, "Error: user context is empty", nil, http.StatusInternalServerError)
		return
	}

	userID, err := u.ParseUUID(w, userClaims.ID)
	if err != nil {
		return
	}

	dayID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		u.Send(w, "Invalid day ID", nil, http.StatusBadRequest)
		return
	}

	activityID, err := strconv.Atoi(r.PathValue("activity_id"))
	if err != nil {
		u.Send(w, "Invalid activity ID", nil, http.StatusBadRequest)
		return
	}

	err = h.Cereal.DeleteActivity(userID, dayID, activityID)
	if err != nil {
		u.Send(w, "Could not delete activity: "+err.Error(), nil, http.StatusInternalServerError)
		return
	}

	u.Send(w, "Activity deleted", nil, http.StatusOK)
}
