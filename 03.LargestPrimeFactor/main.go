/*
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/
package main

import (
	"fmt"
	"math"
)

func findTheLargestPrime(n int) int {
	a := int(math.Sqrt(float64(n)))
	for i := 2; i < a; i++ {
		if n%i == 0 {
			n /= i
			a = int(math.Sqrt(float64(n)))
		}
	}
	return n
}

func main() {
	prime := findTheLargestPrime(600851475143)
	fmt.Println(prime)
}
