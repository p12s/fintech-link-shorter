package handler

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

type errorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	logrus.Error(message)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(errorResponse{Message: message}); err != nil {
		http.Error(w, "Error: an error occurred - the link could not be shortened.", http.StatusInternalServerError)
		return
	}
}