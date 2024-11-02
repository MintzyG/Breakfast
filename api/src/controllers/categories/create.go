package categories

import (
	BFE "breakfast/_internal/errors"
	RSP "breakfast/_internal/response"
	"breakfast/models"
	DB "breakfast/repositories/category"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

var uncheckedFields = map[string]bool{"UserID": true, "ID": true, "Description": true}

func createCategory(w http.ResponseWriter, r *http.Request) {
	var c models.Category
	err := json.NewDecoder(r.Body).Decode(&c)
	if BFE.HandleError(w, err) {
		return
	}

	err = models.IsModelValid(c, uncheckedFields)
	if BFE.HandleError(w, err) {
		return
	}

	claims, err := models.GetUserClaims(r)
	if BFE.HandleError(w, err) {
		return
	}

	c.UserId, _ = uuid.Parse(claims.UserID)
	err = DB.CreateCategory(&c)
	if BFE.HandleError(w, err) {
		return
	}

	RSP.SendObjectResponse(w, http.StatusCreated, c)
}
