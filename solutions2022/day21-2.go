package solutions2022

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func day21p2() {
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
	high := 4000000000000
	low := 0

	for low <= high {
		mid := low + ((high - low) / 2)
		m["humn"] = strconv.Itoa(mid)
		test := rootEquality(m, "root")
		if test < 0 {
			low = mid
		} else if test == 0 {
			fmt.Println(mid)
			break
		} else {
			high = mid
		}

		fmt.Println(test, mid)
	}

	m["humn"] = strconv.Itoa(3555057453229)
	fmt.Println(rootEquality(m, "root"))
	fmt.Println("another bug.. multiple valid answers (due to integer division)... it only accepts the lowest one")

	timeElapsed := time.Since(start)
	fmt.Println("This function took", timeElapsed)

	// Check for any errors that occurred while reading the file.
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
}

func rootEquality(m map[string]string, name string) int {
	input := strings.Split(m[name], " ")

	return monkeyMath(m, input[2]) - monkeyMath(m, input[0])
}
