package yogurt

import (
	BFE "breakfast/_internal/errors"
	RSP "breakfast/_internal/response"
	"breakfast/models"
	DB "breakfast/repositories/yogurt"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

func getTaskByID(w http.ResponseWriter, r *http.Request) {
	task_idStr := r.PathValue("id")
	task_id, err := strconv.Atoi(task_idStr)
	if BFE.HandleError(w, err) {
		return
	}

	claims, err := models.GetUserClaims(r)
	if BFE.HandleError(w, err) {
		return
	}

	user_id, _ := uuid.Parse(claims.UserID)
	tasks, err := DB.GetTaskByID(task_id, user_id)
	if BFE.HandleError(w, err) {
		return
	}

	RSP.SendObjectResponse(w, http.StatusOK, tasks)
}
