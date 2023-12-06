package solutions2023

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day5p2() {
	// Open the file.
	file, err := os.Open("/Users/cameron.bedard/Documents/FunRepos/AdventOfCode2022/inputs2023/day5.txt")
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

	//parse seed locations
	tmp := strings.Split(lines[0], ":")
	tmp[1] = strings.Trim(tmp[1], " ")
	seedsStr := strings.Split(tmp[1], " ")

	seeds := make([]int, len(seedsStr))
	for i := range seedsStr {
		x, _ := strconv.Atoi(seedsStr[i])
		seeds[i] = x
	}

	min := 2000000000
	transforms := getTransformArr(lines)

	for seedI := 0; seedI < len(seeds); seedI += 2 {
		minSeed := seeds[seedI]
		maxSeed := seeds[seedI] + seeds[seedI+1] - 1
		for s := minSeed; s <= maxSeed; s++ {
			simSeed := s
			//run thru sim for seed s
			for i := 0; i < 7; i++ {
				for j := 0; j < len(transforms[i]); j++ {
					if inRange(simSeed, transforms[i][j].left, transforms[i][j].right) {
						simSeed += transforms[i][j].amount
						break
					}
				}
			}
			if simSeed < min {
				min = simSeed
			}
		}
		fmt.Println("range complete", seedI, "/", len(seeds))
	}

	fmt.Println("min seed: ", min)
}

// returns len(7)*len(transforms) with each transform being l, r, amt
func getTransformArr(lines []string) [][]trans {
	transforms := make([][]trans, 7)
	transformI := 0
	//parse transformations
	for i := 3; i < len(lines); i++ {
		if lines[i] == "" {
			i++
			transformI++
		} else {
			//source, destination, range
			xyz := strings.Split(lines[i], " ")
			dest, _ := strconv.Atoi(xyz[0])
			source, _ := strconv.Atoi(xyz[1])
			transformRange, _ := strconv.Atoi(xyz[2])
			transformAmt := dest - source

			tmp := trans{source, source + transformRange - 1, transformAmt}
			transforms[transformI] = append(transforms[transformI], tmp)
		}
	}
	return transforms
}

type trans struct {
	left   int
	right  int
	amount int
}
