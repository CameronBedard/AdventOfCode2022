package solutions2023

import (
	"fmt"
	"strings"
)

func Day15p1(lines []string) {
	sequences := strings.Split(lines[0], ",")

	sum := 0
	for i := range sequences {
		curr := 0
		for j := range sequences[i] {
			curr += int(sequences[i][j])
			curr *= 17
			curr = curr % 256
		}
		sum += curr
	}

	fmt.Println("total", sum)
}
