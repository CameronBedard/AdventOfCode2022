package solutions2023

import (
	"fmt"
)

func Day10p1(lines []string) {
	start := Point{}
	for i := range lines {
		for j := range lines[0] {
			if lines[i][j] == 'S' {
				start = Point{i, j}
			}
		}
	}
	//from S, try pathing across all 4 neighbours, print all loops
	//not allowed - to | or | to -, need corner pipe
	//each piece should have piece(symbol, currDir, nextDir) where currDir is ex. {1,0} (north)
	// and nextDir is the transform on the dir or it returns {0,0}
	directions := Direction{
		north: Point{-1, 0},
		south: Point{1, 0},
		east:  Point{0, 1},
		west:  Point{0, -1},
	}

	runSim(lines, start, directions.north, directions)
	runSim(lines, start, directions.south, directions)
	runSim(lines, start, directions.east, directions)
	runSim(lines, start, directions.west, directions)
	//answer is loopLength / 2
}

func runSim(lines []string, start Point, dir Point, directions Direction) {
	nullDir := Point{0, 0}
	tmpStart := addPoints(start, dir)
	loopLength := 1

	for dir != nullDir {
		if lines[tmpStart.y][tmpStart.x] == 'S' {
			fmt.Println("loop length:", loopLength)
			break
		}

		dir = getNextDir(lines, tmpStart, dir, directions)
		tmpStart = addPoints(tmpStart, dir)
		loopLength++
	}
	//fmt.Println("crash ended", tmpStart, loopLength)
}

func getNextDir(grid []string, pos Point, dir Point, directions Direction) Point {
	//check inbounds
	nullPoint := Point{0, 0}

	if !(pos.x >= 0 && pos.y >= 0 && pos.y < len(grid) && pos.x < len(grid[0])) {
		return nullPoint
	} else {
		switch grid[pos.y][pos.x] {
		case '-':
			if dir == directions.west || dir == directions.east {
				return dir
			}
		case '|':
			if dir == directions.north || dir == directions.south {
				return dir
			}
		case 'L':
			if dir == directions.west {
				return directions.north
			} else if dir == directions.south {
				return directions.east
			}
		case 'J':
			if dir == directions.east {
				return directions.north
			} else if dir == directions.south {
				return directions.west
			}
		case '7':
			if dir == directions.east {
				return directions.south
			} else if dir == directions.north {
				return directions.west
			}
		case 'F':
			if dir == directions.north {
				return directions.east
			} else if dir == directions.west {
				return directions.south
			}
		default:
			return nullPoint
		}
	}
	return nullPoint
}

type Point struct {
	y int
	x int
}

type Direction struct {
	north Point
	south Point
	east  Point
	west  Point
}

func addPoints(p1, p2 Point) Point {
	return Point{p1.y + p2.y, p1.x + p2.x}
}
