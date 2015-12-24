/*
The number, 197, is called a circular prime because all rotations of the digits: 197, 971, and 719, are themselves prime.

There are thirteen such primes below 100: 2, 3, 5, 7, 11, 13, 17, 31, 37, 71, 73, 79, and 97.

How many circular primes are there below one million?
*/
package main

import (
	"math"
)

func isPrime(n int) bool {
	if n < 2 {
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

func isCircularPrime(n int) bool {
	numLen := 0
	temp := n
	if temp == 0 {
		numLen = 1
	}
	for temp > 0 {
		numLen++
		temp /= 10
	}
	if numLen == 1 && isPrime(n) {
		return true
	}
	temp = n
	var rotation, end int
	for i := 0; i < numLen; i++ {
		end = temp % 10
		rotation = temp / 10
		rotation += (int(math.Pow10(numLen-1)) * end)
		if !isPrime(rotation) {
			return false
		}
		temp = rotation
	}
	return true
}

func main() {
	count := 0
	for i := 1; i < 1000000; i++ {
		if isCircularPrime(i) {
			count++
		}
	}
	print(count)
}

/*
[ `go run main.go` | done: 875.050415ms ]
	55
*/
