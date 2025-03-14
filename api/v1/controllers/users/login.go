package users

import (
	"breakfast/_internal/cors"
	BFE "breakfast/_internal/errors"
	RSP "breakfast/_internal/response"
	"breakfast/models"
	DB "breakfast/repositories/user"
	"encoding/json"
	"net/http"
)

var uncheckedFields = map[string]bool{"UserID": true, "FirstName": true, "LastName": true}

func loginUser(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
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
