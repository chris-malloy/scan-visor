package routes

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"testing"
)

func TestRoutes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Routes Test Suite")
}

var _ = Describe("The Route Callbacks", func() {
	Context("Home", func() {
		It("Is Sweet", func() {
			headers := make(http.Header)
			headers.Add("Content-type", "text/html")
			w, req := arrangeGet("/")

			Home(w, req)

			assert(w, http.StatusOK, headers, "notNil")
		})
	})
	
	Context("Health Check", func() {
		It("Responds Appropriately", func() {
			headers := make(http.Header)
			headers.Add("Content-type", "application/json")
			w, req := arrangeGet("/")

			HealthCheck(w, req)

			assert(w, http.StatusOK, headers, `{"status":"up"}`)
		})
	})
})
