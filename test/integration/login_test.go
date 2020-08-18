package integration

import . "github.com/onsi/ginkgo"

var _ = Describe("The Login Route", func() {
	Context("When supplied no body", func() {
		It("Responds with bad request", func() {

		})
	})

	Context("When user session expired", func() {
		It("Responds with unauthorized", func() {

		})
	})

	Context("When supplied a bad user/pass", func() {
		It("Responds with forbidden ", func() {

		})
	})


})