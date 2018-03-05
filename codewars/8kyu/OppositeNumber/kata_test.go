package kata

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Opposite Number")
}

var _ = Describe("Opposite function", func() {
	It("should invert number", func() {
		Expect(Opposite(1)).To(Equal(-1))
		Expect(Opposite(0)).To(Equal(0))
		Expect(Opposite(5)).To(Equal(-5))
		Expect(Opposite(-5)).To(Equal(5))
	})
})
