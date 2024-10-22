package yogurt

import (
	BFE "breakfast/errors"
	"breakfast/models"
	DB "breakfast/repositories"
	RSP "breakfast/response"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

var excludeFields = map[string]bool{"UserID": true, "TaskID": true, "Description": true, "Completed": true}

func createTask(w http.ResponseWriter, r *http.Request) {
	var task models.YogurtTask
	err := json.NewDecoder(r.Body).Decode(&task)
	if BFE.HandleError(w, err) {
		return
	}

	err = models.IsModelValid(task, excludeFields)
	if BFE.HandleError(w, err) {
		return
	}

	claims, err := models.GetUserClaims(r)
	if BFE.HandleError(w, err) {
		return
	}

	task.UserID, _ = uuid.Parse(claims.UserID)
	err = DB.CreateYogurtTask(&task)
	if BFE.HandleError(w, err) {
		return
	}

	RSP.SendObjectResponse(w, http.StatusCreated, task)
}
