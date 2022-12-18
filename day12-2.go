package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func day12p2() {
	// Open the file.
	file, err := os.Open("/Users/cameron.bedard/Documents/FunRepos/AdventOfCode2022/inputs/day12.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file.
	scanner := bufio.NewScanner(file)
	row := make([]string, 0)

	// Read the file line by line.
	for scanner.Scan() {
		line := scanner.Text()
		row = append(row, line)
	}

	endNode := Node{}
	minDist := 5000

	for i := 0; i < len(row); i++ {
		for j := 0; j < len(row[i]); j++ {
			if row[i][j] == 'E' {
				endNode = Node{'z', i, j}
				row[i] = strings.Replace(row[i], "E", "z", 1)
			}
		}
	}

	for i := 0; i < len(row); i++ {
		for j := 0; j < len(row[i]); j++ {
			if row[i][j] == 'a' {
				distance := BFS(Node{'a', i, j}, endNode, row)

				if distance > 0 && distance < minDist {
					minDist = distance
				}
			}
		}
	}

	fmt.Println("best distance: ", minDist)

	// Check for any errors that occurred while reading the file.
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
}

func BFS(start Node, end Node, row []string) int {
	dist := [41][67]int{}
	visited := [41][67]bool{}

	queue := NewNodeQueue()
	queue.Push(start)
	dist[start.i][start.j] = 0
	visited[start.i][start.j] = true

	for !queue.IsEmpty() {
		curr := queue.Pop()
		i := curr.i
		j := curr.j
		currDist := dist[i][j]

		//check for end node
		if i == end.i && j == end.j {
			return currDist
		}

		if validIJ(i-1, j, len(row), len(row[0])) {
			if int(row[i-1][j])-int(row[i][j]) <= 1 && !(visited[i-1][j]) {
				queue.Push(Node{row[i-1][j], i - 1, j})
				dist[i-1][j] = currDist + 1
				visited[i-1][j] = true
			}
		}
		if validIJ(i+1, j, len(row), len(row[0])) {
			if int(row[i+1][j])-int(row[i][j]) <= 1 && !visited[i+1][j] {
				queue.Push(Node{row[i+1][j], i + 1, j})
				dist[i+1][j] = currDist + 1
				visited[i+1][j] = true
			}
		}
		if validIJ(i, j-1, len(row), len(row[0])) {
			if int(row[i][j-1])-int(row[i][j]) <= 1 && !visited[i][j-1] {
				queue.Push(Node{row[i][j-1], i, j - 1})
				dist[i][j-1] = currDist + 1
				visited[i][j-1] = true
			}
		}
		if validIJ(i, j+1, len(row), len(row[0])) {
			if int(row[i][j+1])-int(row[i][j]) <= 1 && !(visited[i][j+1]) {
				queue.Push(Node{row[i][j+1], i, j + 1})
				dist[i][j+1] = currDist + 1
				visited[i][j+1] = true
			}
		}
	}
	//end node not reachable
	return -1
}
