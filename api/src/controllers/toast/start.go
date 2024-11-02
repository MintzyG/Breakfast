package toast

import (
	BFE "breakfast/_internal/errors"
	RSP "breakfast/_internal/response"
	"breakfast/models"
	DB "breakfast/repositories/toast"
	"net/http"
)

var configStart = models.ValidationConfig{
	IgnoreFields: map[string]bool{
		"Description": true, // Optional field
	},
	ForbiddenFields: map[string]bool{
		"user_id":    true, // Set by server
		"session_id": true, // Set by server
		"duration":   true, // Calculated on stopSession
		"end_time":   true, // Set by stopSession
	},
}

func startSession(w http.ResponseWriter, r *http.Request) {
	var session models.Toast
	_, err := models.FillModelFromJSON(r, &session, configStart)
	if BFE.HandleError(w, err) {
		return
	}

	err = DB.StartToastSession(&session)
	if BFE.HandleError(w, err) {
		return
	}

	RSP.SendObjectResponse(w, http.StatusCreated, session)
}
