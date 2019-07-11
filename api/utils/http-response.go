package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type cuponServiceResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Result  interface{} `json:"result,omitempty"`
}

func WriteResponse(w http.ResponseWriter, code int, message string, result interface{}) {
	jsonResponse, err := json.Marshal(cuponServiceResponse{
		Code:    code,
		Message: message,
		Result:  result,
	})
	if err != nil {
		log.Print(err)
	}

	w.WriteHeader(code)
	w.Write(jsonResponse)
}
