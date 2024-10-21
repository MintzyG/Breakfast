package errors

import (
	RSP "breakfast/response"
	"fmt"
	"net/http"
)

type ErrorType struct {
    HttpCode  int
    ErrorCode string
}

type BFError struct {
  HttpCode     int    `json:"http_code"`
  ErrorCode    string `json:"error_code"`
  ErrorMessage error  `json:"-"`
}

func (e BFError) Error() string {
    if e.ErrorMessage != nil {
        return fmt.Sprintf("%s: %s", e.ErrorCode, e.ErrorMessage.Error())
    }
    return e.ErrorCode
}

func (e BFError) Unwrap() error {
    return e.ErrorMessage
}

var (
    ErrAuth              = ErrorType{http.StatusUnauthorized, "ERR_AUTH"}
    ErrClaims            = ErrorType{http.StatusUnauthorized, "ERR_CLAIMS"}
    ErrConflict          = ErrorType{http.StatusConflict, "ERR_CONFLICT"}
    ErrDatabase          = ErrorType{http.StatusInternalServerError, "ERR_DATABASE"}
    ErrTokenExpired      = ErrorType{http.StatusUnauthorized, "ERR_TOKEN_EXPIRED"}
    ErrForbidden         = ErrorType{http.StatusForbidden, "ERR_FORBIDDEN"}
    ErrHeaderMalformed   = ErrorType{http.StatusUnauthorized, "ERR_HEADER_MALFORMED"}
    ErrHeaderMissing     = ErrorType{http.StatusUnauthorized, "ERR_HEADER_MISSING"}
    ErrJSON              = ErrorType{http.StatusBadRequest, "ERR_JSON"}
    ErrMissingFields     = ErrorType{http.StatusUnprocessableEntity, "ERR_MISSING_FIELDS"}
    ErrNotImplemented    = ErrorType{http.StatusNotImplemented, "ERR_NOT_IMPLEMENTED"}
    ErrPassword          = ErrorType{http.StatusUnauthorized, "ERR_PASSWORD"}
    ErrRateLimitExceeded = ErrorType{http.StatusTooManyRequests, "ERR_RATE_LIMIT_EXCEEDED"}
    ErrResourceNotFound  = ErrorType{http.StatusNotFound, "ERR_RESOURCE_NOT_FOUND"}
    ErrServer            = ErrorType{http.StatusInternalServerError, "ERR_SERVER"}
    ErrUnauthorized      = ErrorType{http.StatusUnauthorized, "ERR_UNAUTHORIZED"}
    ErrUnprocessable     = ErrorType{http.StatusUnprocessableEntity, "ERR_UNPROCESSABLE"}
    ErrUserNotFound      = ErrorType{http.StatusNotFound, "ERR_USER_NOT_FOUND"}
    ErrValidation        = ErrorType{http.StatusBadRequest, "ERR_VALIDATION"}
)

func NewBFError(errorType ErrorType, err error) BFError {
    return BFError{
        HttpCode:     errorType.HttpCode,
        ErrorCode:    errorType.ErrorCode,
        ErrorMessage: err,
    }
}

func HandleError(w http.ResponseWriter, err error) bool {
    if err == nil {
        return false
    }

    if bfErr, ok := err.(BFError); ok {
        var errorMessage string
        if bfErr.ErrorMessage == nil {
           return false
        }
        errorMessage = bfErr.ErrorMessage.Error()
        RSP.SendErrorResponse(w, bfErr.HttpCode, bfErr.ErrorCode, errorMessage)
    } else {
      if bfErr.ErrorMessage == nil {
         return false
      }
      RSP.SendErrorResponse(w, http.StatusInternalServerError, "UNKNOWN_ERROR", err.Error())
    }
    return true
}

