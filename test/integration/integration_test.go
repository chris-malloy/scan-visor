package integration

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestAPI(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Scan Visor Integration Test Suite")
}

var _ bool = Describe("The API Server", func() {
	Context("While running and reachable", func() {
		It("Responds to the health check route", func() {
			runTest("/health", http.StatusOK, `{"status":"up"}`)
		})
	})
})

func runTest(route string, expectedStatus int, expectedBody string) {
	response := get("/health")

	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	Expect(response.StatusCode).To(Equal(http.StatusOK))
	Expect(err).To(BeNil(), "Reading body from response failed")
	Expect(strings.TrimSpace(string(body))).To(Equal(`{"status":"up"}`))
}
