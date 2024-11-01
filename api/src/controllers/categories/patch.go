package categories

import (
	BFE "breakfast/_internal/errors"
	"breakfast/models"
	DB "breakfast/repositories/category"
	RSP "breakfast/_internal/response"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

func patchCategory(w http.ResponseWriter, r *http.Request) {
	category_idStr := r.PathValue("id")
	category_id, err := strconv.Atoi(category_idStr)
	if BFE.HandleError(w, err) {
		return
	}

	claims, err := models.GetUserClaims(r)
	if BFE.HandleError(w, err) {
		return
	}

	var updates map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&updates)
	if BFE.HandleError(w, BFE.New(BFE.ErrJSON, err)) {
		return
	}

	user_id, _ := uuid.Parse(claims.UserID)
	err = DB.PatchCategory(category_id, user_id, updates)
	if BFE.HandleError(w, err) {
		return
	}

	RSP.SendSuccessResponse(w, http.StatusOK, "Successfully patched category!")
}
