package solutions2023

import (
	"fmt"
	"strconv"
	"strings"
)

func Day15p2(lines []string) {
	sequences := strings.Split(lines[0], ",")
	boxes := make([][]lense, 256)
	for i := range boxes {
		boxes[i] = make([]lense, 0)
	}

	for i := range sequences {
		if strings.Contains(sequences[i], "-") {
			//removing
			str := sequences[i][:len(sequences[i])-1]
			h := hash(str)

			for j := range boxes[h] {
				if boxes[h][j].key == str {
					boxes[h] = remove(boxes[h], j)
					break
				}
			}
		} else {
			tmp := strings.Split(sequences[i], "=")
			val, _ := strconv.Atoi(tmp[1])
			l := lense{tmp[0], val}

			h := hash(l.key)

			found := false
			for j := range boxes[h] {
				if boxes[h][j].key == l.key {
					boxes[h][j] = l
					found = true
				}
			}
			if !found {
				boxes[h] = append(boxes[h], l)
			}
		}
	}

	sum := 0
	for i := range boxes {
		for j := range boxes[i] {
			sum += (i + 1) * (j + 1) * boxes[i][j].val
		}
	}
	fmt.Println("total", sum)
}

type lense struct {
	key string
	val int
}

func hash(str string) int {
	curr := 0
	for j := range str {
		curr += int(str[j])
		curr *= 17
		curr = curr % 256
	}
	return curr
}

func remove(slice []lense, s int) []lense {
	return append(slice[:s], slice[s+1:]...)
}
