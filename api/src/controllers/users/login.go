package users

import (
	"breakfast/models"
	DB "breakfast/repositories"
	RSP "breakfast/response"
	"database/sql"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func loginUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		RSP.SendErrorResponse(w, http.StatusBadRequest, "Error parsing JSON", "JSON_ERROR")
		return
	}

	if user.Email == "" {
		RSP.SendErrorResponse(w, http.StatusUnprocessableEntity, "Email field empty", "MISSING_EMAIL")
		return
	}

	if user.Password == "" {
		RSP.SendErrorResponse(w, http.StatusUnprocessableEntity, "Password field empty", "MISSING_PASSWORD")
		return
	}

	db_user, err := DB.GetUserByEmail(user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			RSP.SendErrorResponse(w, http.StatusNotFound, "User not found", "USER_NOT_FOUND")
			return
		}
		RSP.SendErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("Database error: %v", err.Error()), "DATABASE_ERROR")
		return
	}

	err = models.CheckUserPassword(db_user.Password, user.Password)
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			RSP.SendErrorResponse(w, http.StatusUnauthorized, "Wrong password", "PASSWORD_ERROR")
			return
		}
		RSP.SendErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("Database error: %v", err.Error()), "DATABASE_ERROR")
		return
	}

	jwtToken, err := generateJWTToken(*db_user)
	if err != nil {
		RSP.SendErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("Error creating auth token: %v", err.Error()), "AUTH_ERROR")
		return
	}

	RSP.SendSuccessResponse(w, http.StatusOK, jwtToken)
}
