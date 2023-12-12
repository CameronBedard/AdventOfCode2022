package solutions2023

import (
	"fmt"
	"strconv"
	"strings"
)

func Day12p1(lines []string) {
	formation := make([]string, len(lines))
	springs := make([][]int, len(lines))
	for i := range lines {
		line := strings.Split(lines[i], " ")
		formation[i] = line[0]

		springTmp := strings.Split(line[1], ",")
		for j := range springTmp {
			x, _ := strconv.Atoi(springTmp[j])
			springs[i] = append(springs[i], x)
		}
	}

	//2d dp, ?.##?###?#?##? 1,6,1,2
	//i is formation, j is springs
	sum := 0
	for i := range lines {
		sum += possible(0, 0, formation[i], springs[i])
	}
	fmt.Println("total", sum)
}

func possible(i, j int, formation string, springs []int) int {
	if i >= len(formation) {
		if j < len(springs) {
			return 0
		}
		return 1
	}

	res := 0
	if formation[i] == '.' {
		res = possible(i+1, j, formation, springs)
	} else {
		if formation[i] == '?' {
			res += possible(i+1, j, formation, springs)
		}
		if j < len(springs) {
			count := 0
			for k := i; k < len(formation); k++ {
				if count > springs[j] || formation[k] == '.' || (count == springs[j] && formation[k] == '?') {
					break
				}
				count++
			}
			if count == springs[j] {
				if i+count < len(formation) && formation[i+count] != '#' {
					res += possible(i+count+1, j+1, formation, springs)
				} else {
					res += possible(i+count, j+1, formation, springs)
				}
			}
		}
	}
	return res
}
