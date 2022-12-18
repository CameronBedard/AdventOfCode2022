package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	low  int
	high int
}

func day4p2() {
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
	var arr []Pair
	pairs := 0
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		newL := strings.Split(line, ",")
		arr1 := strings.Split(newL[0], "-")
		arr2 := strings.Split(newL[1], "-")

		a1, err := strconv.Atoi(arr1[0])
		a2, err := strconv.Atoi(arr1[1])
		arr = append(arr, Pair{a1, a2})

		b1, err := strconv.Atoi(arr2[0])
		b2, err := strconv.Atoi(arr2[1])
		arr = append(arr, Pair{b1, b2})

		if err != nil {
			fmt.Println("error parsing ints")
		}

		pairs += 2
	}

	numPairs := 0
	for i := 0; i < pairs-1; i += 2 {
		if overlap(arr[i].low, arr[i].high, arr[i+1].low, arr[i+1].high) {
			numPairs++
		}
		fmt.Println(i, i+1)
	}

	fmt.Println("pairs: ", numPairs)

	// Check for any errors that occurred while reading the file.
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
}

func overlap(a1 int, a2 int, b1 int, b2 int) bool {
	if a1 < b1 && a2 < b1 {
		return false
	}
	if b1 < a1 && b2 < a1 {
		return false
	}

	return true
}
