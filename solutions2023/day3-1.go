package solutions2023

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Day3p1() {
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
	valid := make([][]bool, 0)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)

		arr := make([]bool, len(line))
		valid = append(valid, arr)
	}
	//eof

	//loop thru pos[i][j], if not . or 0-9 then its a symbol, mark all adjacent squares as valid
	//loop again and connect digits into a subStr if any char is on a valid[i][j] position, add it to sum
	n := len(grid)
	m := len(grid[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if isSymbol(grid[i][j]) {
				//upper row
				if inBounds(n, m, i-1, j) {
					valid[i-1][j] = true
				}
				if inBounds(n, m, i-1, j-1) {
					valid[i-1][j-1] = true
				}
				if inBounds(n, m, i-1, j+1) {
					valid[i-1][j+1] = true
				}
				//L and R
				if inBounds(n, m, i, j-1) {
					valid[i][j-1] = true
				}
				if inBounds(n, m, i, j+1) {
					valid[i][j+1] = true
				}
				//Bottom row
				if inBounds(n, m, i+1, j) {
					valid[i+1][j] = true
				}
				if inBounds(n, m, i+1, j-1) {
					valid[i+1][j-1] = true
				}
				if inBounds(n, m, i+1, j+1) {
					valid[i+1][j+1] = true
				}
			}
		}
	}

	total := 0
	for i := 0; i < n; i++ {

		validNum := false
		currNum := ""
		for j := 0; j < m; j++ {
			if isNum(grid[i][j]) {
				currNum += string(grid[i][j])
				if valid[i][j] {
					validNum = true
				}
			} else {
				if validNum {
					x, _ := strconv.Atoi(currNum)
					total += x
				} //wont add to total if valid num is in last pos
				currNum = ""
				validNum = false
			}
		}

		//if valid num at end of line
		if validNum && len(currNum) > 0 {
			x, _ := strconv.Atoi(currNum)
			total += x
			currNum = ""
			validNum = false
		}
	}

	fmt.Println("total: ", total)

	// Check for any errors that occurred while reading the file.
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
}

func inBounds(n, m, i, j int) bool {
	return i >= 0 && j >= 0 && i < n && j < m
}

func isSymbol(x uint8) bool {
	if x == '.' || (x >= 48 && x <= 57) {
		return false
	}
	return true
}

func isNum(x uint8) bool {
	if x >= 48 && x <= 57 {
		return true
	}
	return false
}
