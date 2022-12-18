package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func day18p2() {
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
	minX := 100000
	minY := 100000
	minZ := 100000
	maxX := -100000
	maxY := -100000
	maxZ := -100000

	// Read the file line by line.
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, ",")

		arr = append(arr, cube{atoi(nums[0]), atoi(nums[1]), atoi(nums[2]), 6})

		//getting 3d bounds
		n := len(arr) - 1
		if arr[n].x > maxX {
			maxX = arr[n].x
		}
		if arr[n].y > maxY {
			maxY = arr[n].y
		}
		if arr[n].z > maxZ {
			maxZ = arr[n].z
		}

		if arr[n].x < minX {
			minX = arr[n].x
		}
		if arr[n].y < minY {
			minY = arr[n].y
		}
		if arr[n].z < minZ {
			minZ = arr[n].z
		}
	}

	fmt.Println(minX, minY, minZ)
	fmt.Println(maxX, maxY, maxZ)

	//fill air not detected by outside surface BFS with lava
	outerAir := BFS3D(arr)

	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			for z := minZ; z <= maxZ; z++ {
				if !cubeExists(arr, x, y, z) && !outerAir[cube{x, y, z, -1}] {
					arr = append(arr, cube{x, y, z, 6})
				}
			}
		}
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

func cubeExists(arr []cube, x int, y int, z int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i].x == x && arr[i].y == y && arr[i].z == z {
			return true
		}
	}
	return false
}

func BFS3D(lava []cube) map[cube]bool {
	visited := make(map[cube]bool)

	start := cube{-1, -1, -1, -1}
	queue := make([]cube, 0)
	queue = append(queue, start)
	visited[start] = true

	for len(queue) != 0 {
		curr := queue[0]
		queue = queue[1:]
		x := curr.x
		y := curr.y
		z := curr.z

		if validRange(x-1, y, z) && !cubeExists(lava, x-1, y, z) && !visited[cube{x - 1, y, z, -1}] {
			air := cube{x - 1, y, z, -1}
			queue = append(queue, air)
			visited[air] = true
		}
		if validRange(x+1, y, z) && !cubeExists(lava, x+1, y, z) && !visited[cube{x + 1, y, z, -1}] {
			air := cube{x + 1, y, z, -1}
			queue = append(queue, air)
			visited[air] = true
		}
		if validRange(x, y-1, z) && !cubeExists(lava, x, y-1, z) && !visited[cube{x, y - 1, z, -1}] {
			air := cube{x, y - 1, z, -1}
			queue = append(queue, air)
			visited[air] = true
		}
		if validRange(x, y+1, z) && !cubeExists(lava, x, y+1, z) && !visited[cube{x, y + 1, z, -1}] {
			air := cube{x, y + 1, z, -1}
			queue = append(queue, air)
			visited[air] = true
		}
		if validRange(x, y, z-1) && !cubeExists(lava, x, y, z-1) && !visited[cube{x, y, z - 1, -1}] {
			air := cube{x, y, z - 1, -1}
			queue = append(queue, air)
			visited[air] = true
		}
		if validRange(x, y, z+1) && !cubeExists(lava, x, y, z+1) && !visited[cube{x, y, z + 1, -1}] {
			air := cube{x, y, z + 1, -1}
			queue = append(queue, air)
			visited[air] = true
		}
	}

	return visited
}

func validRange(x int, y int, z int) bool {
	//between -1,-1,-1 and 20,20,20
	if x < -1 || x > 20 || y < -1 || y > 20 || z < -1 || z > 20 {
		return false
	}
	return true
}
