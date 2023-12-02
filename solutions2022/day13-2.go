package solutions2022

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

func day13p2() {
	// Open the file.
	file, err := os.Open("/Users/cameron.bedard/Documents/FunRepos/AdventOfCode2022/inputs2022/day13.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	total := 1
	packets := make([]any, 0)

	// Create a scanner to read the file.
	scanner := bufio.NewScanner(file)
	// Read the file line by line.
	for scanner.Scan() {
		line1 := scanner.Text()
		scanner.Scan()
		line2 := scanner.Text()

		var a, b any
		json.Unmarshal([]byte(line1), &a)
		json.Unmarshal([]byte(line2), &b)
		packets = append(packets, a, b)

		scanner.Scan()
		scanner.Text() //blank line
	}

	var a, b any
	json.Unmarshal([]byte("[[2]]"), &a)
	json.Unmarshal([]byte("[[6]]"), &b)
	packets = append(packets, a, b)

	sort.Slice(packets, func(i, j int) bool {
		return cmp(packets[i], packets[j]) < 0
	})

	for i := 0; i < len(packets); i++ {
		if fmt.Sprint(packets[i]) == "[[2]]" || fmt.Sprint(packets[i]) == "[[6]]" {
			total *= i + 1
		}
	}

	fmt.Println("total:", total)

	// Check for any errors that occurred while reading the file.
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
}
