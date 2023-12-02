package solutions2022

import (
	"bufio"
	"fmt"
	"os"
)

func day17p2() {
	// Open the file.
	file, err := os.Open("/Users/cameron.bedard/Documents/FunRepos/AdventOfCode2022/inputs/day17.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	shapes := make([][]string, 0)
	shapes = append(shapes, []string{"####"})
	shapes = append(shapes, []string{".#.", "###", ".#."})
	shapes = append(shapes, []string{"..#", "..#", "###"})
	shapes = append(shapes, []string{"#", "#", "#", "#"})
	shapes = append(shapes, []string{"##", "##"})

	cave := make([]string, 0)
	for i := 0; i < 100000; i++ {
		cave = append(cave, ".......")
	}

	// Create a scanner to read the file.
	scanner := bufio.NewScanner(file)
	// Read the file line by line.
	scanner.Scan()
	wind := scanner.Text()
	windI := 0
	minLevel := len(cave)     //we start with high number
	rotation := len(wind) * 5 //how many shapes before the cave starts repeating
	rockHeightAtRockI := make([]int, 0)
	rockHeightAtRockI = append(rockHeightAtRockI, 0)

	for shape := 0; shape < rotation; shape++ {
		shapeI := shape % 5
		offset := 2
		level := minLevel - 3 - len(shapes[shapeI])

		for true {
			//wind
			if wind[windI] == '<' {
				if validPush(offset-1, level, shapes[shapeI], cave) {
					offset--
				}
			}
			if wind[windI] == '>' {
				if validPush(offset+1, level, shapes[shapeI], cave) {
					offset++
				}
			}
			windI = (windI + 1) % len(wind)

			//drop
			if validDrop(offset, level+1, shapes[shapeI], cave) {
				level++
			} else {
				placeRock(offset, level, shapes[shapeI], cave)
				if level < minLevel {
					minLevel = level
				}
				rockHeightAtRockI = append(rockHeightAtRockI, len(cave)-minLevel)
				break //next shape
			}
		}
	}

	m := make(map[[100]int]bool)
	a := [100]int{}
	for i := 1500; i < 1600; i++ {
		a[i-1500] = rockHeightAtRockI[i] - rockHeightAtRockI[i-1]
	}
	m[a] = true
	cycleRepeats := -1

	for i := 1551; i < 5000; i++ {
		a2 := [100]int{}
		for j := i; j < i+100; j++ {
			a2[j-i] = rockHeightAtRockI[j] - rockHeightAtRockI[j-1]
		}

		if m[a2] == true {
			cycleRepeats = i
			fmt.Println("cyclic repeats between: ", 1500, cycleRepeats)
			break
		}
	}

	//cycle occurs from 1500 -> (cycleRepeats-1)
	fmt.Println(rockHeightAtRockI[cycleRepeats]-rockHeightAtRockI[1500], 3239-1500)
	rocks := 1000000000000
	totalHeight := rockHeightAtRockI[1500]
	rocks -= 1500

	cycleLength := cycleRepeats - 1500
	cycleNumber := rocks / cycleLength
	heightEachCycle := rockHeightAtRockI[cycleRepeats] - rockHeightAtRockI[1500]

	//run cycles over height and rock numbers
	totalHeight += cycleNumber * heightEachCycle
	rocks -= cycleLength * cycleNumber

	//when finding the last remainder it must be started from where the cycle would be at that point
	totalHeight += rockHeightAtRockI[rocks+1500] - rockHeightAtRockI[1500]

	fmt.Println("highest rock:", totalHeight)

	// Check for any errors that occurred while reading the file.
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
}
