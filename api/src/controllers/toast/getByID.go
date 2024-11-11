package toast

import (
	"breakfast/_internal/cors"
	BFE "breakfast/_internal/errors"
	RSP "breakfast/_internal/response"
	"breakfast/models"
	DB "breakfast/repositories/toast"
	"net/http"
	"strconv"
)

func getSessionByID(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	session_idStr := r.PathValue("id")
	session_id, err := strconv.Atoi(session_idStr)
	if BFE.HandleError(w, err) {
		return
	}

	id, err := models.GetUserID(r)
	if BFE.HandleError(w, err) {
		return
	}

	session, err := DB.GetSessionByID(session_id, id)
	if BFE.HandleError(w, err) {
		return
	}

	RSP.SendObjectResponse(w, http.StatusOK, session)
}
