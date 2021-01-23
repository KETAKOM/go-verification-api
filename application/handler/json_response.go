package handler

import (
	"encoding/json"
	"net/http"
)

type ResponseData struct {
	Status int         "json: status"
	Data   interface{} "json: data"
}

func ResponseJson(w http.ResponseWriter, status int, data interface{}) {
	response := ResponseData{status, data}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
