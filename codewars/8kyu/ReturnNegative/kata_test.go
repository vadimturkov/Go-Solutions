package kata

import (
	"math/rand"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Return Negative")
}

func abs(x int) int {
	m := x >> 31
	return (x + m) ^ m
}

const NUM_RANDOM_TESTS = 10
const MAX_X = 1000

var _ = Describe("Test Example", func() {
	It("should test that the MakeNegative returns the negative value", func() {
		Expect(MakeNegative(42)).To(Equal(-42))
		Expect(MakeNegative(-9)).To(Equal(-9))
		Expect(MakeNegative(0)).To(Equal(0))
	})
	It("random positives", func() {
		for i:=0; i < NUM_RANDOM_TESTS; i++ {
			x := rand.Intn(MAX_X - 1) + 1
			Expect(MakeNegative(x)).To(Equal(-abs(x)))
		}
	})
	It("random negatives", func() {
		for i:=0; i < NUM_RANDOM_TESTS; i++ {
			x := -rand.Intn(MAX_X)
			Expect(MakeNegative(x)).To(Equal(-abs(x)))
		}
	})
})