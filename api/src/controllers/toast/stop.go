package toast

import (
	BFE "breakfast/_internal/errors"
	RSP "breakfast/_internal/response"
	"breakfast/models"
	DB "breakfast/repositories/toast"
	"fmt"
	"net/http"
)

var configStop = models.ValidationConfig{
	IgnoreFields: map[string]bool{
		"Description": true, // Optional field
	},
	ForbiddenFields: map[string]bool{
		"user_id":     true, // Already set
		"duration":    true, // Calculated on stopSession
		"start_time":  true, // Already set
		"title":       true, // Already set
		"category_id": true, // Already set
	},
}

func stopSession(w http.ResponseWriter, r *http.Request) {
	var session models.Toast
	_, err := models.FillModelFromJSON(r, &session, configStop)
	if BFE.HandleError(w, err) {
		return
	}

	fmt.Println(session.SessionID)
	err = DB.StopToastSession(&session)
	if BFE.HandleError(w, err) {
		return
	}

	RSP.SendObjectResponse(w, http.StatusCreated, session)
}
