package server

import (
	"encoding/json"
	"log"
	"net/http"
)

type ServerStatus struct {
	Status string `json:"status"`
}

func StatusCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(ServerStatus{"up"})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalf("could not encode json: %s\n", err.Error())
	}
}
