package utils

import (
	"encoding/json"
	"metric-hub/model"
	"net/http"
)

func HttpResponse(w http.ResponseWriter, status bool, code int, err error, data interface{}) {

	response := model.HttpResponse{
		Status: status,
		Code:   code,
		Data:   data,
	}
	if err != nil {
		response.Error = err.Error()
	}

	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}
