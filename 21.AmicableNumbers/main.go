/*
Let d(n) be defined as the sum of proper divisors of n (numbers less than n which divide evenly into n).
If d(a) = b and d(b) = a, where a ≠ b, then a and b are an amicable pair and each of a and b are called amicable numbers.

For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110; therefore d(220) = 284. The proper divisors of 284 are 1, 2, 4, 71 and 142; so d(284) = 220.

Evaluate the sum of all the amicable numbers under 10000.
*/
package main

import (
	"fmt"
)

var sumOfDivisors [10000]int

func initData() {
	sum := 0
	for i := 2; i < 10000; i++ {
		sum = 0
		for j := 1; j < i; j++ {
			if i%j == 0 {
				sum += j
			}
		}
		sumOfDivisors[i] = sum
	}
}

func sumOfAmicableNumbers() (sum int) {
	initData()
	for i := 2; i < 10000; i++ {
		if sumOfDivisors[i] < 10000 && sumOfDivisors[i] > 1 {
			//exclude the situation 'd(a) = b and d(b) = a and a = b'
			if sumOfDivisors[i] != i && sumOfDivisors[sumOfDivisors[i]] == i {
				sum += i
			}
		}
	}
	return
}

func main() {
	num := sumOfAmicableNumbers()
	fmt.Println(num)
}
