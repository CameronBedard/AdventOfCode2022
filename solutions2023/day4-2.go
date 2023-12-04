package solutions2023

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day4p2() {
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
	lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	//eof

	copies := make([]int, len(lines))
	for i := range copies {
		copies[i] = 1
	}

	for i := range lines {
		tmp := strings.Split(lines[i], ":")
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
					matches++
				}
			}
		}
		//fmt.Println("card ", i+1, " has ", matches, " matches.")

		//for the next $matches cards, take copy[i] and add that to their copy[i]
		for copyI := i + 1; copyI <= i+matches; copyI++ {
			copies[copyI] += copies[i]
		}
		//fmt.Println(copies)
	}

	//sum copies arr
	sum := 0
	for i := range copies {
		sum += copies[i]
	}
	fmt.Println("total: ", sum)

	// Check for any errors that occurred while reading the file.
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
}
