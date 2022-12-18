package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func day1p1() {
	// Open the file.
	file, err := os.Open("/Users/cameron.bedard/Documents/FunRepos/AdventOfCode2022/inputs/day1.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file.
	scanner := bufio.NewScanner(file)

	// Read the file line by line.
	max := 0
	current := 0
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		if line == "" {
			current = 0
		} else {
			myInt, err := strconv.Atoi(line)

			// Check for errors
			if err != nil {
				// Handle the error
			} else {
				current += myInt
				if current > max {
					max = current
				}
			}
		}
	}

	fmt.Println("max: ", max)

	// Check for any errors that occurred while reading the file.
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
}
