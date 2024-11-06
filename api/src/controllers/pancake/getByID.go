package pancake

import (
	BFE "breakfast/_internal/errors"
	RSP "breakfast/_internal/response"
	"breakfast/models"
	DB "breakfast/repositories/pancake"
	"net/http"
	"strconv"
)

func getNoteByID(w http.ResponseWriter, r *http.Request) {
	note_idStr := r.PathValue("id")
	note_id, err := strconv.Atoi(note_idStr)
	if BFE.HandleError(w, err) {
		return
	}

	user_id, err := models.GetUserID(r)
	if BFE.HandleError(w, err) {
		return
	}

	note, err := DB.GetNoteByID(note_id, user_id)
	if BFE.HandleError(w, err) {
		return
	}

	RSP.SendObjectResponse(w, http.StatusOK, note)
}
