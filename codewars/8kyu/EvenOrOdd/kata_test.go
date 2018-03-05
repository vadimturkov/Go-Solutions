package kata

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Return Negative")
}

var _ = Describe("Test Example", func() {
	It("should test that the solution returns the correct value", func() {
		Expect(EvenOrOdd(1)).To(Equal("Odd"))
		Expect(EvenOrOdd(2)).To(Equal("Even"))
		for i := 0; i < 1000; i++ {
			res := "Odd"
			if i % 2 == 0 {
				res = "Even"
			}
			Expect(EvenOrOdd(i)).To(Equal(res))
		}
	})
})