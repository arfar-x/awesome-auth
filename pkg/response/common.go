package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
	Err     string `json:"error"`
}

func NewResponse(message string, code int, data any, err error) *Response {
	var errorMessage string
	if err != nil {
		errorMessage = err.Error()
	}
	return &Response{
		Code:    code,
		Message: message,
		Data:    data,
		Err:     errorMessage,
	}
}

func JsonResponse(message string, code int, data any, err error) *Response {
	response := NewResponse(message, code, data, err)
	if _, err := json.Marshal(response); err != nil {
		panic("Json response binding failed.")
	}
	return response
}

func Success(message string, data any) (int, *Response) {
	return http.StatusOK, JsonResponse(message, http.StatusOK, data, nil)
}

func CreatedResource(message string, data any) (int, *Response) {
	return http.StatusCreated, JsonResponse(message, http.StatusCreated, data, nil)
}

func Unauthorized(message string, data any) (int, *Response) {
	return http.StatusUnauthorized, JsonResponse(message, http.StatusUnauthorized, data, nil)
}

func InternalError(message string, data any) (int, *Response) {
	return http.StatusInternalServerError, JsonResponse(message, http.StatusInternalServerError, data, nil)
}
