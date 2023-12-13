package solutions2023

import (
	"fmt"
	"slices"
)

func Day13p2(lines []string) {
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
		prevX := isVertSymetric(grids[i])
		prevY := isHorizSymetric(grids[i])
	Sim:
		for j := range grids[i] {
			for k := range grids[i][0] {
				newGrid := make([]string, len(grids[i]))
				copy(newGrid, grids[i])
				if newGrid[j][k] == '.' {
					newGrid[j] = newGrid[j][0:k] + "#" + newGrid[j][k+1:]
				} else {
					newGrid[j] = newGrid[j][0:k] + "." + newGrid[j][k+1:]
				}

				x := isVertSymetric2(newGrid)
				y := isHorizSymetric2(newGrid)

				if len(x) > 0 {
					for match := range x {
						if x[match] != prevX {
							sum += x[match]
							break Sim
						}
					}
				}
				if len(y) > 0 {
					for match := range y {
						if y[match] != prevY {
							sum += y[match] * 100
							break Sim
						}
					}
				}
				if j == len(grids[i])-1 && k == len(grids[i][0])-1 {
					fmt.Println("found no match", i, len(grids[i]), len(grids[i][0]))
				}
			}
		}
	}
	fmt.Println("total", sum)
}

func print2d(str []string) {
	for i := range str {
		fmt.Println(str[i])
	}
}

func isVertSymetric2(grid []string) []int {
	mirror := make([]int, 0)
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
			mirror = append(mirror, marker)
		}
	}
	return mirror
}

func isHorizSymetric2(grid []string) []int {
	mirror := make([]int, 0)
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
			mirror = append(mirror, marker)
		}
	}
	return mirror
}
