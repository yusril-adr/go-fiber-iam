package errors

import (
	"net/http"
)

type AppError struct {
	StatusCode int
	Message    string
	Trace      string
}

func (e *AppError) Error() string {
	return e.Message
}

func BadRequest(msg string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		Message:    msg,
	}
}

func NotFound(msg string) *AppError {
	return &AppError{
		StatusCode: http.StatusNotFound,
		Message:    msg,
	}
}

func Internal(msg string) *AppError {
	return &AppError{
		StatusCode: http.StatusInternalServerError,
		Message:    msg,
	}
}

func Unauthorized(msg string) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		Message:    msg,
	}
}
