package bf_response

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

func SendErrorResponse(w http.ResponseWriter, status int, code string, context string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(ErrorResponse{
		Code:    code,
		Message: context,
	})

	if err != nil {
		log.Printf("Failed to encode JSON error response: %v", err)
		http.Error(w, "An unexpected error occurred", http.StatusInternalServerError)
	}
}

func SendSuccessResponse(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(SuccessResponse{
		Message: message,
	})

	if err != nil {
		log.Printf("Failed to encode JSON success response: %v", err)
		http.Error(w, "An unexpected error occurred", http.StatusInternalServerError)
	}
}

func SendObjectResponse[T any](w http.ResponseWriter, statusCode int, obj T) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(obj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
