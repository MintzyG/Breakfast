package maple

import (
	"breakfast/_internal/cors"
	BFE "breakfast/_internal/errors"
	RSP "breakfast/_internal/response"
	"breakfast/models"
	DB "breakfast/repositories/maple"
	"net/http"
)

var configCreate = models.ValidationConfig{
	ForbiddenFields: map[string]bool{
		"habit_id":   true, // Set by server
		"user_id":    true, // Set by server
		"curr_streak": true, // Set by server
		"highest_streak": true, // Set by server
		"days_performed": true, // Set by MarkDay
	},
}

func createHabit(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	var m models.Maple
	_, err := models.FillModelFromJSON(r, &m, configCreate)
	if BFE.HandleError(w, err) {
		return
	}

	err = DB.CreateHabit(&m)
	if BFE.HandleError(w, err) {
		return
	}

	RSP.SendObjectResponse(w, http.StatusCreated, m)
}
