/*
The number 3797 has an interesting property. Being prime itself, it is possible to continuously remove digits from left to right, and remain prime at each stage: 3797, 797, 97, and 7. Similarly we can work from right to left: 3797, 379, 37, and 3.

Find the sum of the only eleven primes that are both truncatable from left to right and right to left.

NOTE: 2, 3, 5, and 7 are not considered to be truncatable primes.
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

func isLeftToRight(n, length int) bool {
	newN := 0
	for i := 1; i < length; i++ {
		newN = n % int(math.Pow10(length-i))
		if !isPrime(newN) {
			return false
		}
	}
	return true
}

func isRightToLeft(n, length int) bool {
	newN := 0
	for i := 1; i < length; i++ {
		newN = n / 10
		if !isPrime(newN) {
			return false
		}
		n = newN
	}
	return true
}

func isTruncatablePrime(n int) bool {
	if !isPrime(n) {
		return false
	}
	length, temp := 0, n
	if temp == 0 {
		length = 1
	}
	for temp > 0 {
		length++
		temp /= 10
	}
	if length == 1 {
		return false
	} else {
		if isLeftToRight(n, length) && isRightToLeft(n, length) {
			return true
		} else {
			return false
		}
	}
}

func main() {
	count, sum := 0, 0
	for i := 10; true; i++ {
		if isTruncatablePrime(i) {
			count++
			sum += i
		}
		if count == 11 {
			break
		}
	}
	println(sum)
}

/*
[ `go run main.go` | done: 604.837395ms ]
	748317
*/
