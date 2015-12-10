/*
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
Find the largest palindrome made from the product of two 3-digit numbers.
*/
package main

import (
	"fmt"
	// "strconv"
)

// func isPalidrome(n int) bool {
// 	str := strconv.Itoa(n)
// 	length := len(str)
// 	var k int = length / 2
// 	for i := 0; i < k; i++ {
// 		if str[i] != str[length-i-1] {
// 			return false
// 		}
// 	}
// 	return true

// }

func reverse(n int) (reversed int) {
	for n > 0 {
		reversed = 10*reversed + n%10
		n /= 10
	}
	return
}

func isPalidrome(num int) bool {
	return num == reverse(num)
}

func findThePalidrome() (num1, num2, pal int) {
	temp := 0
	pal = 0
	for i := 101; i < 1000; i++ {
		for j := 101; j < 1000; j++ {
			temp = i * j
			if isPalidrome(temp) && temp > pal {
				pal = temp
				num1 = i
				num2 = j
			}
		}
	}
	return
}
func main() {
	a, b, c := findThePalidrome()
	fmt.Println(a, b, c)
}
