package solutions2023

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Day8p1() {
	// Open the file.
	file, err := os.Open("/Users/cameron.bedard/Documents/FunRepos/AdventOfCode2022/inputs2023/day8.txt")
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

	directions := lines[0]
	nodes := make(map[string]*Node)

	//DRM = (DLQ, BGR)
	for i := 2; i < len(lines); i++ {
		lines[i] = strings.ReplaceAll(lines[i], " ", "")
		tmp := strings.Split(lines[i], "=")
		tmp[1] = strings.Trim(tmp[1], "()")

		nval := tmp[0]
		lr := strings.Split(tmp[1], ",")
		n := Node{nval, lr[0], lr[1]}
		nodes[nval] = &n
	}

	n := *nodes["AAA"]
	for loop := 0; loop < 1000000000; loop++ {
		i := loop % len(directions)

		if directions[i] == 'L' {
			n = *nodes[n.left]
		} else {
			n = *nodes[n.right]
		}
		if n.val == "ZZZ" {
			fmt.Println(loop + 1)
			break
		}
	}
}

type Node struct {
	val   string
	left  string
	right string
}
