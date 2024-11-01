package users

import (
	BFE "breakfast/_internal/errors"
	"breakfast/models"
	DB "breakfast/repositories/user"
	RSP "breakfast/_internal/response"
	"encoding/json"
	"net/http"
)

var uncheckedFields = map[string]bool{"UserId":  true, "FirstName": true, "LastName": true}

func loginUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if BFE.HandleError(w, err) {
		return
	}

	err = models.IsModelValid(user, uncheckedFields)
	if BFE.HandleError(w, err) {
		return
	}

	db_user, err := DB.GetUserByEmail(user.Email)
	if BFE.HandleError(w, err) {
		return
	}

	err = models.CheckUserPassword(db_user.Password, user.Password)
	if BFE.HandleError(w, err) {
		return
	}

	jwtToken, err := generateJWTToken(*db_user)
	if BFE.HandleError(w, err) {
		return
	}

	RSP.SendSuccessResponse(w, http.StatusOK, jwtToken)
}
