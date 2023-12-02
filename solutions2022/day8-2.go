package solutions2022

import (
	"bufio"
	"fmt"
	"os"
)

func day8p2() {
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

	max := 0
	for i := 0; i < 99; i++ {
		for j := 0; j < 99; j++ {
			scenicScore := checkUp2(i, j, arr) * checkDown2(i, j, arr) * checkLeft2(i, j, arr) * checkRight2(i, j, arr)

			if scenicScore > max {
				max = scenicScore
			}
		}
	}
	fmt.Println("scenic score:", max)

	// Check for any errors that occurred while reading the file.
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
}

func checkUp2(i int, j int, arr [99][99]int) int {
	height := arr[i][j]
	count := 0
	i--
	for i >= 0 {
		count++
		if arr[i][j] >= height {
			break
		}
		i--
	}
	return count
}

func checkDown2(i int, j int, arr [99][99]int) int {
	height := arr[i][j]
	count := 0
	i++
	for i < 99 {
		count++
		if arr[i][j] >= height {
			break
		}
		i++
	}
	return count
}

func checkLeft2(i int, j int, arr [99][99]int) int {
	height := arr[i][j]
	count := 0
	j--
	for j >= 0 {
		count++
		if arr[i][j] >= height {
			break
		}
		j--
	}
	return count
}

func checkRight2(i int, j int, arr [99][99]int) int {
	height := arr[i][j]
	count := 0
	j++
	for j < 99 {
		count++
		if arr[i][j] >= height {
			break
		}
		j++
	}
	return count
}
