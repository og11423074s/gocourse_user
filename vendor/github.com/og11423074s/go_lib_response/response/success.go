package response

import (
	"encoding/json"
	"github.com/og11423074s/gocourse_meta/meta"
	"net/http"
)

type SuccessResponse struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Meta    *meta.Meta  `json:"meta,omitempty"`
}

func OK(msg string, data interface{}, meta *meta.Meta) Response {
	return success(msg, data, meta, http.StatusOK)
}

func Created(msg string, data interface{}, meta *meta.Meta) Response {
	return success(msg, data, meta, http.StatusCreated)
}

func Accepted(msg string, data interface{}, meta *meta.Meta) Response {
	return success(msg, data, meta, http.StatusAccepted)
}

func NonAuthoritativeInformation(msg string, data interface{}, meta *meta.Meta) Response {
	return success(msg, data, meta, http.StatusNonAuthoritativeInfo)
}

func NoContent(msg string, data interface{}, meta *meta.Meta) Response {
	return success(msg, data, meta, http.StatusNoContent)
}

func ResetContent(msg string, data interface{}, meta *meta.Meta) Response {
	return success(msg, data, meta, http.StatusResetContent)
}

func PartialContent(msg string, data interface{}, meta *meta.Meta) Response {
	return success(msg, data, meta, http.StatusPartialContent)
}

func success(msg string, data interface{}, meta *meta.Meta, code int) Response {
	return &SuccessResponse{
		Message: msg,
		Data:    data,
		Meta:    meta,
		Status:  code,
	}
}

// GetData override methods
func (s *SuccessResponse) GetData() interface{} {
	return s.Data
}

// Error override methods
func (s *SuccessResponse) Error() string {
	return ""
}

// GetBody override methods
func (s *SuccessResponse) GetBody() ([]byte, error) {
	return json.Marshal(s) // convert struct to json
}

// StatusCode override methods
func (s *SuccessResponse) StatusCode() int {
	return s.Status
}
