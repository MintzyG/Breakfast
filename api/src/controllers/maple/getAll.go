package maple

import (
	"breakfast/_internal/cors"
	BFE "breakfast/_internal/errors"
	RSP "breakfast/_internal/response"
	"breakfast/models"
	DB "breakfast/repositories/maple"
	"net/http"
)

func getAllHabits(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	user_id, err := models.GetUserID(r)
	if BFE.HandleError(w, err) {
		return
	}

	habits, err := DB.GetAllHabits(user_id)
	if BFE.HandleError(w, err) {
		return
	}

	RSP.SendObjectResponse(w, http.StatusOK, habits)
}
