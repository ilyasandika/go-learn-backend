package helper

import "uaspw2/models/web/response"

func CreateSuccessResponse(code int, message string, data interface{}) response.SuccessResponse {
	return response.SuccessResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func CreateErrorResponse(code int, message string, error interface{}) response.ErrorResponse {
	return response.ErrorResponse{
		Code:    code,
		Message: message,
		Error:   error,
	}
}
