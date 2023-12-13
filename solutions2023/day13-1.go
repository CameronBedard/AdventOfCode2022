package solutions2023

import (
	"fmt"
	"slices"
)

func Day13p1(lines []string) {
	grids := make([][]string, 1)
	gridI := 0
	for i := range lines {
		if lines[i] == "" {
			gridI++
			grids = append(grids, make([]string, 0))
		} else {
			grids[gridI] = append(grids[gridI], lines[i])
		}
	}
	//all lines that have a reflected equal must match
	sum := 0
	for i := range grids {
		x := isVertSymetric(grids[i])
		y := isHorizSymetric(grids[i])

		if x != -1 {
			sum += x
		}
		if y != -1 {
			sum += y * 100
		}
	}
	fmt.Println("total", sum)
}

func isVertSymetric(grid []string) int {
	mirror := -1
	for marker := 1; marker < len(grid[0]); marker++ {
		p1, p2 := marker-1, marker
		//start counting cols
		for p2 < len(grid[0]) && p1 >= 0 {
			str1, str2 := make([]uint8, len(grid)), make([]uint8, len(grid))
			for i := range grid {
				str1[i] = grid[i][p1]
				str2[i] = grid[i][p2]
			}
			if !slices.Equal(str1, str2) {
				break
			}
			p2++
			p1--
		}
		if p2 == len(grid[0]) || p1 == -1 { //if we reached the end
			return marker
		}
	}
	return mirror
}

func isHorizSymetric(grid []string) int {
	mirror := -1
	for marker := 1; marker < len(grid); marker++ {
		p1, p2 := marker-1, marker
		//start counting cols
		for p2 < len(grid) && p1 >= 0 {
			if grid[p1] != grid[p2] {
				break
			}
			p2++
			p1--
		}
		if p2 == len(grid) || p1 == -1 { //if we reached the end
			return marker
		}
	}
	return mirror
}
