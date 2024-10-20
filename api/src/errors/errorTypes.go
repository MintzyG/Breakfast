package errors

import (
	RSP "breakfast/response"
	"net/http"
)

type BFError struct {
	HttpCode     int    `json:"http_code"`
	ErrorCode    string `json:"error_code"`
	ErrorContext string `json:"error_context"`
}

func (e BFError) Error() string {
	return e.ErrorContext
}

const (
	CLAIMS_ERROR_CODE        = "CLAIMS_ERROR"
	DATABASE_ERROR_CODE      = "DATABASE_ERROR"
	HEADER_MALFORMED_CODE    = "HEADER_MALFORMED"
	HEADER_MISSING_CODE      = "HEADER_MISSING"
	JSON_ERROR_CODE          = "JSON_ERROR"
	MISSING_FIELDS_CODE      = "MISSING_FIELDS"
	SERVER_ERROR_CODE        = "SERVER_ERROR"
	UNAUTHORIZED_ACCESS_CODE = "UNAUTHORIZED_ACCESS"
	USER_NOT_FOUND_CODE      = "USER_NOT_FOUND"
	PASSWORD_ERROR_CODE      = "PASSWORD_ERROR"
  AUTH_ERROR_CODE          = "AUTH_ERROR"
)

var (
	CLAIMS_ERROR        = BFError{HttpCode: http.StatusUnauthorized, ErrorCode: CLAIMS_ERROR_CODE, ErrorContext: ""}
	DATABASE_ERROR      = BFError{HttpCode: http.StatusInternalServerError, ErrorCode: DATABASE_ERROR_CODE, ErrorContext: ""}
	HEADER_MALFORMED    = BFError{HttpCode: http.StatusUnauthorized, ErrorCode: HEADER_MALFORMED_CODE, ErrorContext: ""}
	HEADER_MISSING      = BFError{HttpCode: http.StatusUnauthorized, ErrorCode: HEADER_MISSING_CODE, ErrorContext: ""}
	JSON_ERROR          = BFError{HttpCode: http.StatusBadRequest, ErrorCode: JSON_ERROR_CODE, ErrorContext: ""}
	MISSING_FIELDS      = BFError{HttpCode: http.StatusUnprocessableEntity, ErrorCode: MISSING_FIELDS_CODE, ErrorContext: ""}
	SERVER_ERROR        = BFError{HttpCode: http.StatusInternalServerError, ErrorCode: SERVER_ERROR_CODE, ErrorContext: ""}
	UNAUTHORIZED_ACCESS = BFError{HttpCode: http.StatusUnauthorized, ErrorCode: UNAUTHORIZED_ACCESS_CODE, ErrorContext: ""}
	USER_NOT_FOUND      = BFError{HttpCode: http.StatusNotFound, ErrorCode: USER_NOT_FOUND_CODE, ErrorContext: ""}
	PASSWORD_ERROR      = BFError{HttpCode: http.StatusUnauthorized, ErrorCode: PASSWORD_ERROR_CODE, ErrorContext: ""}
  AUTH_ERROR          = BFError{HttpCode: http.StatusUnauthorized, ErrorCode: AUTH_ERROR_CODE, ErrorContext: ""}
)

var errorMap = map[string]BFError{
	CLAIMS_ERROR_CODE:        CLAIMS_ERROR,
	DATABASE_ERROR_CODE:      DATABASE_ERROR,
	HEADER_MALFORMED_CODE:    HEADER_MALFORMED,
	HEADER_MISSING_CODE:      HEADER_MISSING,
	JSON_ERROR_CODE:          JSON_ERROR,
	MISSING_FIELDS_CODE:      MISSING_FIELDS,
	SERVER_ERROR_CODE:        SERVER_ERROR,
	UNAUTHORIZED_ACCESS_CODE: UNAUTHORIZED_ACCESS,
	USER_NOT_FOUND_CODE:      USER_NOT_FOUND,
	PASSWORD_ERROR_CODE:      PASSWORD_ERROR,
  AUTH_ERROR_CODE:          AUTH_ERROR,
}

func NewBFError(errorCode, errorContext string) BFError {
	baseError, exists := errorMap[errorCode]

	if !exists {
		baseError = BFError{
			HttpCode:     http.StatusInternalServerError,
			ErrorCode:    "UNKNOWN_ERROR",
			ErrorContext: "An unknown error occurred",
		}
	}

	baseError.ErrorContext = errorContext
	return baseError
}

func HandleError(w http.ResponseWriter, err error) bool {
	if err == nil {
		return false
	}

	if bfErr, ok := err.(BFError); ok {
		RSP.SendErrorResponse(w, bfErr.HttpCode, bfErr.ErrorCode, bfErr.ErrorContext)
	} else {
		RSP.SendErrorResponse(w, http.StatusInternalServerError, "UNKNOWN_ERROR", err.Error())
	}
	return true
}
