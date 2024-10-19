package categories

import (
	"breakfast/models"
	DB "breakfast/repositories"
	RSP "breakfast/response"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func getAllCategories(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value("claims").(*models.UserClaims)
	if !ok {
		RSP.SendErrorResponse(w, http.StatusUnauthorized, "Claims missing", "CLAIMS_ERROR")
		return
	}

	user_id, err := uuid.Parse(claims.UserID)
	if err != nil {
		RSP.SendErrorResponse(w, http.StatusUnauthorized, "Invalid User ID", "USER_ERROR")
		return
	}

	categories, err := DB.GetAllCategories(user_id)
	if err != nil {
		RSP.SendErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("ERROR: %v", err.Error()), "DATABASE_ERROR")
		return
	}

	RSP.SendObjectResponse(w, http.StatusOK, categories)
}
