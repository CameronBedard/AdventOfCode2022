package solutions2022

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func day2p1() {
	// Open the file.
	file, err := os.Open("/Users/cameron.bedard/Documents/FunRepos/AdventOfCode2022/inputs2022/day2.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file.
	scanner := bufio.NewScanner(file)

	// Read the file line by line.
	score := 0
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		round := strings.Split(line, " ")

		if round[1] == "X" {
			score += 1
			if round[0] == "A" {
				score += 3
			}
			if round[0] == "B" {
				score += 0
			}
			if round[0] == "C" {
				score += 6
			}
		}
		if round[1] == "Y" {
			score += 2
			if round[0] == "A" {
				score += 6
			}
			if round[0] == "B" {
				score += 3
			}
			if round[0] == "C" {
				score += 0
			}
		}
		if round[1] == "Z" {
			score += 3
			if round[0] == "A" {
				score += 0
			}
			if round[0] == "B" {
				score += 6
			}
			if round[0] == "C" {
				score += 3
			}
		}
	}

	fmt.Println("score: ", score)

	// Check for any errors that occurred while reading the file.
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
}
