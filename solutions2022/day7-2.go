package solutions2022

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func day7p2() {
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

	//have 70M, need 30M free space, total taken is 43.3M (rootSum)
	//select min directory for which rootSum - (directory) < 40M
	//search tree
	queue := make([]folder, 0)
	queue = append(queue, root)
	total := folderSumBFS(root)
	min := 70000000

	for len(queue) != 0 {
		curr := queue[0]
		queue = queue[1:]
		sum := folderSumBFS(curr)

		for i := 0; i < len(curr.children); i++ {
			queue = append(queue, *curr.children[i])
		}

		if total-sum < 40000000 {
			if sum < min {
				min = sum
			}
		}
	}

	fmt.Println("min directory: ", min)

	// Check for any errors that occurred while reading the file.
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
}
