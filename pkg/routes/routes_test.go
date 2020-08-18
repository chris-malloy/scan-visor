package routes

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"strings"
	"testing"
)

func TestRoutes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Routes Test Suite")
}

var _ = Describe("The Utility Routes", func() {
	Context("Home", func() {
		It("Is Sweet", func() {
			headers := make(http.Header)
			headers.Add("Content-type", "text/html")
			w, req := arrangeRequest("GET", "/", nil)

			Home(w, req)

			assert(w, http.StatusOK, headers, "notNil")
		})
	})
	
	Context("Health Check", func() {
		It("Responds Appropriately", func() {
			headers := make(http.Header)
			headers.Add("Content-type", "application/json")
			w, req := arrangeRequest("GET", "/health", nil)

			HealthCheck(w, req)

			assert(w, http.StatusOK, headers, `{"status":"up"}`)
		})
	})
})

var _ = Describe("The Login Route", func() {
	Context("when making a non-POST request", func() {
		It("responds with method not allowed", func() {
			headers := make(http.Header)
			headers.Add("Content-type", "application/json")
			body := strings.NewReader("{}")
			w, req := arrangeRequest("GET", "/login", body)

			Login(w, req)

			assert(w, http.StatusMethodNotAllowed, headers, `{"error":"method is not allowed"}`)
		})
	})

	Context("when body is empty", func() {
		It("responds with bad request", func() {
			headers := make(http.Header)
			headers.Add("Content-type", "application/json")
			w, req := arrangeRequest("POST", "/login", nil)

			Login(w, req)

			assert(w, http.StatusBadRequest, headers, `{"error":"body cannot be empty"}`)
		})
	})
})
