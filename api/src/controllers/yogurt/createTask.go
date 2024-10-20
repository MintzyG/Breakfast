package yogurt

import (
	"breakfast/models"
	DB "breakfast/repositories"
	RSP "breakfast/response"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

var excludeFields = map[string]bool{"UserID": true, "TaskID": true, "Description": true}

func createTask(w http.ResponseWriter, r *http.Request) {
	var task models.YogurtTask
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		RSP.SendErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", err.Error()), "JSON_ERROR")
		return
	}

	err = models.IsModelValid(task, excludeFields)
	if err != nil {
		RSP.SendErrorResponse(w, http.StatusUnprocessableEntity, err.Error(), "MISSING_FIELDS")
		return
	}

	claims, ok := r.Context().Value("claims").(*models.UserClaims)
	if !ok {
		RSP.SendErrorResponse(w, http.StatusUnauthorized, "Claims missing", "CLAIMS_ERROR")
		return
	}

	task.UserID, _ = uuid.Parse(claims.UserID)
	err = DB.CreateYogurtTask(&task)
	if err != nil {
		RSP.SendErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("Error creating task: %v", err.Error()), "DATABASE_ERROR")
		return
	}

	RSP.SendObjectResponse(w, http.StatusCreated, task)
}
