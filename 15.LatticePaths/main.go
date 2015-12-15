/*
Starting in the top left corner of a 2×2 grid, and only being able to move to the right and down,
there are exactly 6 routes to the bottom right corner.
(...)
How many such routes are there through a 20×20 grid?
*/
package main

import (
	"fmt"
)

func countLatticePaths(row, col int) int {
	if row == 1 || col == 1 {
		return row + col
	} else {
		return countLatticePaths(row-1, col) + countLatticePaths(row, col-1)
	}
}

func main() {
	paths := countLatticePaths(20, 20) //paths = 137846528820, cost 2m49s
	fmt.Println(paths)
}
