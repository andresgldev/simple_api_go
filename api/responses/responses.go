package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type errorResponse struct {
	Error string `json:"error"`
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JSON(w, statusCode, errorResponse{err.Error()})
		return
	}
	JSON(w, http.StatusBadRequest, nil)
}
