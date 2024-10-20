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
	AUTH_ERROR_CODE           = "AUTH_ERROR"
	CLAIMS_ERROR_CODE         = "CLAIMS_ERROR"
	CONFLICT_ERROR_CODE       = "CONFLICT_ERROR"
	DATABASE_ERROR_CODE       = "DATABASE_ERROR"
	EXPIRED_TOKEN_CODE        = "EXPIRED_TOKEN"
	FORBIDDEN_ACTION_CODE     = "FORBIDDEN_ACTION"
	HEADER_MALFORMED_CODE     = "HEADER_MALFORMED"
	HEADER_MISSING_CODE       = "HEADER_MISSING"
	JSON_ERROR_CODE           = "JSON_ERROR"
	MISSING_FIELDS_CODE       = "MISSING_FIELDS"
	NOT_IMPLEMENTED_CODE      = "NOT_IMPLEMENTED"
	PASSWORD_ERROR_CODE       = "PASSWORD_ERROR"
	RATE_LIMIT_EXCEEDED_CODE  = "RATE_LIMIT_EXCEEDED"
	RESOURCE_NOT_FOUND_CODE   = "RESOURCE_NOT_FOUND"
	SERVER_ERROR_CODE         = "SERVER_ERROR"
	UNAUTHORIZED_ACCESS_CODE  = "UNAUTHORIZED_ACCESS"
	UNPROCESSABLE_ENTITY_CODE = "UNPROCESSABLE_ENTITY"
	USER_NOT_FOUND_CODE       = "USER_NOT_FOUND"
	VALIDATION_ERROR_CODE     = "VALIDATION_ERROR"
)

var (
	AUTH_ERROR           = BFError{HttpCode: http.StatusUnauthorized, ErrorCode: AUTH_ERROR_CODE, ErrorContext: ""}
	CLAIMS_ERROR         = BFError{HttpCode: http.StatusUnauthorized, ErrorCode: CLAIMS_ERROR_CODE, ErrorContext: ""}
	CONFLICT_ERROR       = BFError{HttpCode: http.StatusConflict, ErrorCode: CONFLICT_ERROR_CODE, ErrorContext: ""}
	DATABASE_ERROR       = BFError{HttpCode: http.StatusInternalServerError, ErrorCode: DATABASE_ERROR_CODE, ErrorContext: ""}
	EXPIRED_TOKEN        = BFError{HttpCode: http.StatusUnauthorized, ErrorCode: EXPIRED_TOKEN_CODE, ErrorContext: ""}
	FORBIDDEN_ACTION     = BFError{HttpCode: http.StatusForbidden, ErrorCode: FORBIDDEN_ACTION_CODE, ErrorContext: ""}
	HEADER_MALFORMED     = BFError{HttpCode: http.StatusUnauthorized, ErrorCode: HEADER_MALFORMED_CODE, ErrorContext: ""}
	HEADER_MISSING       = BFError{HttpCode: http.StatusUnauthorized, ErrorCode: HEADER_MISSING_CODE, ErrorContext: ""}
	JSON_ERROR           = BFError{HttpCode: http.StatusBadRequest, ErrorCode: JSON_ERROR_CODE, ErrorContext: ""}
	MISSING_FIELDS       = BFError{HttpCode: http.StatusUnprocessableEntity, ErrorCode: MISSING_FIELDS_CODE, ErrorContext: ""}
	NOT_IMPLEMENTED      = BFError{HttpCode: http.StatusNotImplemented, ErrorCode: NOT_IMPLEMENTED_CODE, ErrorContext: ""}
	PASSWORD_ERROR       = BFError{HttpCode: http.StatusUnauthorized, ErrorCode: PASSWORD_ERROR_CODE, ErrorContext: ""}
	RATE_LIMIT_EXCEEDED  = BFError{HttpCode: http.StatusTooManyRequests, ErrorCode: RATE_LIMIT_EXCEEDED_CODE, ErrorContext: ""}
	RESOURCE_NOT_FOUND   = BFError{HttpCode: http.StatusNotFound, ErrorCode: RESOURCE_NOT_FOUND_CODE, ErrorContext: ""}
	SERVER_ERROR         = BFError{HttpCode: http.StatusInternalServerError, ErrorCode: SERVER_ERROR_CODE, ErrorContext: ""}
	UNAUTHORIZED_ACCESS  = BFError{HttpCode: http.StatusUnauthorized, ErrorCode: UNAUTHORIZED_ACCESS_CODE, ErrorContext: ""}
	UNPROCESSABLE_ENTITY = BFError{HttpCode: http.StatusUnprocessableEntity, ErrorCode: UNPROCESSABLE_ENTITY_CODE, ErrorContext: ""}
	USER_NOT_FOUND       = BFError{HttpCode: http.StatusNotFound, ErrorCode: USER_NOT_FOUND_CODE, ErrorContext: ""}
	VALIDATION_ERROR     = BFError{HttpCode: http.StatusBadRequest, ErrorCode: VALIDATION_ERROR_CODE, ErrorContext: ""}
)

var errorMap = map[string]BFError{
	AUTH_ERROR_CODE:           AUTH_ERROR,
	CLAIMS_ERROR_CODE:         CLAIMS_ERROR,
	CONFLICT_ERROR_CODE:       CONFLICT_ERROR,
	DATABASE_ERROR_CODE:       DATABASE_ERROR,
	EXPIRED_TOKEN_CODE:        EXPIRED_TOKEN,
	FORBIDDEN_ACTION_CODE:     FORBIDDEN_ACTION,
	HEADER_MALFORMED_CODE:     HEADER_MALFORMED,
	HEADER_MISSING_CODE:       HEADER_MISSING,
	JSON_ERROR_CODE:           JSON_ERROR,
	MISSING_FIELDS_CODE:       MISSING_FIELDS,
	NOT_IMPLEMENTED_CODE:      NOT_IMPLEMENTED,
	PASSWORD_ERROR_CODE:       PASSWORD_ERROR,
	RATE_LIMIT_EXCEEDED_CODE:  RATE_LIMIT_EXCEEDED,
	RESOURCE_NOT_FOUND_CODE:   RESOURCE_NOT_FOUND,
	SERVER_ERROR_CODE:         SERVER_ERROR,
	UNAUTHORIZED_ACCESS_CODE:  UNAUTHORIZED_ACCESS,
	UNPROCESSABLE_ENTITY_CODE: UNPROCESSABLE_ENTITY,
	USER_NOT_FOUND_CODE:       USER_NOT_FOUND,
	VALIDATION_ERROR_CODE:     VALIDATION_ERROR,
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
