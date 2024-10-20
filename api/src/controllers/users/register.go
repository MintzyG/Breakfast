package users

import (
	BFE "breakfast/errors"
	"breakfast/models"
	DB "breakfast/repositories"
	RSP "breakfast/response"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

var excludeFieldsRegister = map[string]bool{"UserID": true}

func registerUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		RSP.SendErrorResponse(w, http.StatusBadRequest, "Error parsing JSON", "JSON_ERROR")
		return
	}

	err = models.IsModelValid(user, excludeFieldsRegister)
	if err != nil {
		RSP.SendErrorResponse(w, http.StatusUnprocessableEntity, err.Error(), "MISSING_FIELDS")
		return
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		RSP.SendErrorResponse(w, http.StatusBadRequest, err.Error(), "PASSWORD_ERROR")
		return
	}
	user.Password = string(bytes)

	user.UserID = uuid.New()
	err = DB.CreateUser(&user)
  if BFE.HandleError(w, err) { return }

	jwtToken, err := generateJWTToken(user)
	if err != nil {
		RSP.SendErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("Error authenticating user: %v", err.Error()), "AUTH_ERROR")
		return
	}

	RSP.SendSuccessResponse(w, http.StatusCreated, jwtToken)
}
