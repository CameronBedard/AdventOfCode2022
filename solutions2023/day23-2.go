package solutions2023

import (
	"fmt"
)

func Day23p2(grid []string) {
	start := Point{0, 1}
	end := Point{len(grid) - 1, len(grid[0]) - 2}
	directions := Direction{
		north: Point{-1, 0},
		south: Point{1, 0},
		east:  Point{0, 1},
		west:  Point{0, -1},
	}

	best := maxDFS2(grid, start, directions, 0, end)

	fmt.Println("most scenic:", best)
}

var visited = make(map[Point]bool)
var maxSteps = 0

func maxDFS2(grid []string, pos Point, directions Direction, steps int, end Point) int {
	if visited[pos] {
		return -1
	} else if pos == end {
		return steps
	} else {
		visited[pos] = true
		n, s, e, w := addPoints(pos, directions.north), addPoints(pos, directions.south), addPoints(pos, directions.east), addPoints(pos, directions.west)

		best := 0
		for _, v := range []Point{n, s, e, w} {
			if inGrid(grid, v) && grid[v.y][v.x] != '#' {
				best = max(maxDFS2(grid, v, directions, steps+1, end), best)
			}
		}
		visited[pos] = false
		return best
	}
}
