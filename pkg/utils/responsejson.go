package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type jsonResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Meta    interface{} `json:"meta,omitempty"`
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(jsonResponse{
		Code:    code,
		Message: "Success",
		Data:    payload,
		Meta:    nil,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	write, err := w.Write(response)
	if err != nil {
		log.Println(write)
	}
}
