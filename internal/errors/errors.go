package errors

import (
	"errors"
	"net/http"
)

type ErrorResponse struct {
	StatusCode int
	Err        error
}

func NotFoundError(msg string) *ErrorResponse {
	return &ErrorResponse{
		StatusCode: http.StatusNotFound,
		Err:        errors.New(msg),
	}
}

func InternalError(msg string) *ErrorResponse {
	return &ErrorResponse{
		StatusCode: http.StatusInternalServerError,
		Err:        errors.New(msg),
	}
}

func BadRequest(msg string) *ErrorResponse {
	return &ErrorResponse{
		StatusCode: http.StatusBadRequest,
		Err:        errors.New(msg),
	}
}
