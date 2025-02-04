package maple

import (
	"breakfast/_internal/cors"
	BFE "breakfast/_internal/errors"
	RSP "breakfast/_internal/response"
	"breakfast/models"
	DB "breakfast/repositories/maple"
	"net/http"
	"strconv"
)

func getHabitByID(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	habit_idStr := r.PathValue("id")
	habit_id, err := strconv.Atoi(habit_idStr)
	if BFE.HandleError(w, err) {
		return
	}

	user_id, err := models.GetUserID(r)
	if BFE.HandleError(w, err) {
		return
	}

	habit, err := DB.GetHabitByID(habit_id, user_id)
	if BFE.HandleError(w, err) {
		return
	}

	RSP.SendObjectResponse(w, http.StatusOK, habit)
}
