package users

import (
	"breakfast/models"
	DB "breakfast/repositories"
	RSP "breakfast/response"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func registerUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		RSP.SendErrorResponse(w, http.StatusBadRequest, "Error parsing JSON", "JSON_ERROR")
		return
	}

  excludeFields := map[string]bool{ "UserID": true, }
  err = models.IsModelValid(user, excludeFields)
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
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			RSP.SendErrorResponse(w, http.StatusConflict, "User with this email already exists", "USER_EXISTS")
			return
		}
		RSP.SendErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("Error creating user: %v", err.Error()), "DATABASE_ERROR")
		return
	}

	jwtToken, err := generateJWTToken(user)
	if err != nil {
		RSP.SendErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("Error authenticating user: %v", err.Error()), "AUTH_ERROR")
		return
	}

	RSP.SendSuccessResponse(w, http.StatusCreated, jwtToken)
}
