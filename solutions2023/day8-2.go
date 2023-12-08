package solutions2023

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Day8p2() {
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
	startNodes := make([]string, 0)

	//DRM = (DLQ, BGR)
	for i := 2; i < len(lines); i++ {
		lines[i] = strings.ReplaceAll(lines[i], " ", "")
		tmp := strings.Split(lines[i], "=")
		tmp[1] = strings.Trim(tmp[1], "()")

		nval := tmp[0]
		lr := strings.Split(tmp[1], ",")
		n := Node{nval, lr[0], lr[1]}
		nodes[nval] = &n

		if nval[2] == 'A' {
			startNodes = append(startNodes, nval)
		}
	}

	ans := make(map[int]int)
	for loop := 0; loop < 2000000000; loop++ {
		i := loop % len(directions)

		startNodes = getNext(startNodes, nodes, directions[i])

		for j := range startNodes {
			if startNodes[j][2] == 'Z' {
				ans[j] = loop + 1
			}
		}
		if len(ans) == len(startNodes) {
			break
		}
	}
	fmt.Println(ans)
	//answer is lcm of all cycles
}

func getNext(startNodes []string, nodes map[string]*Node, dir uint8) []string {
	for i := 0; i < len(startNodes); i++ {
		n := *nodes[startNodes[i]]
		if dir == 'L' {
			startNodes[i] = n.left
		} else {
			startNodes[i] = n.right
		}
	}
	return startNodes
}
