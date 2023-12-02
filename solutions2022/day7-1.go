package solutions2022

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func day7p1() {
	// Open the file.
	file, err := os.Open("/Users/cameron.bedard/Documents/FunRepos/AdventOfCode2022/inputs/day7.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	root := folder{"root", nil, make([]*folder, 0), 0}
	root.children = append(root.children, &folder{"/", &root, make([]*folder, 0), 0})
	current := &root

	// Create a scanner to read the file.
	scanner := bufio.NewScanner(file)

	// Read the file line by line.
	for scanner.Scan() {
		line := scanner.Text()
		command := strings.Split(line, " ")
		if command[0] == "$" {
			if command[1] == "cd" {
				if command[2] == ".." {
					if current.parent != nil {
						current = current.parent
					}
				} else {
					for i := 0; i < len(current.children); i++ {
						if command[2] == current.children[i].name {
							current = current.children[i]
						}
					}
				}
			} else {
				//its ls, continue
			}
		} else {
			//its listing files or folders for current
			if command[0] == "dir" {
				current.children = append(current.children, &folder{command[1], current, make([]*folder, 0), 0})
			} else {
				current.size += atoi(command[0])
			}
		}
	}

	//search tree
	queue := make([]folder, 0)
	queue = append(queue, root)
	total := 0

	for len(queue) != 0 {
		curr := queue[0]
		queue = queue[1:]
		sum := folderSumBFS(curr)

		for i := 0; i < len(curr.children); i++ {
			queue = append(queue, *curr.children[i])
		}

		if sum <= 100000 {
			total += sum
		}
	}

	fmt.Println("total: ", total)

	// Check for any errors that occurred while reading the file.
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
}

type folder struct {
	name     string
	parent   *folder
	children []*folder
	size     int
}

func folderSumBFS(root folder) int {
	queue := make([]folder, 0)
	queue = append(queue, root)
	sum := 0

	for len(queue) != 0 {
		curr := queue[0]
		queue = queue[1:]
		sum += curr.size

		for i := 0; i < len(curr.children); i++ {
			queue = append(queue, *curr.children[i])
		}
	}

	return sum
}
