package categories

import (
	"breakfast/models"
	DB "breakfast/repositories"
	RSP "breakfast/response"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

func createCategory(w http.ResponseWriter, r *http.Request) {
	var c models.Category
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		RSP.SendErrorResponse(w, http.StatusBadRequest, "Error parsing JSON", "JSON_ERROR")
		return
	}

	if c.Title == "" || c.Description == "" {
		RSP.SendErrorResponse(w, http.StatusUnprocessableEntity, "Info fields empty", "MISSING_DATA")
		return
	}

	if c.Color == "" || c.TextColor == "" {
		RSP.SendErrorResponse(w, http.StatusUnprocessableEntity, "Color fields empty", "MISSING_DATA")
		return
	}

	if c.Emoji == "" {
		RSP.SendErrorResponse(w, http.StatusUnprocessableEntity, "Emoji fields empty", "MISSING_DATA")
		return
	}

	claims, ok := r.Context().Value("claims").(*models.UserClaims)
	if !ok {
		RSP.SendErrorResponse(w, http.StatusUnauthorized, "Claims missing", "CLAIMS_ERROR")
		return
	}

	id, err := uuid.Parse(claims.UserID)
	if err != nil {
		RSP.SendErrorResponse(w, http.StatusUnauthorized, "Invalid User ID", "USER_ERROR")
		return
	}
	c.UserId = id

	err = DB.CreateCategory(&c)
	if err != nil {
		RSP.SendErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("Error creating category: %v", err.Error()), "DATABASE_ERROR")
		return
	}

	RSP.SendSuccessResponse(w, http.StatusCreated, c.String())
}
