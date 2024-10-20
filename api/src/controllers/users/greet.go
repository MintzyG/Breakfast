package users

import (
	"breakfast/models"
	DB "breakfast/repositories"
	RSP "breakfast/response"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

func greetUserByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	claims, ok := r.Context().Value("claims").(*models.UserClaims)
	if !ok {
		RSP.SendErrorResponse(w, http.StatusUnauthorized, "Claims missing", "CLAIMS_ERROR")
		return
	}

	id, _ := uuid.Parse(idStr)
	user, err := DB.GetUserByID(id)
	if err != nil {
		if err.Error() == "user not found" {
			RSP.SendErrorResponse(w, http.StatusNotFound, "Invalid user", "USER_NOT_FOUND")
		} else {
			RSP.SendErrorResponse(w, http.StatusInternalServerError, "Server Error", "SERVER_ERROR")
		}
		return
	}

	greeting := fmt.Sprintf("Hello %v! I'm %v %v", user, claims.FirstName, claims.LastName)
	RSP.SendSuccessResponse(w, http.StatusOK, greeting)
}
