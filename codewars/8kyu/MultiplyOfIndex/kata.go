// Solution to CodeWars kata: Multiple of index
// https://www.codewars.com/kata/multiple-of-index/

package kata

func multipleOfIndex(ints []int) []int {
	var out []int

	for index, value := range ints {
		if index > 0 && value % index == 0 {
			out = append(out, value)
		}
	}

	return out
}
