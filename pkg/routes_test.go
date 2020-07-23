package server

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestServer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "UnitTestSuite")
}

var _ = Describe("the route callbacks", func() {
	Context("health check", func() {
		It("responds appropriately", func() {
			req := httptest.NewRequest("GET", "/healthcheck", nil)
			w := httptest.NewRecorder()
			HealthCheck(w, req)

			res := w.Result()
			body, _ := ioutil.ReadAll(res.Body)

			Expect(res.StatusCode).To(Equal(200))
			Expect(res.Header.Get("Content-Type")).To(Equal("application/json"))
			Expect(strings.TrimSpace(string(body))).To(Equal(`{"status":"up"}`))
		})
	})
})
