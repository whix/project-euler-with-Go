/*
n! means n × (n − 1) × ... × 3 × 2 × 1

For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.

Find the sum of the digits in the number 100!
*/
package main

import (
	"fmt"
)

var bigNum = [160]int{0: 1} // actully use 158

// func factorial(n int) int {
// 	if n == 0 || n == 1 {
// 		return 1
// 	} else {
// 		return factorial(n-1) * n
// 	}
// }

func factorialSpecial(n int) {
	carry := 0
	for i := 1; i <= n; i++ {
		for k, v := range bigNum {
			v = v*i + carry
			carry = v / 10
			bigNum[k] = v % 10
		}
	}
	fmt.Println(bigNum)
}

func factorialDigitSum(n int) (sum int) {
	factorialSpecial(n)
	for _, v := range bigNum {
		sum += v
	}
	return
}

func main() {
	sum := factorialDigitSum(100)
	fmt.Println(sum)
}
