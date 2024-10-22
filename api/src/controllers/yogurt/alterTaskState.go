package yogurt

import (
	BFE "breakfast/errors"
	"breakfast/models"
	DB "breakfast/repositories"
	RSP "breakfast/response"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

type taskState struct {
  State bool `json:"state"`
}

func alterTaskState(w http.ResponseWriter, r *http.Request) {
  task_idStr := r.PathValue("id")
	task_id, err := strconv.Atoi(task_idStr)
	if BFE.HandleError(w, err) {
		return
  }

  var state taskState
  err = json.NewDecoder(r.Body).Decode(&state)
  if BFE.HandleError(w, err) {
    return
  }

  claims, err := models.GetUserClaims(r)
  if BFE.HandleError(w, err) {
    return
  }

  user_id, _ := uuid.Parse(claims.UserID)
  err = DB.AlterTaskCompletedStatus(task_id, user_id, state.State)
  if BFE.HandleError(w, err) {
    return
  }

  RSP.SendSuccessResponse(w, http.StatusOK, "State changed successfully!")
}
