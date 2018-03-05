package kata

import (
	"fmt"
	"math/rand"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Multiply of Index")
}

func solution(array []int) []int {
	result := []int{}
	for i, num := range array {
		if i > 0 {
			if num%i == 0 {
				result = append(result, num)
			}
		}
	}
	return result
}

var _ = Describe("Basic tests", func() {
	It("should return the correct values", func() {
		Expect(multipleOfIndex([]int{22, -6, 32, 82, 9, 25})).To(ConsistOf(-6, 32, 25))
		Expect(multipleOfIndex([]int{68, -1, 1, -7, 10, 10})).To(ConsistOf(-1, 10))
		Expect(multipleOfIndex([]int{11, -11})).To(ConsistOf(-11))
		Expect(multipleOfIndex([]int{-56, -85, 72, -26, -14, 76, -27, 72, 35, -21, -67, 87, 0, 21, 59, 27, -92, 68})).To(ConsistOf(-85, 72, 0, 68))
		Expect(multipleOfIndex([]int{28, 38, -44, -99, -13, -54, 77, -51})).To(ConsistOf(38, -44, -99))
		Expect(multipleOfIndex([]int{-1, -49, -1, 67, 8, -60, 39, 35})).To(ConsistOf(-49, 8, -60, 35))
	})
})

var _ = Describe("passes random tests", func() {
	randomInteger := func(a, b int) int {
		return rand.Intn((b-a)+1) + a
	}

	for i := 1; i <= 94; i++ {
		var nums []int
		lengthNums := randomInteger(2, 20)

		for j := 1; j <= lengthNums; j++ {
			nums = append(nums, randomInteger(-100, 100))
		}

		It(fmt.Sprintf(`Testing №%v for %v`, i+3, nums), func() {
			Expect(multipleOfIndex(nums)).To(Equal(solution(nums)))
		})
	}
})