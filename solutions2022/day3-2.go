package solutions2022

import (
	"bufio"
	"fmt"
	"os"
)

func day3p2() {
	// Open the file.
	file, err := os.Open("/Users/cameron.bedard/Documents/FunRepos/AdventOfCode2022/inputs2022/day3.txt")
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
	groupI := 0
	arr := make([]int, 52)
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)

		for i := 0; i < len(line); i++ {
			if line[i] < 97 {
				if int(arr[line[i]-39]) == groupI {
					arr[line[i]-39] = groupI + 1
				}
				if int(arr[line[i]-39]) == 3 {
					prio += int(line[i] - 38)
					fmt.Printf("%v", arr)
					break
				}
			} else {
				if int(arr[line[i]-97]) == groupI {
					arr[line[i]-97] = groupI + 1
				}
				if int(arr[line[i]-97]) == 3 {
					prio += int(line[i] - 96)
					fmt.Printf("%v", arr)
					break
				}
			}
		}

		groupI++
		if groupI == 3 {
			groupI = 0
			arr = make([]int, 52)
		}
	}

	fmt.Println("prio: ", prio)

	// Check for any errors that occurred while reading the file.
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
}
