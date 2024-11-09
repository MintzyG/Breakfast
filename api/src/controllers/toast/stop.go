package toast

import (
	BFE "breakfast/_internal/errors"
	RSP "breakfast/_internal/response"
  "breakfast/_internal/cors"
	"breakfast/models"
	DB "breakfast/repositories/toast"
	"net/http"
	"strconv"
)

func stopSession(w http.ResponseWriter, r *http.Request) {
  cors.EnableCors(&w)
	session_idStr := r.PathValue("id")
	session_id, err := strconv.Atoi(session_idStr)
	if BFE.HandleError(w, err) {
		return
	}

	user_id, err := models.GetUserID(r)
	if BFE.HandleError(w, err) {
		return
	}

	session, err := DB.StopToastSession(session_id, user_id)
	if BFE.HandleError(w, err) {
		return
	}

	RSP.SendObjectResponse(w, http.StatusCreated, session)
}
