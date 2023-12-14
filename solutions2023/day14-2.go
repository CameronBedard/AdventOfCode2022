package solutions2023

import (
	"fmt"
)

func Day14p2(lines []string) {
	grid := make([][]uint8, len(lines))
	for i := range lines {
		grid[i] = make([]uint8, len(lines[i]))
		for j := range grid[i] {
			grid[i][j] = lines[i][j]
		}
	}

	m := make(map[[100][100]uint8]int)
	rollNorth(grid)
	rotateClockwise(grid)
	rollNorth(grid)
	rotateClockwise(grid)
	rollNorth(grid)
	rotateClockwise(grid)
	rollNorth(grid)

	breakpoint := 1000000000
	for i := 1; i <= 1000000000; i++ {
		for j := 0; j < 4; j++ {
			rotateClockwise(grid)
			rollNorth(grid)
		}
		a := getKey(grid)
		if m[a] != 0 && breakpoint == 1000000000 {
			fmt.Println(i, m[a])
			breakpoint = i + (1000000000-m[a])%(i-m[a]) - 1
		} else {
			m[a] = i
		}
		if i == breakpoint {
			break
		}
	}
	rotateClockwise(grid)

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

func getKey(grid [][]uint8) [100][100]uint8 {
	var arr [100][100]uint8
	for i := range grid {
		copy(arr[i][:], grid[i])
	}

	return arr
}

func rollNorth(grid [][]uint8) {
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
}

func rotateClockwise(a [][]uint8) {
	N := len(a)
	for i := 0; i < N/2; i++ {
		for j := i; j < N-i-1; j++ {
			// Swap elements of each cycle
			// in clockwise direction
			temp := a[i][j]
			a[i][j] = a[N-1-j][i]
			a[N-1-j][i] = a[N-1-i][N-1-j]
			a[N-1-i][N-1-j] = a[j][N-1-i]
			a[j][N-1-i] = temp
		}
	}
}
