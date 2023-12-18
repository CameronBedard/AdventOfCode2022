package solutions2023

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var count = 0

func Day18p1(lines []string) {
	grid := make([][]uint8, 400)
	for i := range grid {
		grid[i] = make([]uint8, 700)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}

	direction := map[string]Point{
		"D": south,
		"U": north,
		"L": west,
		"R": east,
	}

	pos := Point{300, 300}
	for i := range lines {
		inputs := strings.Split(lines[i], " ")
		dir := direction[inputs[0]]
		dist, _ := strconv.Atoi(inputs[1])

		for j := 0; j < dist; j++ {
			pos = addPoints(pos, dir)
			grid[pos.y][pos.x] = '#'
			count++
		}
	}

	//writeToFile(grid)
	DFS(grid, 225, 165)

	fmt.Println("total", count)
}

func DFS(grid [][]uint8, i, j int) {
	grid[i][j] = '#'
	count++

	if grid[i+1][j] == '.' {
		DFS(grid, i+1, j)
	}
	if grid[i-1][j] == '.' {
		DFS(grid, i-1, j)
	}
	if grid[i][j+1] == '.' {
		DFS(grid, i, j+1)
	}
	if grid[i][j-1] == '.' {
		DFS(grid, i, j-1)
	}
}

func writeToFile(grid [][]uint8) {
	file, err := os.OpenFile("inputs2023/output.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("File does not exists or cannot be created")
		os.Exit(1)
	}
	defer file.Close()

	w := bufio.NewWriter(file)

	for i := range grid {
		for j := range grid[i] {
			fmt.Fprint(w, string(grid[i][j]))
		}
		fmt.Fprintln(w)
	}

	w.Flush()
}
