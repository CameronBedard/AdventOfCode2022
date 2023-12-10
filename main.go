package main

import (
	//"AOC/solutions2022"
	"AOC/solutions2023"
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("/Users/cameron.bedard/Documents/FunRepos/AdventOfCode2022/inputs2023/day10.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	solutions2023.Day10p2(lines)
}
