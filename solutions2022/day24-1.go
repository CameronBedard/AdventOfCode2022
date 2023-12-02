package solutions2022

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func day24p1() {
	// Open the file.
	file, err := os.Open("/Users/cameron.bedard/Documents/FunRepos/AdventOfCode2022/inputs2022/day24.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	field := make([]string, 0)
	start := Point{0, 0}
	end := Point{0, 0}
	blizzards := make([]Blizzard, 0)

	// Create a scanner to read the file.
	scanner := bufio.NewScanner(file)
	// Read the file line by line.
	for scanner.Scan() {
		field = append(field, scanner.Text())
	}

	//set start and end
	start.x = strings.Index(field[0], ".")
	end.y = len(field) - 1
	end.x = strings.Index(field[end.y], ".")

	for y, row := range field {
		for x, _ := range row {
			c := row[x]
			switch c {
			case '^':
				blizzards = append(blizzards, Blizzard{
					Point{y, x}, UP, Point{len(field) - 2, x},
				})
			case 'v':
				blizzards = append(blizzards, Blizzard{
					Point{y, x}, DOWN, Point{1, x},
				})
			case '<':
				blizzards = append(blizzards, Blizzard{
					Point{y, x}, LEFT, Point{y, len(field[0]) - 2},
				})
			case '>':
				blizzards = append(blizzards, Blizzard{
					Point{y, x}, RIGHT, Point{y, 1},
				})
			}
		}
	}

	fmt.Println("total:", blizzardBFS(field, blizzards, start, end))
}

func inBounds(field []string, pos Point) bool {
	return pos.x >= 0 && pos.x < len(field[0]) && pos.y >= 0 && pos.y < len(field)
}

func blizzardBFS(field []string, blizzards []Blizzard, start, target Point) int {
	minutes := 0
	currentStep := make(map[Point]bool)
	currentStep[start] = true

	for !currentStep[target] {
		whereBlizzards := make(map[Point]bool)
		for i, b := range blizzards {
			bb := addPoints(b.pos, b.dir)
			if inBounds(field, bb) {
				if field[bb.y][bb.x] == '#' {
					blizzards[i].pos = b.wrap
				} else {
					blizzards[i].pos = bb
				}
			}
			whereBlizzards[blizzards[i].pos] = true
		}

		newStep := make(map[Point]bool)
		for pos := range currentStep {
			if !(whereBlizzards[pos]) {
				newStep[pos] = true
			}
			for _, d := range directions {
				newPos := addPoints(pos, d)
				if inBounds(field, newPos) && field[newPos.y][newPos.x] != '#' && !whereBlizzards[newPos] {
					newStep[newPos] = true
				}
			}
		}
		currentStep = newStep
		minutes++
	}

	return minutes
}

type Blizzard struct {
	pos, dir, wrap Point
}

var (
	UP    = Point{-1, 0}
	DOWN  = Point{1, 0}
	LEFT  = Point{0, -1}
	RIGHT = Point{0, 1}
)

var directions = []Point{
	UP,
	DOWN,
	LEFT,
	RIGHT,
}
