package solutions2023

import (
	"bufio"
	"fmt"
	"os"
)

func Day1p1() {
	// Open the file.
	file, err := os.Open("/Users/cameron.bedard/Documents/FunRepos/AdventOfCode2022/inputs2023/day1.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file.
	scanner := bufio.NewScanner(file)

	// Read the file line by line.
	total := 0
	for scanner.Scan() {
		line := scanner.Text()

		//first num
		for i := 0; i < len(line); i++ {
			if line[i] >= 48 && line[i] <= 57 {
				x := line[i] - 48
				total += int(x) * 10
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if line[i] >= 48 && line[i] <= 57 {
				x := line[i] - 48
				total += int(x)
				break
			}
		}
	}
	//eof

	fmt.Println("total: ", total)

	// Check for any errors that occurred while reading the file.
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
}
