package solutions2023

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day4p1() {
	// Open the file.
	file, err := os.Open("/Users/cameron.bedard/Documents/FunRepos/AdventOfCode2022/inputs2023/day4.txt")
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

		tmp := strings.Split(line, ":")
		tmp2 := strings.Split(tmp[1], "|")
		//tmp2[0] winning nums, tmp2[1] our nums
		tmp2[0] = strings.Trim(tmp2[0], " ")
		tmp2[1] = strings.Trim(tmp2[1], " ")
		winNums := strings.Split(tmp2[0], " ")
		ourNums := strings.Split(tmp2[1], " ")

		winMap := make(map[int]bool, 0)
		for _, v := range winNums {
			x, _ := strconv.Atoi(v)
			winMap[x] = true
		}

		matches := 0
		for _, v := range ourNums {
			x, err := strconv.Atoi(v)
			if err == nil {
				if winMap[x] == true {
					if matches == 0 {
						matches = 1
					} else {
						matches *= 2
					}
				}
			}
		}

		total += matches
	}
	//eof

	fmt.Println("total: ", total)

	// Check for any errors that occurred while reading the file.
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
}
