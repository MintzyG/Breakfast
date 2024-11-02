package toast

import (
	BFE "breakfast/_internal/errors"
	"breakfast/models"
	DB "breakfast/repositories/toast"
	RSP "breakfast/_internal/response"
	"net/http"
	"strconv"
)

func getSessionByID(w http.ResponseWriter, r *http.Request) {
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
