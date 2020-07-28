package routes

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Routes Test Suite")
}

var _ = Describe("the route callbacks", func() {
	Context("home", func() {
		It("is sweet", func() {
			headers := make(http.Header)
			headers.Add("Content-type", "text/html")
			w, req := arrangeGet("/")

			Home(w, req)

			assert(w, http.StatusOK, headers, "notNil")
		})
	})
	Context("health check", func() {
		It("responds appropriately", func() {
			headers := make(http.Header)
			headers.Add("Content-type", "application/json")
			w, req := arrangeGet("/")

			HealthCheck(w, req)

			assert(w, http.StatusOK, headers, `{"status":"up"}`)
		})
	})
})
