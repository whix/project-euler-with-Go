/*
145 is a curious number, as 1! + 4! + 5! = 1 + 24 + 120 = 145.

Find the sum of all numbers which are equal to the sum of the factorial of their digits.

Note: as 1! = 1 and 2! = 2 are not sums they are not included.
*/
package main

// 9!=362880, 9!*7=2540160 < 9999999... therefore, the number less than 9999999
// 不知道能不能把界限再继续缩小
const limit = 9999999

func factorial(n int) int {
	if n == 1 || n == 0 {
		return 1
	} else {
		return factorial(n-1) * n
	}
}

func isCuriousNumber(n int) bool {
	if n == 1 || n == 2 {
		return false
	}
	temp := n
	sum := 0
	digit := 0
	for temp > 0 {
		digit = temp % 10
		temp /= 10
		sum += factorial(digit)

	}
	return sum == n
}

func sumOfDigits() (result int) {
	for i := 1; i < limit; i++ {
		if isCuriousNumber(i) {
			result += i
		}
	}
	return
}

func main() {
	print(sumOfDigits())
}

/*
[ `go run main.go` | done: 1.433626622s ]
	40730
*/
