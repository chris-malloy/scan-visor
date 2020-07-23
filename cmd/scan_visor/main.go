package main

import (
	"fmt"
	"log"
	"net/http"
	"scan-visor/pkg"
)

func main() {
	fmt.Println("The Scan Visor")

	http.HandleFunc("/hello", server.StatusCheck)

	err := http.ListenAndServe(":9000", nil)

	if err != nil {
		log.Fatalf("could not start server: %s\n", err.Error())
	}

}
