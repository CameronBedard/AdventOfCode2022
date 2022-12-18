package main

import (
	"bufio"
	"fmt"
	"os"
)

func day8p1() {
	// Open the file.
	file, err := os.Open("/Users/cameron.bedard/Documents/FunRepos/AdventOfCode2022/inputs/day8.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file.
	scanner := bufio.NewScanner(file)
	arr := [99][99]int{}
	lineN := 0

	// Read the file line by line.
	for scanner.Scan() {
		line := scanner.Text()

		for i := 0; i < len(line); i++ {
			arr[lineN][i] = int(line[i] - 48)
		}

		lineN++
	}

	total := 0
	for i := 0; i < 99; i++ {
		for j := 0; j < 99; j++ {
			if checkUp(i, j, arr) || checkDown(i, j, arr) || checkLeft(i, j, arr) || checkRight(i, j, arr) {
				total++
			}
		}
	}
	fmt.Println("visible trees:", total)

	// Check for any errors that occurred while reading the file.
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
}

func checkUp(i int, j int, arr [99][99]int) bool {
	height := arr[i][j]
	i--
	for i >= 0 {
		if arr[i][j] >= height {
			return false
		}
		i--
	}
	return true
}

func checkDown(i int, j int, arr [99][99]int) bool {
	height := arr[i][j]
	i++
	for i < 99 {
		if arr[i][j] >= height {
			return false
		}
		i++
	}
	return true
}

func checkLeft(i int, j int, arr [99][99]int) bool {
	height := arr[i][j]
	j--
	for j >= 0 {
		if arr[i][j] >= height {
			return false
		}
		j--
	}
	return true
}

func checkRight(i int, j int, arr [99][99]int) bool {
	height := arr[i][j]
	j++
	for j < 99 {
		if arr[i][j] >= height {
			return false
		}
		j++
	}
	return true
}
