package solutions2022

import (
	"bufio"
	"fmt"
	"os"
)

func day23p2() {
	// Open the file.
	file, err := os.Open("/Users/cameron.bedard/Documents/FunRepos/AdventOfCode2022/inputs2022/day23.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	lineN := 0
	positions := make([]Point, 0)

	// Create a scanner to read the file.
	scanner := bufio.NewScanner(file)
	// Read the file line by line.
	for scanner.Scan() {
		line := scanner.Text()

		for i := 0; i < len(line); i++ {
			if line[i] == '#' {
				positions = append(positions, Point{lineN + 150, i + 150})
			}
		}
		lineN++
	}

	//create empty field
	field := make([][]int, 400)
	for i := range field {
		field[i] = make([]int, 400)
	}

	//elves array from parsed positions
	elves := make([]elf, len(positions))
	for i := range positions {
		elves[i] = elf{positions[i], Point{0, 0}}
	}

	//fill initial starting positions
	fillField(field, elves)

	//dirs vector arr
	dirs := [4][3]Point{
		{Point{-1, -1}, Point{-1, 0}, Point{-1, 1}}, /* North */
		{Point{1, -1}, Point{1, 0}, Point{1, 1}},    /* South */
		{Point{-1, -1}, Point{0, -1}, Point{1, -1}}, /* West */
		{Point{-1, 1}, Point{0, 1}, Point{1, 1}},    /* East */
	}

	offset := 0 //whether to start at north, etc
	for round := 0; round < 10000; round++ {
		proposedPoints := make(map[Point]int)

		//list proposals for all elves
		for i := range elves {
			//check if elf is lonely
			isLonely := true
			for index := 0; index < 4; index++ {
				for j := 0; j < 3; j++ {
					p1 := addPoints(elves[i].pos, dirs[index][j])
					if field[p1.y][p1.x] == 1 {
						isLonely = false
					}
				}
			}
			if isLonely {
				elves[i].proposed = Point{-1, -1}
				continue
			}

			//check direction
			for d := offset; d < offset+4; d++ {
				dIndex := d % 4

				p1 := addPoints(elves[i].pos, dirs[dIndex][0])
				p2 := addPoints(elves[i].pos, dirs[dIndex][1])
				p3 := addPoints(elves[i].pos, dirs[dIndex][2])

				if field[p1.y][p1.x] == 0 && field[p2.y][p2.x] == 0 && field[p3.y][p3.x] == 0 { //direction is empty
					elves[i].proposed = p2
					proposedPoints[p2] = proposedPoints[p2] + 1 //tally in our proposal list
					//fmt.Println("elf at pos", elves[i].pos, "proposes moving", dIndex, "to", elves[i].proposed)
					break
				}
			}
		}

		//update positions
		nilProp := Point{-1, -1}
		for i := range elves {
			if proposedPoints[elves[i].proposed] == 1 && elves[i].proposed != nilProp { // only elf to try that position
				elves[i].pos = elves[i].proposed
			}
		}
		for i := range elves {
			elves[i].proposed = nilProp
		}

		emptyField(field)
		fillField(field, elves) // clear and repaint field for next round
		offset++

		if len(proposedPoints) == 0 {
			fmt.Println("round ", round+1)
			break
		}
	}

	// Check for any errors that occurred while reading the file.
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
}
