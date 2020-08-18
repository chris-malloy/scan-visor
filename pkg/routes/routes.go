package routes

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type HealthStatus struct {
	Status string `json:"status"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprint(w, "<h1>Scan Visor</h1><p>With the Scan Visor, Samus can scan certain objects to learn more about them or perform a function. When a scan has finished, the game action freezes to allow Samus to read its contents. If Samus scans a creature, enemy or boss, Samus will gain useful hints how to beat them, such as weak spots and behavioral patterns.</p><p>It also does auth.</p> ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalf("could not write response: %s\n", err.Error())
	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	writeJsonOrFail(HealthStatus{"up"}, w)
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	handleBadMethod("POST", w, r)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("error reading body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		writeJsonOrFail(ErrorResponse{"error reading request body"}, w)
	}

	fmt.Println(body)
	if len(body) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		writeJsonOrFail(ErrorResponse{"body cannot be empty"}, w)
	}
}
