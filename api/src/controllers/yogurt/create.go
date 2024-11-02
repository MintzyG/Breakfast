package yogurt

import (
	BFE "breakfast/_internal/errors"
	RSP "breakfast/_internal/response"
	"breakfast/models"
	DB "breakfast/repositories/yogurt"
	"net/http"
)

var configCreate = models.ValidationConfig{
	IgnoreFields: map[string]bool{
		"description": true, // Optional field
	},
	ForbiddenFields: map[string]bool{
		"task_id": true, // Set by server
		"user_id":    true, // Set by server
	},
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var task models.YogurtTask
  _, err := models.FillModelFromJSON(r, &task, configCreate)
  if BFE.HandleError(w, err) {
    return
  }

	err = DB.CreateYogurtTask(&task)
	if BFE.HandleError(w, err) {
		return
	}

	RSP.SendObjectResponse(w, http.StatusCreated, task)
}
