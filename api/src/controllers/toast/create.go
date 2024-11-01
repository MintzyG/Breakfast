package toast

import (
	BFE "breakfast/_internal/errors"
	"breakfast/models"
	DB "breakfast/repositories/toast"
	RSP "breakfast/_internal/response"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
)

var uncheckedFields = map[string]bool{"UserID": true, "SessionID": true, "Description": true, "Duration": true}

func createSession(w http.ResponseWriter, r *http.Request) {
	var t models.Toast
	err := json.NewDecoder(r.Body).Decode(&t)
	if BFE.HandleError(w, err) {
		return
	}

	err = models.IsModelValid(t, uncheckedFields)
	if BFE.HandleError(w, err) {
		return
	}

	claims, err := models.GetUserClaims(r)
	if BFE.HandleError(w, err) {
		return
	}

	t.UserID, _ = uuid.Parse(claims.UserID)
	err = DB.CreateToastSession(&t)
	if BFE.HandleError(w, err) {
		return
	}

	RSP.SendObjectResponse(w, http.StatusCreated, t)
}
