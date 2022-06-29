package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/riku-smile/go-stations/model"
)

// A HealthzHandler implements health check endpoint.
type HealthzHandler struct{}

// NewHealthzHandler returns HealthzHandler based http.Handler.
func NewHealthzHandler() *HealthzHandler {
	return &HealthzHandler{}
}

// ServeHTTP implements http.Handler interface.
func (h *HealthzHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hr := &model.HealthzResponse{
		Message: "OK",
	}
	if err := json.NewEncoder(w).Encode(hr); err != nil {
		log.Println(err)
	}
}
