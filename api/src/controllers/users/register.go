package users

import (
	BFE "breakfast/_internal/errors"
	"breakfast/models"
	DB "breakfast/repositories/user"
	RSP "breakfast/_internal/response"
	"encoding/json"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

var uncheckedFieldsRegister = map[string]bool{"UserID": true}

func registerUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if BFE.HandleError(w, err) {
		return
	}

	err = models.IsModelValid(user, uncheckedFieldsRegister)
	if BFE.HandleError(w, err) {
		return
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if BFE.HandleError(w, err) {
		return
	}
	user.Password = string(bytes)

	user.UserID = uuid.New()
	err = DB.CreateUser(&user)
	if BFE.HandleError(w, err) {
		return
	}

	jwtToken, err := generateJWTToken(user)
	if BFE.HandleError(w, err) {
		return
	}

	RSP.SendSuccessResponse(w, http.StatusCreated, jwtToken)
}
