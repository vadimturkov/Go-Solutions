// Solution to CodeWars kata: To square(root) or not to square(root)
// https://www.codewars.com/kata/to-square-root-or-not-to-square-root/

package kata

import "math"

func SquareOrSquareRoot(arr []int) []int {
	var result []int

	for _, number := range arr {
		if sqrt := math.Sqrt(float64(number)); sqrt == math.Trunc(sqrt) {
			result = append(result, int(sqrt))
		} else {
			result = append(result, number * number)
		}
	}

	return result
}