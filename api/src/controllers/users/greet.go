package users

import (
	BFE "breakfast/errors"
	"breakfast/models"
	DB "breakfast/repositories"
	RSP "breakfast/response"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

func greetUserByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	claims, err := models.GetUserClaims(r)
	if BFE.HandleError(w, err) {
		return
	}

	id, _ := uuid.Parse(idStr)
	user, err := DB.GetUserByID(id)
	if BFE.HandleError(w, err) {
		return
	}

	greeting := fmt.Sprintf("Hello %v! I'm %v %v", user, claims.FirstName, claims.LastName)
	RSP.SendSuccessResponse(w, http.StatusOK, greeting)
}
