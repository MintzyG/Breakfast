package pancake

import (
	BFE "breakfast/_internal/errors"
	RSP "breakfast/_internal/response"
	"breakfast/models"
	DB "breakfast/repositories/pancake"
	"net/http"
)

func getAllNotes(w http.ResponseWriter, r *http.Request) {
	user_id, err := models.GetUserID(r)
	if BFE.HandleError(w, err) {
		return
	}

	notes, err := DB.GetAllNotes(user_id)
	if BFE.HandleError(w, err) {
		return
	}

	RSP.SendObjectResponse(w, http.StatusOK, notes)
}
