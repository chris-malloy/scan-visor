package integration

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"testing"
)

func TestAPI(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Scan Visor Integration Test Suite")
}

var _ bool = Describe("The API Server", func() {
	Context("While running and reachable", func() {
		It("Responds to the home route", func() {
			runGetTest("/", http.StatusOK, "notNil")
		})
	})

	Context("While running and reachable", func() {
		It("Responds to the health check route", func() {
			runGetTest("/health", http.StatusOK, `{"status":"up"}`)
		})
	})


})
