package yogurt

import (
	BFE "breakfast/_internal/errors"
	RSP "breakfast/_internal/response"
	"breakfast/models"
	DB "breakfast/repositories/yogurt"
	"net/http"
	"strconv"
)

var configPatch = models.ValidationConfig{
	IgnoreFields: map[string]bool{
		"emoji":       true, // Optional field
		"title":       true, // Optional field
		"description": true, // Optional field
		"completed":   true, // Optional field
		"task_size":   true, // Optional field
		"difficulty":  true, // Optional field
		"priority":    true, // Optional field
		"category_id": true, // Optional field
	},
	ForbiddenFields: map[string]bool{
		"task_id": true, // Set by server
		"user_id": true, // Set by server
	},
}

func patchTask(w http.ResponseWriter, r *http.Request) {
	task_idStr := r.PathValue("id")
	task_id, err := strconv.Atoi(task_idStr)
	if BFE.HandleError(w, err) {
		return
	}

	var task models.YogurtTask
	fields, err := models.FillModelFromJSON(r, &task, configPatch)
	if BFE.HandleError(w, err) {
		return
	}

	task.TaskID = task_id
	err = DB.PatchTask(task, fields)
	if BFE.HandleError(w, err) {
		return
	}

	RSP.SendSuccessResponse(w, http.StatusOK, "State changed successfully!")
}
