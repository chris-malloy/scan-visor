package routes

import (
	. "github.com/onsi/gomega"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
)

func arrangeGet(route string) (*httptest.ResponseRecorder, *http.Request)  {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", route, nil)
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