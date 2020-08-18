package routes

import (
	. "github.com/onsi/gomega"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
)

func arrangeRequest(method string, route string, body io.Reader) (*httptest.ResponseRecorder, *http.Request)  {
	w := httptest.NewRecorder()
	req := httptest.NewRequest(method, route, body)
	return w, req
}

func assert(w *httptest.ResponseRecorder, expectedStatus int, expectedHeaders http.Header, expectedBody string) {
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	Expect(res.StatusCode).To(Equal(expectedStatus))
	for headerKey, headerValue := range expectedHeaders {
		Expect(res.Header.Get(headerKey)).To(Equal(headerValue[0]))
	}

	if expectedBody == "notNil" {
		Expect(body).ToNot(BeNil())
	} else {
		Expect(strings.TrimSpace(string(body))).To(Equal(expectedBody))
	}
}