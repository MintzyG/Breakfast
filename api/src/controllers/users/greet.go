package users

import (
	DB "breakfast/repositories"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

func greetUserByID(w http.ResponseWriter, r *http.Request) {
	// Extract the id from the path
	idStr := r.PathValue("id")

	// Convert the string id to a UUID
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid UUID", http.StatusBadRequest)
		return
	}

	// Get user from database
	user, err := DB.GetUserByID(id)
	if err != nil {
		if err.Error() == "user not found" {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, "Server error", http.StatusInternalServerError)
		}
		return
	}

	// Greet the user
	greeting := fmt.Sprintf("Hello %v!", user)
	fmt.Fprint(w, greeting)
}
