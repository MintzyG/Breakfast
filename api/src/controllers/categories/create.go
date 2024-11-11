package categories

import (
	"breakfast/_internal/cors"
	BFE "breakfast/_internal/errors"
	RSP "breakfast/_internal/response"
	"breakfast/models"
	DB "breakfast/repositories/category"
	"net/http"
)

var configCreate = models.ValidationConfig{
	IgnoreFields: map[string]bool{
		"description": true, // Optional
	},
	ForbiddenFields: map[string]bool{
		"user_id":     true, // Set by server
		"category_id": true, // Set by server
	},
}

func createCategory(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	var c models.Category
	_, err := models.FillModelFromJSON(r, &c, configCreate)
	if BFE.HandleError(w, err) {
		return
	}

	err = DB.CreateCategory(&c)
	if BFE.HandleError(w, err) {
		return
	}

	RSP.SendObjectResponse(w, http.StatusCreated, c)
}
