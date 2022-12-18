package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func day4p1() {
	// Open the file.
	file, err := os.Open("/Users/cameron.bedard/Documents/FunRepos/AdventOfCode2022/inputs/day4.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file.
	scanner := bufio.NewScanner(file)

	// Read the file line by line.
	// rucksack length always even, 2 compartments
	pairs := 0
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		newL := strings.Split(line, ",")
		arr1 := strings.Split(newL[0], "-")
		arr2 := strings.Split(newL[1], "-")

		a1, err := strconv.Atoi(arr1[0])
		a2, err := strconv.Atoi(arr1[1])

		b1, err := strconv.Atoi(arr2[0])
		b2, err := strconv.Atoi(arr2[1])

		if err != nil {
			fmt.Println("error parsing ints")
		}

		if a1 >= b1 && a2 <= b2 {
			pairs++
		} else if b1 >= a1 && b2 <= a2 {
			pairs++
		}
	}

	fmt.Println("pairs: ", pairs)

	// Check for any errors that occurred while reading the file.
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
}
