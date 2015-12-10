/*
The sum of the squares of the first ten natural numbers is,

1^2 + 2^2 + ... + 10^2 = 385
The square of the sum of the first ten natural numbers is,

(1 + 2 + ... + 10)^2 = 55^2 = 3025
Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
*/
package main

import (
	"fmt"
	"math"
)

func diffBetweenTwoSquare() int {
	sumOfSquare, squareOfSum := 0, 0
	for i := 1; i <= 100; i++ {
		sumOfSquare += int(math.Pow(float64(i), 2))
		squareOfSum += i
	}
	squareOfSum = int(math.Pow(float64(squareOfSum), 2))
	return squareOfSum - sumOfSquare
}

func main() {
	diff := diffBetweenTwoSquare()
	fmt.Println(diff)
}
