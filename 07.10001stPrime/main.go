/*
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
What is the 10 001st prime number?
*/
package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n == 1 {
		return false
	}
	if n == 2 {
		return true
	}
	var a int = 0
	a = int(math.Sqrt(float64(n)))
	for i := 2; i <= a; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func find10001stPrime() (num int) {
	cnt := 0
	for i := 1; true; i++ {
		if isPrime(i) {
			cnt++
			if cnt == 10001 {
				num = i
				goto A
			}
		}
	}
A:
	return num
}

func main() {
	num := find10001stPrime()
	fmt.Println(num)
}
