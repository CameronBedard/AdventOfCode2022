package solutions2023

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day6p1() {
	// Open the file.
	file, err := os.Open("/Users/cameron.bedard/Documents/FunRepos/AdventOfCode2022/inputs2023/day6.txt")
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

	timeTmp := strings.Split(lines[0], " ")
	distTmp := strings.Split(lines[1], " ")
	time := make([]int, len(timeTmp))
	dist := make([]int, len(distTmp))

	for i := range time {
		x, _ := strconv.Atoi(timeTmp[i])
		y, _ := strconv.Atoi(distTmp[i])
		time[i] = x
		dist[i] = y
	}

	total := 0
	for r := 0; r < len(time); r++ {
		ways := 0

		for hold := 1; hold <= time[r]; hold++ {
			timeLeft := time[r] - hold
			boatDist := hold * timeLeft

			if boatDist > dist[r] {
				ways++
			}
		}

		if total == 0 {
			total = ways
		} else {
			total *= ways
		}
		fmt.Println(ways, total)
	}

	fmt.Println("ways: ", total)
}
