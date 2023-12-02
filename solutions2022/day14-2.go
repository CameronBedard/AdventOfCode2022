package solutions2022

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func day14p2() {
	// Open the file.
	file, err := os.Open("/Users/cameron.bedard/Documents/FunRepos/AdventOfCode2022/inputs/day14.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file.
	scanner := bufio.NewScanner(file)
	cave := make([][]int, 700)
	maxY := 0

	for i := 0; i < 700; i++ {
		cave[i] = make([]int, 1000)
	}

	// Read the file line by line.
	for scanner.Scan() {
		line := scanner.Text()
		coordinates := strings.Split(line, " -> ")
		list := make([]Point, 0)

		for i := 0; i < len(coordinates); i++ {
			xy := strings.Split(coordinates[i], ",")

			list = append(list, Point{atoi(xy[1]), atoi(xy[0])})

			if atoi(xy[1]) > maxY {
				maxY = atoi(xy[1])
			}
		}

		for i := 1; i < len(list); i++ {
			drawLine(list[i], list[i-1], cave)
		}
	}

	fmt.Println("maxY:", maxY)
	drawLine(Point{maxY + 2, 0}, Point{maxY + 2, 999}, cave)

	sandCount := 0
	for dropSand2(0, 500, cave) {
		sandCount++
	}

	fmt.Println("drops of sand: ", sandCount)

	// Check for any errors that occurred while reading the file.
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
}

func dropSand2(y int, x int, cave [][]int) bool {
	if cave[0][500] == 1 {
		return false
	}
	if cave[y+1][x] == 0 {
		return dropSand(y+1, x, cave)
	} else if cave[y+1][x-1] == 0 {
		return dropSand(y+1, x-1, cave)
	} else if cave[y+1][x+1] == 0 {
		return dropSand(y+1, x+1, cave)
	}

	cave[y][x] = 1
	return true
}
