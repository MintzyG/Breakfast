package toast

import (
	BFE "breakfast/_internal/errors"
	RSP "breakfast/_internal/response"
	"breakfast/models"
	DB "breakfast/repositories/toast"
	"errors"
	"net/http"
)

var configStart = models.ValidationConfig{
	IgnoreFields: map[string]bool{
		"description": true, // Optional field
	},
	ForbiddenFields: map[string]bool{
		"user_id":    true, // Set by server
		"session_id": true, // Set by server
		"duration":   true, // Calculated on stopSession
		"end_time":   true, // Set by stopSession
		"active":     true, // Server Handled
	},
}

func startSession(w http.ResponseWriter, r *http.Request) {
	var session models.Toast
	_, err := models.FillModelFromJSON(r, &session, configStart)
	if BFE.HandleError(w, err) {
		return
	}

	sessions, err := DB.GetAllSessions(session.UserID)
	if BFE.HandleError(w, err) {
		return
	}

	for _, se := range sessions {
		if se.Active {
			BFE.HandleError(w, BFE.New(BFE.ErrUnprocessable, errors.New("Already has an active session")))
			return
		}
	}

	session.Active = true
	err = DB.StartToastSession(&session)
	if BFE.HandleError(w, err) {
		return
	}

	RSP.SendObjectResponse(w, http.StatusCreated, session)
}
