package solutions2023

import (
	"fmt"
)

func Day21p1(lines []string) {
	start := Point{0, 0}
	for i := range lines {
		for j := range lines[0] {
			if lines[i][j] == 'S' {
				start = Point{i, j}
			}
		}
	}

	points64 := make(map[Point]bool)
	directions := Direction{
		north: Point{-1, 0},
		south: Point{1, 0},
		east:  Point{0, 1},
		west:  Point{0, -1},
	}

	depthDFS(lines, start, points64, directions, 0)

	fmt.Println("total", len(points64))
}

var state = make(map[dfsState]bool)

func depthDFS(grid []string, pos Point, points64 map[Point]bool, directions Direction, steps int) {
	if state[dfsState{pos, steps}] {
		//continue
	} else if steps == 64 {
		points64[pos] = true
	} else {
		state[dfsState{pos, steps}] = true
		n, s, e, w := addPoints(pos, directions.north), addPoints(pos, directions.south), addPoints(pos, directions.east), addPoints(pos, directions.west)

		if inGrid(grid, n) && grid[n.y][n.x] != '#' {
			depthDFS(grid, n, points64, directions, steps+1)
		}
		if inGrid(grid, s) && grid[s.y][s.x] != '#' {
			depthDFS(grid, s, points64, directions, steps+1)
		}
		if inGrid(grid, e) && grid[e.y][e.x] != '#' {
			depthDFS(grid, e, points64, directions, steps+1)
		}
		if inGrid(grid, w) && grid[w.y][w.x] != '#' {
			depthDFS(grid, w, points64, directions, steps+1)
		}
	}
}

type dfsState struct {
	p     Point
	depth int
}

func inGrid(grid []string, pos Point) bool {
	return pos.x >= 0 && pos.y >= 0 && pos.y < len(grid) && pos.x < len(grid[0])
}
