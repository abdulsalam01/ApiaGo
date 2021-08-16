package utils

import (
	"encoding/json"
	"net/http"
)

func JsonResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func ErrorResponse(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JsonResponse(w, statusCode, struct {
			Message string
		}{
			Message: err.Error(),
		})

		return
	}

	JsonResponse(w, http.StatusBadRequest, nil)
}
