package toast

import (
  JSON "breakfast/_internal/json"
	BFE "breakfast/_internal/errors"
	"breakfast/models"
	DB "breakfast/repositories/toast"
	RSP "breakfast/_internal/response"
	"github.com/google/uuid"
	"net/http"
)

var config = models.ValidationConfig{
  IgnoreFields: map[string]bool{
    "Description": true,  // Optional field
  },
  ForbiddenFields: map[string]bool{
    "user_id": true,     // Set by server
    "session_id": true,  // Set by server
    "duration": true,    // Calculated field
    "end_time": true,    // Set by stopSession
  },
}

func startSession(w http.ResponseWriter, r *http.Request) {
	var t models.Toast

  fields, err := JSON.NewBFDecoder(r.Body).Model(&t)
  if BFE.HandleError(w, err) {
    return
  }

	err = models.ValidateModel(t, fields, config)
	if BFE.HandleError(w, err) {
		return
	}

	claims, err := models.GetUserClaims(r)
	if BFE.HandleError(w, err) {
		return
	}

	t.UserID, _ = uuid.Parse(claims.UserID)
  ts, err := DB.StartToastSession(&t)
	if BFE.HandleError(w, err) {
		return
	}

	RSP.SendObjectResponse(w, http.StatusCreated, ts)
}
