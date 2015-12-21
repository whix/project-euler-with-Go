/*
Euler discovered the remarkable quadratic formula:

n² + n + 41

It turns out that the formula will produce 40 primes for the consecutive values n = 0 to 39. However, when n = 40, 402 + 40 + 41 = 40(40 + 1) + 41 is divisible by 41, and certainly when n = 41, 41² + 41 + 41 is clearly divisible by 41.

The incredible formula  n² − 79n + 1601 was discovered, which produces 80 primes for the consecutive values n = 0 to 79. The product of the coefficients, −79 and 1601, is −126479.

Considering quadratics of the form:

n² + an + b, where |a| < 1000 and |b| < 1000

where |n| is the modulus/absolute value of n
e.g. |11| = 11 and |−4| = 4
Find the product of the coefficients, a and b, for the quadratic expression that produces the maximum number of primes for consecutive values of n, starting with n = 0.
*/
package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	sqrt := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqrt; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func findABValue() (a, b int) {
	count := 0
	for i := -999; i < 1000; i++ {
		for j := -999; j < 1000; j++ {
			for k := 0; true; k++ {
				result := k*k + i*k + j
				if !isPrime(result) {
					if (k - 1) > count {
						count = k - 1
						a = i
						b = j
					}
					break
				}
			}
		}
	}
	return
}

func main() {
	a, b := findABValue()
	fmt.Println(a, b, a*b)
}

/*
[ `go run main.go` | done: 436.05055ms ]
	-61 971 -59231
*/
