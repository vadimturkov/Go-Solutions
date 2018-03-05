package kata

import (
	"math"
	"math/rand"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Square or not")
}

var _ = Describe("Example Tests", func() {
	tests := [][][]int{
		{{4, 3, 9, 7, 2, 1},{2, 9, 3, 49, 4, 1}},
		{{100, 101, 5, 5, 1, 1},{10, 10201, 25, 25, 1, 1}},
		{{1, 2, 3, 4, 5, 6},{1, 4, 9, 2, 25, 36}},
	}
	It("should test that the solution returns the correct value", func() {
		for _, t := range tests {
			Expect(SquareOrSquareRoot(t[0])).To(Equal(t[1]))
		}
	})
})

var _ = Describe("Random Tests", func(){
	It("should test that the solution returns the correct value with random numbers", func() {
		for i := 0; i < 15; i++ {
			arr := testcase()
			Expect(SquareOrSquareRoot(arr)).To(Equal(reference(arr)))
		}
	})
})

func testcase() []int {
	rang := rand.Intn(20) + 3
	test := make([]int, rang)
	for i := 0; i < rang; i++ {
		test[i] = rand.Intn(100) + 1
	}
	return test
}

func reference(arr []int) []int{
	var r []int
	for _, v := range arr {
		if math.Sqrt(float64(v)) == float64(int(math.Sqrt(float64(v)))) {
			r = append(r, int(math.Sqrt(float64(v))))
		} else {
			r = append(r, v * v)
		}
	}
	return r
}