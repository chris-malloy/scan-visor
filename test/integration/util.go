package integration

import (
	. "github.com/onsi/ginkgo"
	"net/http"
	"os"
	"strings"
)

func get(route string) *http.Response {
	prefixCheck(route)

	port := os.Getenv("API_TCP_PORT")
	if port == "" {
		port = "9000"
		//GinkgoT().Fatalf("error: API_TCP_PORT is not set")
	}
	response, err := http.Get("http://localhost:" + port + route)

	if err != nil {
		GinkgoT().Fatalf("caught error when creating the request: %v", err.Error())
	}

	return response
}

func prefixCheck(route string) {
	if !strings.HasPrefix(route, "/") {
		GinkgoT().Fatalf("routes must start with a forward slash. Offender: %v", route)
	}
}
