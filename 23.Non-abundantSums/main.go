/*
A perfect number is a number for which the sum of its proper divisors is exactly equal to the number.
For example, the sum of the proper divisors of 28 would be 1 + 2 + 4 + 7 + 14 = 28, which means that 28 is a perfect number.

A number n is called deficient if the sum of its proper divisors is less than n and it is called abundant if this sum exceeds n.

As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16, the smallest number that can be written as the sum of two abundant numbers is 24.
By mathematical analysis, it can be shown that all integers greater than 28123 can be written as the sum of two abundant numbers.
However, this upper limit cannot be reduced any further by analysis even though it is known that the greatest number that cannot be expressed as the sum of two abundant numbers is less than this limit.

Find the sum of all the positive integers which cannot be written as the sum of two abundant numbers.
*/
package main

import (
	"fmt"
	"math"
)

const (
	limit = 28123
)

var abundant [limit + 1]bool

func sumOfdivitors(n int) int {
	sum := 1
	sqrt := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqrt; i++ {
		if n%i == 0 {
			sum += i
			if i != (n / i) {
				sum += (n / i)
			}
		}
	}
	return sum
}

func isAbundant(n int) bool {
	return n < sumOfdivitors(n)
}

func initData() {
	for i := 12; i <= limit; i++ {
		if isAbundant(i) {
			abundant[i] = true
		}
	}
}

func nonAbundantSums() (sum int) {
	initData()
	for i := 1; i < 24; i++ {
		sum += i
	}
	for j := 25; j <= limit; j++ {
		for k := 12; k < j; k++ {
			if abundant[k] {
				if abundant[j-k] {
					goto Next
				}
			}
		}
		sum += j
	Next:
	}
	return
}

func main() {
	fmt.Println(nonAbundantSums())

}
