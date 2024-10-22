package yogurt

import (
	BFE "breakfast/errors"
	"breakfast/models"
	DB "breakfast/repositories"
	RSP "breakfast/response"
	"net/http"

	"github.com/google/uuid"
)

func getAllTasks(w http.ResponseWriter, r *http.Request) {
  claims, err := models.GetUserClaims(r)
  if BFE.HandleError(w, err) {
    return
  }

  user_id, _ := uuid.Parse(claims.UserID)
  tasks, err := DB.GetAllTasks(user_id)
  if BFE.HandleError(w, err) {
    return
  }

  RSP.SendObjectResponse(w, http.StatusOK, tasks)
}
