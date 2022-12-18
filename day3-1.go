package main

import (
	"bufio"
	"fmt"
	"os"
)

func day3p1() {
	// Open the file.
	file, err := os.Open("/Users/cameron.bedard/Documents/FunRepos/AdventOfCode2022/inputs/day3.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file.
	scanner := bufio.NewScanner(file)

	// Read the file line by line.
	// rucksack length always even, 2 compartments
	prio := 0
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		arr := make([]int, 52)

		for i := 0; i < len(line)/2; i++ {
			if line[i] < 97 {
				arr[line[i]-39]++
			} else {
				arr[line[i]-97]++
			}
		}

		for i := len(line) / 2; i < len(line); i++ {
			if line[i] < 97 {
				if arr[line[i]-39] > 0 {
					prio += int(line[i] - 38)
					break
				}
			} else {
				if arr[line[i]-97] > 0 {
					prio += int(line[i] - 96)
					break
				}
			}
		}
	}

	fmt.Println("prio: ", prio)

	// Check for any errors that occurred while reading the file.
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
}
