package solutions2023

import (
	"fmt"
)

func Day16p2(lines []string) {
	ans := 0
	for i := range lines {
		start := Point{i, 0}
		dir := Point{0, 1}
		ans = max(ans, energizedSim(lines, start, dir))

		start = Point{i, len(lines[0]) - 1}
		dir = Point{0, -1}
		ans = max(ans, energizedSim(lines, start, dir))
	}

	for i := range lines[0] {
		start := Point{0, i}
		dir := Point{1, 0}
		ans = max(ans, energizedSim(lines, start, dir))

		start = Point{len(lines) - 1, i}
		dir = Point{-1, 0}
		ans = max(ans, energizedSim(lines, start, dir))
	}

	fmt.Println(ans)
}

func energizedSim(grid []string, start Point, dir Point) int {
	directions := Direction{
		north: Point{-1, 0},
		south: Point{1, 0},
		east:  Point{0, 1},
		west:  Point{0, -1},
	}
	visited := make(map[Point]bool)
	queue := make([][]Point, 1)
	queue[0] = []Point{start, dir}
	loopCheck := make(map[hist]bool)

	for len(queue) > 0 {
		next := nextDir(grid, queue[0][0], queue[0][1], directions, visited)

		if len(next) == 0 {
			if len(queue) == 1 {
				queue = nil
				break
			} else {
				queue = queue[1:]
			}
		} else if len(next) == 2 {
			//add new dir to the slice
			queue = append(queue, []Point{addPoints(next[1], queue[0][0]), next[1]})
			//get next point
			queue[0][0] = addPoints(next[0], queue[0][0])
			queue[0][1] = next[0]
		} else {
			queue[0][0] = addPoints(next[0], queue[0][0])
			queue[0][1] = next[0]
		}

		curr := hist{queue[0][0], queue[0][1]}
		if loopCheck[curr] == true {
			if len(queue) == 1 {
				queue = nil
			} else {
				queue = queue[1:]
			}
		} else {
			loopCheck[curr] = true
		}
	}

	return len(visited)
}
