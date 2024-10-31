package toast

import (
	BFE "breakfast/errors"
	"breakfast/models"
	DB "breakfast/repositories/toast"
	RSP "breakfast/response"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
)

var uncheckedFieldsStart = map[string]bool{"UserID": true, "SessionID": true, "Duration": true, "EndTime": true, "Description": true}

func startSession(w http.ResponseWriter, r *http.Request) {
	var t models.Toast
	err := json.NewDecoder(r.Body).Decode(&t)
	if BFE.HandleError(w, err) {
		return
	}

	err = models.IsModelValid(t, uncheckedFieldsStart)
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
