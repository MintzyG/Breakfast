package toast

import (
	BFE "breakfast/_internal/errors"
	RSP "breakfast/_internal/response"
  "breakfast/_internal/cors"
	"breakfast/models"
	DB "breakfast/repositories/toast"
	"errors"
	"net/http"
	"strconv"
)

var configPatch = models.ValidationConfig{
	IgnoreFields: map[string]bool{
		"session_name": true, // Optional
		"description":  true, // Optional
		"start_time":   true, // Optional
		"end_time":     true, // Optional
		"category_id":  true, // Optional
	},
	ForbiddenFields: map[string]bool{
		"user_id":    true, // Set by server
		"session_id": true, // Set by server
		"duration":   true, // Set by server
		"active":     true, // Server Handled
	},
}

func patchSession(w http.ResponseWriter, r *http.Request) {
  cors.EnableCors(&w)
	session_idStr := r.PathValue("id")
	session_id, err := strconv.Atoi(session_idStr)
	if BFE.HandleError(w, err) {
		return
	}

	var session models.Toast
	fields, err := models.FillModelFromJSON(r, &session, configPatch)
	if BFE.HandleError(w, err) {
		return
	}

	session.SessionID = session_id
	if fields["start_time"] || fields["end_time"] {
		if session.EndTime.Before(session.StartTime) {
			BFE.HandleError(w, BFE.New(BFE.ErrUnprocessable, errors.New("EndTime can't be before StartTime")))
			return
		}

		session.Duration = int64(session.EndTime.Sub(session.StartTime).Seconds())
		fields["duration"] = true
	}

	err = DB.PatchSession(session, fields)
	if BFE.HandleError(w, err) {
		return
	}

	RSP.SendSuccessResponse(w, http.StatusOK, "Successfully patched session!")
}
