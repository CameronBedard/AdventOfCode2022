package solutions2023

import (
	"fmt"
	"strconv"
	"strings"
)

func Day12p2(lines []string) {
	formation := make([]string, len(lines))
	springs := make([][]int, len(lines))
	for i := range lines {
		line := strings.Split(lines[i], " ")
		formation[i] = line[0]
		for j := 0; j < 4; j++ {
			formation[i] += "?" + line[0]
		}

		springTmp := strings.Split(line[1], ",")
		for j := range springTmp {
			x, _ := strconv.Atoi(springTmp[j])
			springs[i] = append(springs[i], x)
		}
		springsTmp := springs[i]
		for j := 0; j < 4; j++ {
			springs[i] = append(springs[i], springsTmp...)
		}
	}

	//2d dp, ?.##?###?#?##? 1,6,1,2
	//i is formation, j is springs
	sum := 0
	for i := range lines {
		cache := make([][]int, len(formation[i]))
		for x := range formation[i] {
			cache[x] = make([]int, len(springs[i])+1)
			for d := 0; d <= len(springs[i]); d++ {
				cache[x][d] = -1
			}
		}
		sum += possible2(0, 0, formation[i], springs[i], cache)
	}
	fmt.Println("total", sum)
}

func possible2(i, j int, formation string, springs []int, cache [][]int) int {
	if i >= len(formation) {
		if j < len(springs) {
			return 0
		}
		return 1
	}

	if cache[i][j] != -1 {
		return cache[i][j]
	}

	res := 0
	if formation[i] == '.' {
		res = possible2(i+1, j, formation, springs, cache)
	} else {
		if formation[i] == '?' {
			res += possible2(i+1, j, formation, springs, cache)
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
					res += possible2(i+count+1, j+1, formation, springs, cache)
				} else {
					res += possible2(i+count, j+1, formation, springs, cache)
				}
			}
		}
	}
	cache[i][j] = res
	return res
}
