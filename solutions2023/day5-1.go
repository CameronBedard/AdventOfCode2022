package solutions2023

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day5p1() {
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
	hasBeenTransformed := make([]bool, len(seeds))

	//parse transformations
	for i := 3; i < len(lines); i++ {
		if lines[i] == "" {
			i++
			//set hasBeen transformed false
			hasBeenTransformed = make([]bool, len(seeds))
		} else {
			//source, destination, range
			xyz := strings.Split(lines[i], " ")
			dest, _ := strconv.Atoi(xyz[0])
			source, _ := strconv.Atoi(xyz[1])
			transformRange, _ := strconv.Atoi(xyz[2])
			transformAmt := dest - source

			for seed := 0; seed < len(seeds); seed++ {
				if !hasBeenTransformed[seed] && inRange(seeds[seed], source, source+transformRange) {
					//seed moved to dest + offset
					seeds[seed] += transformAmt
					hasBeenTransformed[seed] = true
				}
			}
		}
	}

	min := 2000000000
	for i := range seeds {
		if seeds[i] < min {
			min = seeds[i]
		}
	}
	fmt.Println("min seed: ", min)
}

func inRange(seed, l, r int) bool {
	return seed >= l && seed <= r
}
