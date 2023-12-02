package solutions2023

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Day1p2() {
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
		for i := 0; i < len(line)-1; i++ {
			x := isDigit(line[i])
			if x > 0 {
				total += x * 10
				break
			}
			x = isDigit(line[i+1])
			if x > 0 {
				total += x * 10
				break
			}

			if len(line)-i >= 5 {
				x = isWordNumber(line[i : i+5])
			} else {
				x = isWordNumber(line[i:])
			}
			if x > 0 {
				total += x * 10
				break
			}
		}

		//second num
		for i := len(line) - 1; i >= 0; i-- {
			x := isDigit(line[i])
			if x > 0 {
				total += x
				break
			}

			x = isWordNumber(line[i:])
			if x > 0 {
				total += x
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

func isDigit(x uint8) int {
	if x >= 49 && x <= 57 {
		x -= 48
		return int(x)
	}
	return 0
}

func isWordNumber(str string) int {
	//largest len is five
	numbers := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}

	for i := 0; i < len(numbers); i++ {
		if strings.Contains(str, numbers[i]) {
			return i + 1
		}
	}

	return 0
}
