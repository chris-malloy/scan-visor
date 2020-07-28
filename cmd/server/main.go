package main

import (
	"fmt"
	"log"
	"net/http"
	"scan-visor/pkg/routes"

	// @todo figure out how to properly import other packages, probaly have to fix gopath
)

func main() {
	fmt.Println("Starting up Server for Scan Visor")

	http.HandleFunc("/", routes.Home)
	http.HandleFunc("/health", routes.HealthCheck)

	err := http.ListenAndServe(":9000", nil)

	if err != nil {
		log.Fatalf("could not start server: %s\n", err.Error())
	}

}
