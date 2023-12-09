package solutions2023

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Day9p2(lines []string) {
	entries := make([][]int, len(lines))

	for i := range lines {
		tmp := strings.Split(lines[i], " ")
		line := make([]int, len(tmp))

		for j := range tmp {
			x, _ := strconv.Atoi(tmp[j])
			line[j] = x
		}
		entries[i] = line
	}

	total := 0
	for hist := range entries {
		diffs := make([][]int, 1)
		diffs[0] = entries[hist]

		//while our top row isnt zeroed
		for i := 0; !slices.Equal(diffs[i], make([]int, len(diffs[i]))); i++ {
			//make new arr depth i+1 with -1 length
			diffs = append(diffs, make([]int, len(diffs[i])-1))

			for j := range diffs[i+1] {
				diffs[i+1][j] = diffs[i][j+1] - diffs[i][j]
			}
		}
		//bottom row of diffs is now 0's
		degree := len(diffs) - 1
		diffs[degree] = append(make([]int, 1), diffs[degree]...)

		for depth := degree - 1; depth >= 0; depth-- {
			tmp := make([]int, 1)
			diffs[depth] = append(tmp, diffs[depth]...)
			diffs[depth][0] = diffs[depth][1] - diffs[depth+1][0]
		}

		total += diffs[0][0]
	}

	fmt.Println("total", total)
}
