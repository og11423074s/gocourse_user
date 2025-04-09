package response

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func InternalServerError(message string) Response {
	return errors(message, http.StatusInternalServerError)
}

func NotFound(message string) Response {
	return errors(message, http.StatusNotFound)
}

func Unauthorized(message string) Response {
	return errors(message, http.StatusUnauthorized)
}

func Forbidden(message string) Response {
	return errors(message, http.StatusForbidden)
}

func BadRequest(message string) Response {
	return errors(message, http.StatusBadRequest)
}

func errors(message string, status int) Response {
	return &ErrorResponse{
		Status:  status,
		Message: message,
	}
}

func (e *ErrorResponse) Error() string {
	return e.Message
}

func (e *ErrorResponse) GetBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e *ErrorResponse) StatusCode() int {
	return e.Status
}

func (e *ErrorResponse) GetData() interface{} {
	return nil
}
