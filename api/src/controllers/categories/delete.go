package categories

import (
	BFE "breakfast/_internal/errors"
	RSP "breakfast/_internal/response"
  "breakfast/_internal/cors"
	"breakfast/models"
	DB "breakfast/repositories/category"
	"net/http"
	"strconv"
)

func deleteCategory(w http.ResponseWriter, r *http.Request) {
  cors.EnableCors(&w)
	category_idStr := r.PathValue("id")
	category_id, err := strconv.Atoi(category_idStr)
	if BFE.HandleError(w, err) {
		return
	}

	user_id, err := models.GetUserID(r)
	if BFE.HandleError(w, err) {
		return
	}

	err = DB.DeleteCategory(category_id, user_id)
	if BFE.HandleError(w, err) {
		return
	}

	RSP.SendSuccessResponse(w, http.StatusOK, "Category delete successfully!")
}
