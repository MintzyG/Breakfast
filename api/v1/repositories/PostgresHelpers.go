package repositories

import (
	"github.com/lib/pq"
)

func IsForeignKeyViolation(err error) bool {
	if err, ok := err.(*pq.Error); ok {
		return err.Code == "23503"
	}
	return false
}

func IsUniqueViolation(err error) bool {
	if err, ok := err.(*pq.Error); ok {
		return err.Code == "23505"
	}
	return false
}
