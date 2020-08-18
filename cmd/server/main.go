package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"scan-visor/pkg/routes"

	// @todo figure out how to properly import other packages, probaly have to fix gopath
)

func main() {
	fmt.Println("Starting up Server for Scan Visor")

	http.HandleFunc("/", routes.Home)
	http.HandleFunc("/health", routes.HealthCheck)

	http.HandleFunc("/login", routes.Login)

	port := os.Getenv("API_TCP_PORT")
	if port == "" {
		log.Println("API_TCP_PORT not set explicitly, defaulting to 9000")
		port = "9000"
	}

	err := http.ListenAndServe(":" + port, nil)
	if err != nil {
		log.Fatalf("could not start server: %s\n", err.Error())
	}

}
