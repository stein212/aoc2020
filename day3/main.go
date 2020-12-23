package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	grid := make([]string, 0, 100)

	for {
		var row string
		_, err := fmt.Fscanln(reader, &row)

		if err != nil {
			break
		}

		grid = append(grid, row)
	}

	// fmt.Println(travese(grid, 3, 1))

	mul := travese(grid, 1, 1)
	mul *= travese(grid, 3, 1)
	mul *= travese(grid, 5, 1)
	mul *= travese(grid, 7, 1)
	mul *= travese(grid, 1, 2)

	fmt.Println(mul)
}

func travese(grid []string, right, down int) int {
	// assume non empty grid
	rows := len(grid)
	cols := len(grid[0])

	row, col := 0, 0
	numTrees := 0

	for row < rows {
		if grid[row][col] == '#' {
			numTrees++
		}

		col = (col + right) % cols
		row += down
	}

	return numTrees
}
