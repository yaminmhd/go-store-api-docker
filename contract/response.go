package contract

import (
	"encoding/json"
	"net/http"
)

type ErrorList struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
}

type BaseResponse struct {
	Success bool        `json:"success"`
	Errors  []ErrorList `json:"errors,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func SuccessfulResponse(w http.ResponseWriter, responseData interface{}, success string) {
	successInformation := successObjects[success]

	w.WriteHeader(successInformation.status)
	response := BaseResponse{
		Success: true,
		Data:    responseData,
	}
	responseJSON, _ := json.Marshal(response)
	w.Write(responseJSON)
	return
}

func ErrorResponse(w http.ResponseWriter, errors []string, httpStatus int) {
	w.WriteHeader(httpStatus)
	var errorList []ErrorList
	for _, err := range errors {
		errorList = append(errorList, ErrorList{Message: err})
	}
	response := BaseResponse{
		Success: false,
		Errors:  errorList,
	}
	responseJSON, _ := json.Marshal(response)

	w.Write(responseJSON)
	return
}
