package solutions2023

import (
	"fmt"
)

func Day16p1(lines []string) {
	start := Point{0, 0}
	dir := Point{0, 1}
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
		next := nextDir(lines, queue[0][0], queue[0][1], directions, visited)

		if len(next) == 0 {
			if len(queue) == 1 {
				queue = nil
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

	fmt.Println(len(visited))
}

type hist struct {
	pos Point
	dir Point
}

func nextDir(grid []string, pos Point, dir Point, directions Direction, visited map[Point]bool) []Point {
	arr := make([]Point, 0)
	//check inbounds
	if !(pos.x >= 0 && pos.y >= 0 && pos.y < len(grid) && pos.x < len(grid[0])) {
		return arr
	} else {
		visited[pos] = true
		switch grid[pos.y][pos.x] {
		case '-':
			if dir == directions.north || dir == directions.south {
				arr = append(arr, directions.west, directions.east)
				return arr
			}
		case '|':
			if dir == directions.east || dir == directions.west {
				arr = append(arr, directions.north, directions.south)
				return arr
			}
		case '\\':
			p := Point{dir.x, dir.y}
			arr = append(arr, p)
			return arr
		case '/':
			p := Point{-(dir.x), -(dir.y)}
			arr = append(arr, p)
			return arr
		default:
			//nothing
		}
	}
	arr = append(arr, dir)
	return arr
}
