package solutions2023

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Day3p2() {
	// Open the file.
	file, err := os.Open("/Users/cameron.bedard/Documents/FunRepos/AdventOfCode2022/inputs2023/day3.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file.
	scanner := bufio.NewScanner(file)

	// Read the file line by line.
	grid := make([]string, 0)
	valid := make([][]Pos, 0)
	adjList := make(map[Pos][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)

		arr := make([]Pos, len(line))
		valid = append(valid, arr)
	}
	//eof

	//loop thru pos[i][j], if not . or 0-9 then its a symbol, mark all adjacent squares as valid
	//loop again and connect digits into a subStr if any char is on a valid[i][j] position, add it to sum
	n := len(grid)
	m := len(grid[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '*' {
				//check first row, left, right, last row: iff 2 of these have a digit then our gear is valid
				//upper row
				if inBounds(n, m, i-1, j) {
					valid[i-1][j] = Pos{i, j}

				}
				if inBounds(n, m, i-1, j-1) {
					valid[i-1][j-1] = Pos{i, j}
				}
				if inBounds(n, m, i-1, j+1) {
					valid[i-1][j+1] = Pos{i, j}
				}
				//L and R
				if inBounds(n, m, i, j-1) {
					valid[i][j-1] = Pos{i, j}
				}
				if inBounds(n, m, i, j+1) {
					valid[i][j+1] = Pos{i, j}
				}
				//Bottom row
				if inBounds(n, m, i+1, j) {
					valid[i+1][j] = Pos{i, j}
				}
				if inBounds(n, m, i+1, j-1) {
					valid[i+1][j-1] = Pos{i, j}
				}
				if inBounds(n, m, i+1, j+1) {
					valid[i+1][j+1] = Pos{i, j}
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		validNum := Pos{0, 0}
		currNum := ""
		for j := 0; j < m; j++ {
			if isNum(grid[i][j]) {
				currNum += string(grid[i][j])

				if valid[i][j] != (Pos{0, 0}) {
					validNum = valid[i][j]
				}
			} else {
				if validNum != (Pos{0, 0}) {
					x, _ := strconv.Atoi(currNum)
					adjList[validNum] = append(adjList[validNum], x)
				} //won't add to total if valid num is in last pos
				currNum = ""
				validNum = Pos{0, 0}
			}
		}

		//if valid num at end of line
		if validNum != (Pos{0, 0}) && len(currNum) > 0 {
			x, _ := strconv.Atoi(currNum)
			adjList[validNum] = append(adjList[validNum], x)
			currNum = ""
			validNum = Pos{0, 0}
		}
	}

	total := 0
	for _, nums := range adjList {
		if len(nums) == 2 {
			total += nums[0] * nums[1]
		}
	}

	fmt.Println("total: ", total)

	// Check for any errors that occurred while reading the file.
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
}

type Pos struct {
	y int
	x int
}
