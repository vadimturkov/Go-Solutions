// Solution to CodeWars kata: Even or Odd
// https://www.codewars.com/kata/even-or-odd/

package kata

func EvenOrOdd(number int) string {
	if number % 2 == 0 {
		return "Even"
	}

	return "Odd"
}