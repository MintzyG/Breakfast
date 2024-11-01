package yogurt

import (
	BFE "breakfast/_internal/errors"
	"breakfast/models"
	DB "breakfast/repositories/yogurt"
	RSP "breakfast/_internal/response"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

func deleteTask(w http.ResponseWriter, r *http.Request) {
	category_idStr := r.PathValue("id")
	category_id, err := strconv.Atoi(category_idStr)
	if BFE.HandleError(w, err) {
		return
	}

	claims, err := models.GetUserClaims(r)
	if BFE.HandleError(w, err) {
		return
	}

	user_id, _ := uuid.Parse(claims.UserID)
	err = DB.DeleteTask(category_id, user_id)
	if BFE.HandleError(w, err) {
		return
	}

	RSP.SendSuccessResponse(w, http.StatusOK, "Task delete successfully!")
}
