package helper

import (
	"errors"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"gorm.io/gorm"
)

func GetStatusCode(err error) int {
	if isValidationError(err) {
		return http.StatusBadRequest
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusNotFound
	}

	if errors.Is(err, gorm.ErrMissingWhereClause) {
		return http.StatusBadRequest
	}

	if isPostgresErrorUniqueViolation(err) {
		return http.StatusConflict
	}

	return http.StatusInternalServerError
}

func isValidationError(err error) bool {
	_, ok := err.(validation.Errors)
	return ok
}

func isPostgresErrorUniqueViolation(err error) bool {
	pgError, ok := err.(*pgconn.PgError)
	if ok {
		return pgError.Code == pgerrcode.UniqueViolation
	}
	return false
}
