/*
A permutation is an ordered arrangement of objects. For example, 3124 is one possible permutation of the digits 1, 2, 3 and 4.
If all of the permutations are listed numerically or alphabetically, we call it lexicographic order.
The lexicographic permutations of 0, 1 and 2 are:

012   021   102   120   201   210

What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?
*/
package main

var used []bool

func main() {
	used = make([]bool, 10)
	for i := range used {
		used[i] = false
	}
	remaining := 999999
	perm := 1
	index := 0
	for i := 9; i > 0; i-- {
		perm = 1
		for j := i; j > 0; j-- {
			perm *= j
		}
		index = remaining / perm
		remaining = remaining % perm
		j := 0
		for j = 0; j < index || used[j]; j++ {
			if used[j] {
				index++
			}
		}
		used[j] = true
		print(j)
	}
	println(remaining)
}
