package solutions2023

import (
	"fmt"
)

func Day14p1(lines []string) {
	grid := make([][]uint8, len(lines))
	for i := range lines {
		grid[i] = make([]uint8, len(lines[i]))
		for j := range grid[i] {
			grid[i][j] = lines[i][j]
		}
	}

	for j := 0; j < len(grid[0]); j++ {
		lastStop := 0
		for i := 0; i < len(grid); i++ {
			if grid[i][j] == '#' {
				lastStop = i + 1
			} else if grid[i][j] == 'O' {
				grid[i][j] = '.'
				grid[lastStop][j] = 'O'
				lastStop++
			}
		}
	}

	//print2darr(grid)
	sum := 0
	for j := 0; j < len(grid[0]); j++ {
		for i := 0; i < len(grid); i++ {
			if grid[i][j] == 'O' {
				sum += len(grid) - i
			}
		}
	}

	fmt.Println("total", sum)
}

func print2darr(arr [][]uint8) {
	for i := range arr {
		for j := range arr[i] {
			fmt.Print(string(arr[i][j]))
		}
		fmt.Println()
	}
}
