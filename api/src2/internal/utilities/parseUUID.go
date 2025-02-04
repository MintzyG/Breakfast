package utilities

import (
	"net/http"

	"github.com/google/uuid"
)

func ParseUUID(w http.ResponseWriter, uid string) (uuid.UUID, error) {
	id, err := uuid.Parse(uid)
	if err != nil {
    Send(w, "UUID coudln't be parsed", nil, http.StatusUnauthorized)
		return uuid.UUID{}, err
  }
  return id, nil
}
