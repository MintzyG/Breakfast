package toast

import (
	"breakfast/_internal/cors"
	BFE "breakfast/_internal/errors"
	RSP "breakfast/_internal/response"
	"breakfast/models"
	DB "breakfast/repositories/toast"
	"net/http"
)

func getAllSessions(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	user_id, err := models.GetUserID(r)
	if BFE.HandleError(w, err) {
		return
	}

	sessions, err := DB.GetAllSessions(user_id)
	if BFE.HandleError(w, err) {
		return
	}

	RSP.SendObjectResponse(w, http.StatusOK, sessions)
}
