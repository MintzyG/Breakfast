package yogurt

import (
	"breakfast/_internal/cors"
	BFE "breakfast/_internal/errors"
	RSP "breakfast/_internal/response"
	"breakfast/models"
	DB "breakfast/repositories/yogurt"
	"net/http"
)

func getAllTasks(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	user_id, err := models.GetUserID(r)
	if BFE.HandleError(w, err) {
		return
	}

	tasks, err := DB.GetAllTasks(user_id)
	if BFE.HandleError(w, err) {
		return
	}

	RSP.SendObjectResponse(w, http.StatusOK, tasks)
}
