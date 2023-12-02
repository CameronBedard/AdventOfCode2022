package solutions2022

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func day1p2() {
	// Open the file.
	file, err := os.Open("/Users/cameron.bedard/Documents/FunRepos/AdventOfCode2022/inputs/day1.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file.
	scanner := bufio.NewScanner(file)

	var arr []int // an empty list
	// Read the file line by line.
	current := 0
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		if line == "" {
			arr = append(arr, current)

			current = 0
		} else {
			myInt, err := strconv.Atoi(line)

			// Check for errors
			if err != nil {
				// Handle the error
			} else {
				current += myInt
			}
		}
	}
	arr = append(arr, current)
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))

	fmt.Println("top 3: ", arr[0]+arr[1]+arr[2])

	// Check for any errors that occurred while reading the file.
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
}
