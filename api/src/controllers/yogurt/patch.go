package yogurt

import (
	BFE "breakfast/_internal/errors"
	"breakfast/models"
	DB "breakfast/repositories/yogurt"
	RSP "breakfast/_internal/response"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

func patchTask(w http.ResponseWriter, r *http.Request) {
	task_idStr := r.PathValue("id")
	task_id, err := strconv.Atoi(task_idStr)
	if BFE.HandleError(w, err) {
		return
	}

	claims, err := models.GetUserClaims(r)
	if BFE.HandleError(w, err) {
		return
	}

	var updates map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&updates)
	if BFE.HandleError(w, BFE.New(BFE.ErrJSON, err)) {
		return
	}

	user_id, _ := uuid.Parse(claims.UserID)
	err = DB.PatchTask(task_id, user_id, updates)
	if BFE.HandleError(w, err) {
		return
	}

	RSP.SendSuccessResponse(w, http.StatusOK, "State changed successfully!")
}
