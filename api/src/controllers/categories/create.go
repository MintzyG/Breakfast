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

var excludeFields = map[string]bool{"UserID": true, "ID": true, "Description": true}

func createCategory(w http.ResponseWriter, r *http.Request) {
	var c models.Category
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		RSP.SendErrorResponse(w, http.StatusBadRequest, "Error parsing JSON", "JSON_ERROR")
		return
	}

	err = models.IsModelValid(c, excludeFields)
	if err != nil {
		RSP.SendErrorResponse(w, http.StatusUnprocessableEntity, err.Error(), "MISSING_FIELDS")
		return
	}

	claims, ok := r.Context().Value("claims").(*models.UserClaims)
	if !ok {
		RSP.SendErrorResponse(w, http.StatusUnauthorized, "Claims missing", "CLAIMS_ERROR")
		return
	}

	c.UserId, _ = uuid.Parse(claims.UserID)
	err = DB.CreateCategory(&c)
	if err != nil {
		RSP.SendErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("Error creating category: %v", err.Error()), "DATABASE_ERROR")
		return
	}

	RSP.SendObjectResponse(w, http.StatusCreated, c)
}
