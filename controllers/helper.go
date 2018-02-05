package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

func WriteResponse(w http.ResponseWriter, statusCode int, success bool, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	payload, err := GetResponsePayload(success, message, data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		payload = []byte("Server Error")
		fmt.Println(errors.Wrap(err, "Couldnt marshall json in writeResponse"))
	} else {
		w.WriteHeader(statusCode)
	}
	w.Write(payload)

}

func GetResponsePayload(success bool, message string, data interface{}) ([]byte, error) {
	res := Response{
		Status: ResponseStatus{
			Success: success,
			Message: message,
		},
		Data: data,
	}
	return json.Marshal(res)
}

type Response struct {
	Status ResponseStatus `json:"status"`
	Data   interface{}    `json:"data"`
}
type ResponseStatus struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
