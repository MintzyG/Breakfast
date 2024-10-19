package categories

import (
	"breakfast/models"
	DB "breakfast/repositories"
	RSP "breakfast/response"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

func getCategoryByID(w http.ResponseWriter, r *http.Request) {
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

	category_idStr := r.PathValue("id")
	category_id, err := strconv.Atoi(category_idStr)
	if err != nil {
		RSP.SendErrorResponse(w, http.StatusUnprocessableEntity, "Possibly malformed ID", "ID_ERROR")
		return
	}

	category, err := DB.GetCategoryByID(category_id, user_id)
	if err != nil {
		RSP.SendErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("ERROR: %v", err.Error()), "DATABASE_ERROR")
		return
	}

	RSP.SendObjectResponse(w, http.StatusOK, category)
}
