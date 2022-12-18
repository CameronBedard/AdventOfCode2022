package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func day12p1() {
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

	start := Node{}
	dist := [41][67]int{}
	visited := [41][67]bool{}
	endNode := Node{}

	for i := 0; i < len(row); i++ {
		for j := 0; j < len(row[i]); j++ {
			if row[i][j] == 'S' {
				start = Node{'a', i, j}
				row[i] = strings.Replace(row[i], "S", "a", 1)
			}
			if row[i][j] == 'E' {
				endNode = Node{'z', i, j}
				row[i] = strings.Replace(row[i], "E", "z", 1)
			}
		}
	}

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
		if i == endNode.i && j == endNode.j {
			fmt.Println("final dist: ", currDist)
			break
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

	//fmt.Printf("%v", dist)

	// Check for any errors that occurred while reading the file.
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
}

type Node struct {
	val byte
	i   int
	j   int
}

func validIJ(i int, j int, height int, width int) bool {
	if i >= 0 && j >= 0 && i < height && j < width {
		return true
	}
	return false
}

// to use q:= NewQueue()
type NodeQueue []Node

func NewNodeQueue() *NodeQueue {
	return &NodeQueue{}
}

func (q *NodeQueue) Push(x Node) {
	*q = append(*q, x)
}

func (q *NodeQueue) IsEmpty() bool {
	h := *q
	l := len(h)
	return l == 0
}

func (q *NodeQueue) Pop() Node {
	h := *q
	el := h[0]
	*q = h[1:]

	return el
}
