package toast

import (
	BFE "breakfast/_internal/errors"
	RSP "breakfast/_internal/response"
	"breakfast/models"
	DB "breakfast/repositories/toast"
	"net/http"
	"strconv"
)

func deleteSession(w http.ResponseWriter, r *http.Request) {
	session_idStr := r.PathValue("id")
	session_id, err := strconv.Atoi(session_idStr)
	if BFE.HandleError(w, err) {
		return
	}

	user_id, err := models.GetUserID(r)
	if BFE.HandleError(w, err) {
		return
	}

	err = DB.DeleteToastSession(session_id, user_id)
	if BFE.HandleError(w, err) {
		return
	}

	RSP.SendSuccessResponse(w, http.StatusOK, "Session delete successfully!")
}
