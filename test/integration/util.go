package integration

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func get(route string) *http.Response {
	prefixCheck(route)

	port := os.Getenv("API_TCP_PORT")
	if port == "" {
		port = "9000"
	}

	response, err := http.Get("http://localhost:" + port + route)
	if err != nil {
		GinkgoT().Fatalf("caught error when creating the request: %v", err.Error())
	}

	return response
}

func runGetTest(route string, expectedStatus int, expectedBody string) {
	response := get(route)

	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	Expect(response.StatusCode).To(Equal(expectedStatus))
	Expect(err).To(BeNil(), "Reading body from response failed")
	if expectedBody == "notNil" {
		Expect(body).ToNot(BeNil())
	} else {
		Expect(strings.TrimSpace(string(body))).To(Equal(expectedBody))
	}
}

func prefixCheck(route string) {
	if !strings.HasPrefix(route, "/") {
		GinkgoT().Fatalf("routes must start with a forward slash. Offender: %v", route)
	}
}
