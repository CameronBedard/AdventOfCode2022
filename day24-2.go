package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func day24p2() {
	// Open the file.
	file, err := os.Open("/Users/cameron.bedard/Documents/FunRepos/AdventOfCode2022/inputs/day24.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	field := make([]string, 0)
	start := Point{0, 0}
	end := Point{0, 0}
	blizzards := make([]Blizzard, 0)

	// Create a scanner to read the file.
	scanner := bufio.NewScanner(file)
	// Read the file line by line.
	for scanner.Scan() {
		field = append(field, scanner.Text())
	}

	//set start and end
	start.x = strings.Index(field[0], ".")
	end.y = len(field) - 1
	end.x = strings.Index(field[end.y], ".")

	for y, row := range field {
		for x, _ := range row {
			c := row[x]
			switch c {
			case '^':
				blizzards = append(blizzards, Blizzard{
					Point{y, x}, UP, Point{len(field) - 2, x},
				})
			case 'v':
				blizzards = append(blizzards, Blizzard{
					Point{y, x}, DOWN, Point{1, x},
				})
			case '<':
				blizzards = append(blizzards, Blizzard{
					Point{y, x}, LEFT, Point{y, len(field[0]) - 2},
				})
			case '>':
				blizzards = append(blizzards, Blizzard{
					Point{y, x}, RIGHT, Point{y, 1},
				})
			}
		}
	}

	fmt.Println("total:", traverseField(field, blizzards, start, end)+
		traverseField(field, blizzards, end, start)+
		traverseField(field, blizzards, start, end))
}
