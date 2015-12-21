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

/*
This is easy enough to solve analytically instead of programmatically
There are n! permutations for n digits

10! = 3628800
9!  = 362880
8!  = 40320
7!  = 5040
6!  = 720
5!  = 120
4!  = 24
3!  = 6
2!  = 2
1!  = 1

There are 362880 permutations for each leading digit
There are 40320 permutations for each leading digit pair
There are 5040 permutations for each leading digit trio
...
There are 1 permutations for each leading digit 9-tuple

Since 1000000 < 3*9!, the 1000000th starts with 3rd lowest available digit
Since 1000000 < 7*8!, the 1000000th continues with the 7th lowest available digit
...

To find first number, we would solve 0 = 0*(9!) + 0*(8!) + 0(7!) + 0*(6!) + 0*(5!) 0*(4!) + 0*(3!) + 0*(2!) + 0*(1!)
Which would correspond to choosing the 1st available number for each digit.

So, for the millionth, we have: 999999 = 2*(9!) + 6*(8!) + 6*(7!) + 2*(6!) + 5*(5!) 1*(4!) + 2*(3!) + 1*(2!) + 1*(1!)

[0,1,2,3,4,5,6,7,8,9] choose 3rd lowest: 2
[0,1,3,4,5,6,7,8,9]	choose 7th lowest: 7
[0,1,3,4,5,6,8,9]	choose 7th lowest: 8
[0,1,3,4,5,6,9]		choose 3rd lowest: 3
[0,1,4,5,6,9]		choose 6th lowest: 9
[0,1,4,5,6]			choose 2nd lowest: 1
[0,4,5,6]			choose 3rd lowest: 5
[0,4,6]			choose 2nd lowest: 4
[0,6]				choose 2nd lowest: 6
[0]				choose remaining : 0

So, answer is 2783915460
*/
