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

var configPatch = models.ValidationConfig{
	IgnoreFields: map[string]bool{
		"title":         true, // Optional
		"smallest_unit": true, // Optional
		"category_id":   true, // Optional
	},
	ForbiddenFields: map[string]bool{
		"user_id":        true, // Set by server
		"habit_id":       true, // Set by server
		"curr_streak":    true, // Already set
		"highest_streak": true, // Set by server
		"days_performed": true, // Set by server
	},
}

func patchHabit(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	habit_idStr := r.PathValue("id")
	habit_id, err := strconv.Atoi(habit_idStr)
	if BFE.HandleError(w, err) {
		return
	}

	var habit models.Maple
	fields, err := models.FillModelFromJSON(r, &habit, configPatch)
	if BFE.HandleError(w, err) {
		return
	}

	habit.HabitID = habit_id
	err = DB.PatchHabit(habit, fields)
	if BFE.HandleError(w, err) {
		return
	}

	RSP.SendSuccessResponse(w, http.StatusOK, "Successfully patched habit!")
}
