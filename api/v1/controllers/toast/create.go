package toast

import (
	"breakfast/_internal/cors"
	BFE "breakfast/_internal/errors"
	RSP "breakfast/_internal/response"
	"breakfast/models"
	DB "breakfast/repositories/toast"
	"errors"
	"net/http"
)

var configCreate = models.ValidationConfig{
	IgnoreFields: map[string]bool{
		"description": true, // Optional field
	},
	ForbiddenFields: map[string]bool{
		"user_id":    true, // Set by server
		"session_id": true, // Set by server
		"duration":   true, // Set by server
		"active":     true, // Server Handled
	},
}

func createSession(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	var session models.Toast
	_, err := models.FillModelFromJSON(r, &session, configCreate)
	if BFE.HandleError(w, err) {
		return
	}

	if session.EndTime.Before(session.StartTime) {
		BFE.HandleError(w, BFE.New(BFE.ErrUnprocessable, errors.New("EndTime can't be before StartTime")))
		return
	}

	session.Duration = int64(session.EndTime.Sub(session.StartTime).Seconds())

	err = DB.CreateToastSession(&session)
	if BFE.HandleError(w, err) {
		return
	}

	RSP.SendObjectResponse(w, http.StatusCreated, session)
}
