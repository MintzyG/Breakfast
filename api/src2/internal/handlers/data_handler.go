package handlers

import (
	"breakfast/internal/models"
	"breakfast/internal/services"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func GetUserFromContext(ctx context.Context) *models.UserClaims {
	claims, ok := ctx.Value("user").(*models.UserClaims)
	if !ok {
		return nil
	}
	return claims
}

type DataHandler struct {
  DataService *services.DataService
}

func NewDataHandler(service *services.DataService) *DataHandler {
  return &DataHandler{DataService: service}
}

func (h *DataHandler) HelloMe(w http.ResponseWriter, r *http.Request) {
	userClaims := GetUserFromContext(r.Context())
	if userClaims == nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Context came back empty",
		})
		return
	}
  log.Println(userClaims)

  id, err := uuid.Parse(userClaims.ID)
  if err != nil {
    w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "UUID couldn't be parsed",
		})
		return
  }

  name, err := h.DataService.Me(id)
  if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "User not found or malformed",
		})
		return
  }

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Access granted",
    "greet": "Hello, " + name + "!",
		"userID":  userClaims.ID,
	})
}

