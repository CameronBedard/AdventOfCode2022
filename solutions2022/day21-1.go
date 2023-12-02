package solutions2022

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func day21p1() {
	// Open the file.
	file, err := os.Open("/Users/cameron.bedard/Documents/FunRepos/AdventOfCode2022/inputs2022/day21.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	m := make(map[string]string)

	// Create a scanner to read the file.
	scanner := bufio.NewScanner(file)

	// Read the file line by line.
	for scanner.Scan() {
		line := scanner.Text()
		input := strings.Split(line, ": ")

		m[input[0]] = input[1]
	}
	start := time.Now()

	fmt.Println(monkeyMath(m, "root"))

	timeElapsed := time.Since(start)
	fmt.Println("This function took", timeElapsed)

	// Check for any errors that occurred while reading the file.
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
}

func monkeyMath(m map[string]string, name string) int {
	arr := strings.Split(m[name], " ")

	if len(arr) == 1 {
		return atoi(arr[0])
	} else if arr[1] == "+" {
		return monkeyMath(m, arr[0]) + monkeyMath(m, arr[2])
	} else if arr[1] == "-" {
		return monkeyMath(m, arr[0]) - monkeyMath(m, arr[2])
	} else if arr[1] == "*" {
		return monkeyMath(m, arr[0]) * monkeyMath(m, arr[2])
	} else if arr[1] == "/" {
		return monkeyMath(m, arr[0]) / monkeyMath(m, arr[2])
	}

	return -1000000000 //should not happen
}
