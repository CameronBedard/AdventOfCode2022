package solutions2023

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day2p1() {
	// Open the file.
	file, err := os.Open("/Users/cameron.bedard/Documents/FunRepos/AdventOfCode2022/inputs2023/day2.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file.
	scanner := bufio.NewScanner(file)

	// Read the file line by line.
	total := 0
	lineI := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineI++

		//Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
		turns := strings.Split(line, ":")[1]
		turnArr := strings.Split(turns, ";")

		ballsGood := true
		for turn := 0; turn < len(turnArr); turn++ {
			reveal := strings.Split(turnArr[turn], ",")

			for balls := 0; balls < len(reveal); balls++ {
				arr := strings.Split(reveal[balls], " ")

				number, _ := strconv.Atoi(arr[1])
				color := arr[2]

				if color == "red" && number > 12 {
					ballsGood = false
				}
				if color == "blue" && number > 14 {
					ballsGood = false
				}
				if color == "green" && number > 13 {
					ballsGood = false
				}
			}
		}
		if ballsGood {
			total += lineI
		}
	}
	//eof

	fmt.Println("total: ", total)

	// Check for any errors that occurred while reading the file.
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
}
