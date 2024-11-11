package yogurt

import (
	"breakfast/_internal/cors"
	BFE "breakfast/_internal/errors"
	RSP "breakfast/_internal/response"
	"breakfast/models"
	DB "breakfast/repositories/yogurt"
	"net/http"
	"strconv"
)

func getTaskByID(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	task_idStr := r.PathValue("id")
	task_id, err := strconv.Atoi(task_idStr)
	if BFE.HandleError(w, err) {
		return
	}

	user_id, err := models.GetUserID(r)
	if BFE.HandleError(w, err) {
		return
	}

	tasks, err := DB.GetTaskByID(task_id, user_id)
	if BFE.HandleError(w, err) {
		return
	}

	RSP.SendObjectResponse(w, http.StatusOK, tasks)
}
