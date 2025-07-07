package utils

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

type CustomError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Err     error  `json:"error"`
}

func NewCustomPGError(message string, code int, err error) *CustomError {
	var pgxErr *pgconn.PgError
	if errors.As(err, &pgxErr) {
		if pgxErr.Code == "23503" {
			return &CustomError{
				Message: message,
				Code:    code,
				Err:     err,
			}
		}
		if pgxErr.Code == "23505" {
			return &CustomError{
				Message: message,
				Code:    code,
				Err:     err,
			}
		}
	}

	return &CustomError{
		Message: err.Error(),
		Code:    400,
		Err:     err,
	}
}
