package categories

import (
	BFE "breakfast/_internal/errors"
	RSP "breakfast/_internal/response"
	"breakfast/models"
	DB "breakfast/repositories/category"
	"net/http"
)

func getAllCategories(w http.ResponseWriter, r *http.Request) {
	user_id, err := models.GetUserID(r)
	if BFE.HandleError(w, err) {
		return
	}

	categories, err := DB.GetAllCategories(user_id)
	if BFE.HandleError(w, err) {
		return
	}

	RSP.SendObjectResponse(w, http.StatusOK, categories)
}
