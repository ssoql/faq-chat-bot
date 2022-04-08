package api_errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type ApiError interface {
	Message() string
	Status() int
	Error() string
	Causes() []interface{}
}

type apiError struct {
	ErrMessage string        `json:"message"`
	ErrStatus  int           `json:"status"`
	ErrError   string        `json:"error"`
	ErrCauses  []interface{} `json:"causes"`
}

func (e apiError) Message() string {
	return e.ErrMessage
}

func (e apiError) Status() int {
	return e.ErrStatus
}

func (e apiError) Causes() []interface{} {
	return e.ErrCauses
}

func (e apiError) Error() string {
	return fmt.Sprintf("message: %s - status: %d - error: %s - causes: %v",
		e.ErrMessage, e.ErrStatus, e.ErrError, e.ErrCauses)
}

func NewRestError(message string, status int, err string, causes []interface{}) ApiError {
	return apiError{
		ErrMessage: message,
		ErrStatus:  status,
		ErrError:   err,
		ErrCauses:  causes,
	}
}

func NewApiErrorFromBytes(bytes []byte) (ApiError, error) {
	var apiErr apiError
	if err := json.Unmarshal(bytes, &apiErr); err != nil {
		return nil, errors.New("invalid json")
	}
	return apiErr, nil
}

func NewBadRequestError(message string) ApiError {
	return apiError{
		ErrMessage: message,
		ErrStatus:  http.StatusBadRequest,
		ErrError:   "bad_request",
	}
}

func NewNotFoundError(message string) ApiError {
	return apiError{
		ErrMessage: message,
		ErrStatus:  http.StatusNotFound,
		ErrError:   "not_found",
	}
}

func NewUnauthorizedError(message string) ApiError {
	return apiError{
		ErrMessage: message,
		ErrStatus:  http.StatusUnauthorized,
		ErrError:   "unauthorized",
	}
}

func NewInternalServerError(message string, err error) ApiError {
	result := apiError{
		ErrMessage: message,
		ErrStatus:  http.StatusInternalServerError,
		ErrError:   "internal_server_error",
	}
	if err != nil {
		result.ErrCauses = append(result.ErrCauses, err.Error())
	}
	return result
}
