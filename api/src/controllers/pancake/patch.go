package pancake

import (
	BFE "breakfast/_internal/errors"
	RSP "breakfast/_internal/response"
  "breakfast/_internal/cors"
	"breakfast/models"
	DB "breakfast/repositories/pancake"
	"net/http"
	"strconv"
	"time"
)

var configPatch = models.ValidationConfig{
	IgnoreFields: map[string]bool{
		"title":       true, // Optional
		"content":     true, // Optional
		"category_id": true, // Optional
	},
	ForbiddenFields: map[string]bool{
		"user_id":    true, // Set by server
		"note_id":    true, // Set by server
		"created_at": true, // Already set
		"updated_at": true, // Set by server
	},
}

func patchNote(w http.ResponseWriter, r *http.Request) {
  cors.EnableCors(&w)
	note_idStr := r.PathValue("id")
	note_id, err := strconv.Atoi(note_idStr)
	if BFE.HandleError(w, err) {
		return
	}

	var note models.Pancake
	fields, err := models.FillModelFromJSON(r, &note, configPatch)
	if BFE.HandleError(w, err) {
		return
	}

	note.NoteID = note_id
	fields["updated_at"] = true
	note.UpdatedAt = time.Now()
	err = DB.PatchNote(note, fields)
	if BFE.HandleError(w, err) {
		return
	}

	RSP.SendSuccessResponse(w, http.StatusOK, "Successfully patched note!")
}
