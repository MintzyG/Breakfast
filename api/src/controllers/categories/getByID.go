package categories

import (
	BFE "breakfast/_internal/errors"
	RSP "breakfast/_internal/response"
	"breakfast/models"
	DB "breakfast/repositories/category"
	"net/http"
	"strconv"
)

func getCategoryByID(w http.ResponseWriter, r *http.Request) {
	category_idStr := r.PathValue("id")
	category_id, err := strconv.Atoi(category_idStr)
	if BFE.HandleError(w, err) {
		return
	}

  user_id, err := models.GetUserID(r)
  if BFE.HandleError(w, err) {
    return
  }

	category, err := DB.GetCategoryByID(category_id, user_id)
	if BFE.HandleError(w, err) {
		return
	}

	RSP.SendObjectResponse(w, http.StatusOK, category)
}
