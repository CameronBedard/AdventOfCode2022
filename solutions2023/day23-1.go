package solutions2023

import (
	"fmt"
)

func Day23p1(grid []string) {
	start := Point{0, 1}
	directions := Direction{
		north: Point{-1, 0},
		south: Point{1, 0},
		east:  Point{0, 1},
		west:  Point{0, -1},
	}

	maxDFS(grid, start, directions, 0)

	best := 0
	for k, _ := range state {
		if k.p == (Point{len(grid) - 1, len(grid[0]) - 2}) {
			best = max(best, k.depth)
		}
	}
	fmt.Println("total", best)
}

func maxDFS(grid []string, pos Point, directions Direction, steps int) {
	if state[dfsState{pos, steps}] {
		//continue
	} else {
		state[dfsState{pos, steps}] = true
		n, s, e, w := addPoints(pos, directions.north), addPoints(pos, directions.south), addPoints(pos, directions.east), addPoints(pos, directions.west)

		if inGrid(grid, n) && grid[n.y][n.x] != '#' && grid[n.y][n.x] != 'v' && !state[dfsState{n, steps - 1}] {
			maxDFS(grid, n, directions, steps+1)
		}
		if inGrid(grid, s) && grid[s.y][s.x] != '#' && grid[s.y][s.x] != '^' && !state[dfsState{s, steps - 1}] {
			maxDFS(grid, s, directions, steps+1)
		}
		if inGrid(grid, e) && grid[e.y][e.x] != '#' && grid[e.y][e.x] != '<' && !state[dfsState{e, steps - 1}] {
			maxDFS(grid, e, directions, steps+1)
		}
		if inGrid(grid, w) && grid[w.y][w.x] != '#' && grid[w.y][w.x] != '>' && !state[dfsState{w, steps - 1}] {
			maxDFS(grid, w, directions, steps+1)
		}
	}
}
