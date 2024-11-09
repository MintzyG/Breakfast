package pancake

import (
	BFE "breakfast/_internal/errors"
	RSP "breakfast/_internal/response"
  "breakfast/_internal/cors"
	"breakfast/models"
	DB "breakfast/repositories/pancake"
	"net/http"
)

var configCreate = models.ValidationConfig{
	IgnoreFields: map[string]bool{
		"content": true, // Optional
	},
	ForbiddenFields: map[string]bool{
		"user_id":    true, // Set by server
		"note_id":    true, // Set by server
		"created_at": true, // Set by server
		"updated_at": true, // Set by server
	},
}

func createNote(w http.ResponseWriter, r *http.Request) {
  cors.EnableCors(&w)
	var p models.Pancake
	_, err := models.FillModelFromJSON(r, &p, configCreate)
	if BFE.HandleError(w, err) {
		return
	}

	err = DB.CreateNote(&p)
	if BFE.HandleError(w, err) {
		return
	}

	RSP.SendObjectResponse(w, http.StatusCreated, p)
}
