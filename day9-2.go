package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func day9p2() {
	// Open the file.
	file, err := os.Open("/Users/cameron.bedard/Documents/FunRepos/AdventOfCode2022/inputs/day9.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file.
	scanner := bufio.NewScanner(file)
	rope := make([]Point, 10)
	pointSet := make(map[Point]bool)
	for i := 0; i < 10; i++ {
		rope[i] = Point{1, 1}
	}
	pointSet[rope[9]] = true

	// Read the file line by line.
	for scanner.Scan() {
		line := scanner.Text()

		instr := strings.Split(line, " ")
		direction := instr[0]
		steps := atoi(instr[1])

		for step := 0; step < steps; step++ {
			if direction == "U" {
				rope[0] = add(rope[0], Point{1, 0})
			}
			if direction == "D" {
				rope[0] = add(rope[0], Point{-1, 0})
			}
			if direction == "L" {
				rope[0] = add(rope[0], Point{0, -1})
			}
			if direction == "R" {
				rope[0] = add(rope[0], Point{0, 1})
			}

			for i := 1; i < 10; i++ {
				rope[i] = updateKnot(rope[i], rope[i-1])
			}
			pointSet[rope[9]] = true
		}
	}

	i := 0
	for key, element := range pointSet {
		fmt.Println("Key:", key, "=>", "Element:", element)
		i++
	}

	fmt.Println("diff tail positions:", i)

	// Check for any errors that occurred while reading the file.
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
}

func updateKnot(tailP Point, headP Point) Point {
	//calc difference betten x,y of head and tail
	//if diff is 0 on x or y, add diffVector/2 to y
	//if diff is 2,1 for x,y or y,x we need a diagonal move ceil(of
	xDiff := float64(headP.x - tailP.x)
	yDiff := float64(headP.y - tailP.y)

	if math.Abs(xDiff) < 2 && math.Abs(yDiff) < 2 {
		//do nothing tail is valid position
	} else if xDiff == 0 || yDiff == 0 {
		//vertical/horizontal move
		moveVector := Point{int(yDiff / 2), int(xDiff / 2)}
		tailP = add(tailP, moveVector)
	} else {
		//diagonal move
		if yDiff > 0 {
			yDiff = math.Ceil(yDiff / 2)
		} else {
			yDiff = math.Floor(yDiff / 2)
		}

		if xDiff > 0 {
			xDiff = math.Ceil(xDiff / 2)
		} else {
			xDiff = math.Floor(xDiff / 2)
		}

		moveVector := Point{int(yDiff), int(xDiff)}
		tailP = add(tailP, moveVector)
	}
	return tailP
}
