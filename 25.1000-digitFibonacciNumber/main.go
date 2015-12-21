/*
Hence the first 12 terms will be:

F1 = 1
F2 = 1
F3 = 2
F4 = 3
F5 = 5
F6 = 8
F7 = 13
F8 = 21
F9 = 34
F10 = 55
F11 = 89
F12 = 144
The 12th term, F12, is the first term to contain three digits.

What is the index of the first term in the Fibonacci sequence to contain 1000 digits?
*/
package main

import (
	"fmt"
)

func update(n []int) {
	carry := 0
	for k, v := range n {
		temp := carry + v
		carry = temp / 10
		n[k] = temp % 10
	}
}

func fibonacciSpecial() (result int) {

	a := [1000]int{0: 1}
	b := [1000]int{0: 1}
	var fib [1000]int
	for i := 3; true; i++ {
		for j := 0; j < 1000; j++ {
			fib[j] = a[j] + b[j]
		}
		update(fib[:])
		if fib[999] > 0 {
			result = i
			goto OUT
		}
		a = b
		b = fib
	}
OUT:
	return
}

func main() {
	fmt.Println(fibonacciSpecial())
}
