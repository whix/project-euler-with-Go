/*
2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

What is the sum of the digits of the number 2^1000?
*/
package main

import (
	"fmt"
)

var bignum = [500]int{0: 1}

func mul2(n []int) {
	carry, curr := 0, 0
	for k, v := range n {
		curr = v * 2
		n[k] = curr%10 + carry
		carry = curr / 10
	}

}

func countSum(n []int) (sum int) {
	for _, v := range n {
		sum += v
	}
	return
}

func main() {
	for i := 0; i < 1000; i++ {
		mul2(bignum[:])
	}
	sum := countSum(bignum[:])
	fmt.Println(sum)
}
