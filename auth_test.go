package scan_visor

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestAuth(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Auth Suite")
}

var _ = Describe("username/password auth", func() {
	Context("when passed bad challenge", func() {
		It("fails auth", func() {
			ok, err := authenticate(challenge{"", ""})
			Expect(ok).To(BeFalse())
			Expect(err).ToNot(BeNil())
		})
	})
})

type challenge struct {
	username string
	password string
}

func authenticate(challenge challenge) (bool, error) {
	return false, fmt.Errorf("recieved bad challenge")
}


