package solutions2022

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func day2p2() {
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
			if round[0] == "A" {
				score += 3
			}
			if round[0] == "B" {
				score += 1
			}
			if round[0] == "C" {
				score += 2
			}
		}
		if round[1] == "Y" {
			if round[0] == "A" {
				score += 4
			}
			if round[0] == "B" {
				score += 5
			}
			if round[0] == "C" {
				score += 6
			}
		}
		if round[1] == "Z" {
			if round[0] == "A" {
				score += 8
			}
			if round[0] == "B" {
				score += 9
			}
			if round[0] == "C" {
				score += 7
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
