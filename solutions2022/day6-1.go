package solutions2022

import (
	"bufio"
	"fmt"
	"os"
)

func day6p1() {
	// Open the file.
	file, err := os.Open("/Users/cameron.bedard/Documents/FunRepos/AdventOfCode2022/inputs/day6.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file.
	scanner := bufio.NewScanner(file)

	// Read the file line by line.
	for scanner.Scan() {
		line := scanner.Text()

		for i := 3; i < len(line); i++ {
			if valid(line[i-3 : i+1]) {
				fmt.Println("marker: ", i+1)
				break
			}
		}
	}

	// Check for any errors that occurred while reading the file.
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
}

func valid(s string) bool {
	arr := make([]int, 26)
	fmt.Println("slice: ", s)
	for i := 0; i < len(s); i++ {
		arr[s[i]-97]++
	}

	for i := 0; i < 26; i++ {
		if arr[i] > 1 {
			return false
		}
	}

	return true
}
