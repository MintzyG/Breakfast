package categories

import (
	BFE "breakfast/errors"
	"breakfast/models"
	DB "breakfast/repositories"
	RSP "breakfast/response"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
)

var excludeFields = map[string]bool{"UserID": true, "ID": true, "Description": true}

func createCategory(w http.ResponseWriter, r *http.Request) {
	var c models.Category
	err := json.NewDecoder(r.Body).Decode(&c)
	if BFE.HandleError(w, err) {
		return
	}

	err = models.IsModelValid(c, excludeFields)
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
