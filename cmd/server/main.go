package main

import (
	"fmt"
	"log"
	"net/http"
	"scan-visor/pkg"
)

func main() {
	fmt.Println("Starting up Server for Scan Visor")

	http.HandleFunc("/healthcheck", server.HealthCheck)

	err := http.ListenAndServe(":9000", nil)

	if err != nil {
		log.Fatalf("could not start server: %s\n", err.Error())
	}

}
