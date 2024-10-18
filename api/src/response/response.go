package response

import (
	"log"
	"net/http"
	"encoding/json"
)

type ErrorResponse struct {
    Status  int    `json:"status"`
    Message string `json:"message"`
    Code    string `json:"code,omitempty"`
}

type SuccessResponse struct {
    Status  int    `json:"status"`
    Message string `json:"message"`
}

func SendErrorResponse(w http.ResponseWriter, status int, message string, code string) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)

    err := json.NewEncoder(w).Encode(ErrorResponse{
        Status:  status,
        Message: message,
        Code:    code,
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
        Status:  status,
        Message: message,
    })

    if err != nil {
        log.Printf("Failed to encode JSON success response: %v", err)
        http.Error(w, "An unexpected error occurred", http.StatusInternalServerError)
    }
}
