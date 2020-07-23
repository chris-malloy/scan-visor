package server

import (
	"encoding/json"
	"log"
	"net/http"
)

type HealthStatus struct {
	Status string `json:"status"`
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(HealthStatus{"up"})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalf("could not encode json: %s\n", err.Error())
	}
}
