package users

import (
  BFE "breakfast/errors"
	"breakfast/models"
	DB "breakfast/repositories"
	RSP "breakfast/response"
	"encoding/json"
	"net/http"
)

var excludeFieldsLogin = map[string]bool{"UserID": true, "FirstName": true, "LastName": true}

func loginUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
  if BFE.HandleError(w, err) { return }

	err = models.IsModelValid(user, excludeFieldsLogin)
  if  BFE.HandleError(w, err) { return }

	db_user, err := DB.GetUserByEmail(user.Email)
  if BFE.HandleError(w, err) { return }

	err = models.CheckUserPassword(db_user.Password, user.Password)
  if BFE.HandleError(w, err) { return }

	jwtToken, err := generateJWTToken(*db_user)
  if BFE.HandleError(w, err) { return }

	RSP.SendSuccessResponse(w, http.StatusOK, jwtToken)
}
