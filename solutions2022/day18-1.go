package solutions2022

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func day18p1() {
	// Open the file.
	file, err := os.Open("/Users/cameron.bedard/Documents/FunRepos/AdventOfCode2022/inputs/day18.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file.
	scanner := bufio.NewScanner(file)

	arr := make([]cube, 0)

	// Read the file line by line.
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, ",")

		arr = append(arr, cube{atoi(nums[0]), atoi(nums[1]), atoi(nums[2]), 6})
	}

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if checkX(arr, i, j) || checkY(arr, i, j) || checkZ(arr, i, j) {
				arr[i].sides -= 1
				arr[j].sides -= 1
			}
		}
	}

	surfaceArea := 0
	for i := 0; i < len(arr); i++ {
		surfaceArea += arr[i].sides
	}

	fmt.Println("surface area:", surfaceArea)

	// Check for any errors that occurred while reading the file.
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
}

type cube struct {
	x     int
	y     int
	z     int
	sides int
}

func checkX(arr []cube, i int, j int) bool {
	if arr[i].y == arr[j].y && arr[i].z == arr[j].z {
		if arr[i].x == arr[j].x-1 || arr[i].x == arr[j].x+1 {
			return true
		}
	}
	return false
}

func checkY(arr []cube, i int, j int) bool {
	if arr[i].x == arr[j].x && arr[i].z == arr[j].z {
		if arr[i].y == arr[j].y-1 || arr[i].y == arr[j].y+1 {
			return true
		}
	}
	return false
}

func checkZ(arr []cube, i int, j int) bool {
	if arr[i].x == arr[j].x && arr[i].y == arr[j].y {
		if arr[i].z == arr[j].z-1 || arr[i].z == arr[j].z+1 {
			return true
		}
	}
	return false
}
