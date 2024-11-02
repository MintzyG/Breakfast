package yogurt

import (
	BFE "breakfast/_internal/errors"
	RSP "breakfast/_internal/response"
	"breakfast/models"
	DB "breakfast/repositories/yogurt"
	"net/http"
	"strconv"
)

func deleteTask(w http.ResponseWriter, r *http.Request) {
	category_idStr := r.PathValue("id")
	category_id, err := strconv.Atoi(category_idStr)
	if BFE.HandleError(w, err) {
		return
	}

  user_id, err := models.GetUserID(r)
  if BFE.HandleError(w, err) {
    return
  }

	err = DB.DeleteTask(category_id, user_id)
	if BFE.HandleError(w, err) {
		return
	}

	RSP.SendSuccessResponse(w, http.StatusOK, "Task delete successfully!")
}
