package solutions2023

import (
	"fmt"
	"strings"
)

func Day10p2(lines []string) {
	start := Point{}
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

	//we know from part1 we can run north for the loop, fence is length 13514
	fencePoints := runSim2(lines, start, directions.north, directions)

	//not cleaning the S point to its assumed char got my ass for an hour
	lines[start.y] = strings.ReplaceAll(lines[start.y], "S", "J")

	//get parity for each i, j
	parity := make([][]bool, len(lines))
	for i := range lines {
		parity[i] = make([]bool, len(lines[0]))
		for j := range lines[0] {
			parity[i][j] = verticalParity(lines, i, j, fencePoints)
		}
	}

	//run along the edges and DFS all the connected points into fencePoints
	area := 0
	//vertical edges
	for i := range lines {
		surroundedDFS(lines, Point{i, 0}, fencePoints, directions)
		surroundedDFS(lines, Point{i, len(lines[0]) - 1}, fencePoints, directions)
	}
	//horizontal edges
	for i := range lines[0] {
		surroundedDFS(lines, Point{0, i}, fencePoints, directions)
		surroundedDFS(lines, Point{0, len(lines) - 1}, fencePoints, directions)
	}

	for i := range lines {
		for j := range lines[0] {
			if !fencePoints[Point{i, j}] && parity[i][j] {
				area++
			}
		}
	}

	fmt.Println("area", area, "/", len(lines)*len(lines[0]))
}

func verticalParity(lines []string, row int, end int, fence map[Point]bool) bool {
	vertBars := 0
	last := uint8(0)
	for k := 0; k < end; k++ {
		if fence[Point{row, k}] {
			switch lines[row][k] {
			case '|':
				vertBars++
			case 'F':
				last = 'F'
			case 'J':
				if last == 'F' {
					vertBars++
					last = 0
				}
			case 'L':
				last = 'L'
			case '7':
				if last == 'L' {
					vertBars++
					last = 0
				}
			}
		}
	}

	if vertBars%2 == 1 {
		return true
	} else {
		return false
	}
}

func runSim2(lines []string, start Point, dir Point, directions Direction) map[Point]bool {
	loop := make(map[Point]bool)
	loop[start] = true
	nullDir := Point{0, 0}
	start = addPoints(start, dir)
	loopLength := 1

	for dir != nullDir {
		loop[start] = true
		if lines[start.y][start.x] == 'S' {
			fmt.Println("loop length:", loopLength)
			break
		}

		dir = getNextDir(lines, start, dir, directions)
		start = addPoints(start, dir)
		loopLength++
	}
	//fmt.Println("crash ended", tmpStart, loopLength)
	return loop
}

func surroundedDFS(grid []string, pos Point, fence map[Point]bool, directions Direction) {
	if !(pos.x >= 0 && pos.y >= 0 && pos.y < len(grid) && pos.x < len(grid[0])) {
		//nothing
	} else if fence[pos] {
		//nothing
	} else {
		fence[pos] = true
		n, s, e, w := addPoints(pos, directions.north), addPoints(pos, directions.south), addPoints(pos, directions.east), addPoints(pos, directions.west)

		if !fence[n] {
			surroundedDFS(grid, n, fence, directions)
		}
		if !fence[s] {
			surroundedDFS(grid, s, fence, directions)
		}
		if !fence[e] {
			surroundedDFS(grid, e, fence, directions)
		}
		if !fence[w] {
			surroundedDFS(grid, w, fence, directions)
		}
	}
}
