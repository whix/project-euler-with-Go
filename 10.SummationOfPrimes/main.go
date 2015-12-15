/*
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.
*/
package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n < 0 || n == 1 {
		return false
	}
	temp := int(math.Sqrt(float64(n)))
	for i := 2; i <= temp; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func summationOfPrimes(n int) (result int) {
	for i := 1; i < n; i++ {
		if isPrime(i) {
			result += i
		}
	}
	return
}

func main() {
	result := summationOfPrimes(2000000)
	fmt.Println(result)
}
