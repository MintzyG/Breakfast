package handlers

import (
	"breakfast/internal/models"
	"breakfast/internal/services"
	u "breakfast/internal/utilities"
	"encoding/json"
	"net/http"
)

type AuthHandler struct {
	AuthService *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{AuthService: service}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var data models.UserRegister
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		u.Send(w, err.Error(), nil, http.StatusConflict)
		return
	}

	msg, err := models.ValidateModel(data)
	if err != nil {
		u.Send(w, "Invalid request", msg, http.StatusBadRequest)
		return
	}

	if err := h.AuthService.Register(data.Email, data.Password, data.Name); err != nil {
		u.Send(w, err.Error(), nil, http.StatusConflict)
		return
	}

	u.Send(w, "Created user successfully", nil, http.StatusCreated)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var data models.UserLogin
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		u.Send(w, "Invalid Input", err.Error(), http.StatusBadRequest)
		return
	}

	msg, err := models.ValidateModel(data)
	if err != nil {
		u.Send(w, "Invalid request", msg, http.StatusBadRequest)
		return
	}

	token, err := h.AuthService.Login(data.Email, data.Password)
	if err != nil {
		u.Send(w, "Invalid email or password", err.Error(), http.StatusUnauthorized)
		return
	}

	u.Send(w, "Login successful", map[string]string{"token": token}, http.StatusOK)
}
