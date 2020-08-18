package routes

import (
	"encoding/json"
	"log"
	"net/http"
)

func writeJsonOrFail(writable interface{}, w http.ResponseWriter) {
	err := json.NewEncoder(w).Encode(writable)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalf("could not encode json: %s\n", err.Error())
	}
}

func handleBadMethod(expectedMethod string, w http.ResponseWriter, r *http.Request) {
	if r.Method != expectedMethod {
		w.WriteHeader(http.StatusMethodNotAllowed)
		err := json.NewEncoder(w).Encode(ErrorResponse{"method is not allowed"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatalf("could not encode json: %s\n", err.Error())
		}
	}
}