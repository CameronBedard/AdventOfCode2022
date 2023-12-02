package solutions2022

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func day10p1() {
	// Open the file.
	file, err := os.Open("/Users/cameron.bedard/Documents/FunRepos/AdventOfCode2022/inputs2022/day10.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file.
	scanner := bufio.NewScanner(file)
	signal := 0
	x := 1
	cycle := 0

	// Read the file line by line.
	for scanner.Scan() {
		line := scanner.Text()

		instr := strings.Split(line, " ")

		if len(instr) == 1 {
			//noop
			updateCycle(&cycle, &signal, x)
		} else {
			//addx
			updateCycle(&cycle, &signal, x)
			updateCycle(&cycle, &signal, x)
			x += atoi(instr[1])
		}
	}

	fmt.Println("signal strength:", signal)

	// Check for any errors that occurred while reading the file.
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
}

func updateCycle(cycle *int, signal *int, x int) {
	*cycle++
	if (*cycle-20)%40 == 0 && *cycle <= 220 {
		*signal += (x * (*cycle))
		fmt.Println(*cycle, *signal)
	}
}
