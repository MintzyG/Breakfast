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

var configPatch = models.ValidationConfig{
	IgnoreFields: map[string]bool{
		"title":       true, // Optional
		"description": true, // Optional
		"emoji":       true, // Optional
		"color":       true, // Optional
		"text_color":  true, // Optional
	},
	ForbiddenFields: map[string]bool{
		"user_id":     true, // Set by server
		"category_id": true, // Set by server
	},
}

func patchCategory(w http.ResponseWriter, r *http.Request) {
  cors.EnableCors(&w)
	category_idStr := r.PathValue("id")
	category_id, err := strconv.Atoi(category_idStr)
	if BFE.HandleError(w, err) {
		return
	}

	var category models.Category
	fields, err := models.FillModelFromJSON(r, &category, configPatch)
	if BFE.HandleError(w, err) {
		return
	}

	category.ID = category_id
	err = DB.PatchCategory(category, fields)
	if BFE.HandleError(w, err) {
		return
	}

	RSP.SendSuccessResponse(w, http.StatusOK, "Successfully patched category!")
}
