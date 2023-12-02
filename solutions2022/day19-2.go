package solutions2022

import (
	"bufio"
	"fmt"
	"os"
)

func day19p2() {
	// Open the file.
	file, err := os.Open("/Users/cameron.bedard/Documents/FunRepos/AdventOfCode2022/inputs2022/day19.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	result := 1
	index := 0

	// Create a scanner to read the file.
	scanner := bufio.NewScanner(file)
	// Read the file line by line.
	for scanner.Scan() {
		var id int
		oreRobot := oreRobot{}
		clayRobot := clayRobot{}
		obsRobot := obsRobot{}
		geodeRobot := geodeRobot{}

		fmt.Sscanf(scanner.Text(),
			"Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.",
			&id, &oreRobot.oreCost, &clayRobot.oreCost, &obsRobot.oreCost, &obsRobot.clayCost, &geodeRobot.oreCost, &geodeRobot.obsCost)

		bp := blueprint{
			id,
			oreRobot,
			clayRobot,
			obsRobot,
			geodeRobot,
		}

		local := search(bp, 0, 0, 0, 32, 1, 0, 0, 0, 0)
		result *= local
		globalBest = 0
		fmt.Println(bp, result, local)
		index++
		if index == 3 {
			break
		}
	}

	fmt.Println("total:", result)

	// Check for any errors that occurred while reading the file.
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
}
