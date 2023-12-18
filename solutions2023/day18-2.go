package solutions2023

import (
	"fmt"
	"strconv"
	"strings"
)

func Day18p2(lines []string) {
	direction := map[uint8]Point{
		'1': south,
		'3': north,
		'2': west,
		'0': east,
	}

	p := Point{0, 0}
	area := 0
	for i := range lines {
		inputs := strings.Split(lines[i], " ")[2]
		inputs = strings.Trim(inputs, "(#)")
		dir := direction[inputs[len(inputs)-1]]
		dist, _ := strconv.ParseInt(inputs[:len(inputs)-1], 16, strconv.IntSize)
		//fmt.Println(dir, newNum)

		nextP := addPoints(p, Point{dir.y * int(dist), dir.x * int(dist)})
		area += (p.x * nextP.y) - (p.y * nextP.x) + int(dist)
		p = nextP
	}
	area = area/2 + 1

	fmt.Println("total", area)
}
