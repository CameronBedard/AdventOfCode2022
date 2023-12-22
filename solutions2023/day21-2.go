package solutions2023

import (
	"fmt"
)

func Day21p2(lines []string) {
	start := Point{0, 0}
	for i := range lines {
		for j := range lines[0] {
			if lines[i][j] == 'S' {
				start = Point{i, j}
			}
		}
	}

	directions := Direction{
		north: Point{-1, 0},
		south: Point{1, 0},
		east:  Point{0, 1},
		west:  Point{0, -1},
	}

	depthDFS2(lines, start, directions, 0)

	x, a := 26501365/131, pointList[65]
	b := pointList[65+131] - a
	c := pointList[65+(131*2)] - pointList[65+131]

	lagrange := a + (b * x) + (x*(x-1)/2)*(c-b)

	fmt.Println("total", lagrange)
}

var pointList = make([]int, 66+(131*2))

func depthDFS2(grid []string, pos Point, directions Direction, steps int) {
	if state[dfsState{pos, steps}] || steps == 66+(131*2) {
		//continue
	} else {
		state[dfsState{pos, steps}] = true
		pointList[steps]++
		n, s, e, w := addPoints(pos, directions.north), addPoints(pos, directions.south), addPoints(pos, directions.east), addPoints(pos, directions.west)

		if !isFence(grid, n) {
			depthDFS2(grid, n, directions, steps+1)
		}
		if !isFence(grid, s) {
			depthDFS2(grid, s, directions, steps+1)
		}
		if !isFence(grid, e) {
			depthDFS2(grid, e, directions, steps+1)
		}
		if !isFence(grid, w) {
			depthDFS2(grid, w, directions, steps+1)
		}
	}
}

func isFence(grid []string, pos Point) bool {
	n := len(grid)
	return grid[((pos.y%n)+n)%n][((pos.x%n)+n)%n] == '#'
}
