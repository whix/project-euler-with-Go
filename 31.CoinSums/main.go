/*
In England the currency is made up of pound, £, and pence, p, and there are eight coins in general circulation:

1p, 2p, 5p, 10p, 20p, 50p, £1 (100p) and £2 (200p).
It is possible to make £2 in the following way:

1×£1 + 1×50p + 2×20p + 1×5p + 1×2p + 3×1p
How many different ways can £2 be made using any number of coins?
*/

package main

var coins = []int{1, 2, 5, 10, 20, 50, 100, 200}

func countWays(pence, index int) (ways int) {
	if pence == 0 {
		return 1
	}
	if pence < 0 || index <= 0 {
		return 0
	}
	return countWays(pence, index-1) + countWays(pence-coins[index-1], index)
}

func main() {
	ways := countWays(200, 8)
	println(ways)
}

/*
[ `go run main.go` | done: 160.276724ms ]
	73682
*/
