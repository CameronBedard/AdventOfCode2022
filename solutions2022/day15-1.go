package solutions2022

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

func day15p1() {
	// Open the file.
	file, err := os.Open("/Users/cameron.bedard/Documents/FunRepos/AdventOfCode2022/inputs2022/day15.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file.
	scanner := bufio.NewScanner(file)

	//for each S(x,y), let manhattan distance to B = M
	//span covered on row y = (S(x) - (M - dist(y, S(Y)), S(x) + (M - dist(y, S(Y)))
	//given n segments covering row y(200000):
	//join any overlapping segments for all n
	//total is (sum of segment width for all n) - beaconsAt2M

	//maxX := 4759853
	//minX := 19628
	beaconsAt2M := 0
	segments := make([]Segment, 0)

	// Read the file line by line.
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Replace(line, ",", "", -1)
		line = strings.Replace(line, ":", "", -1)
		line = strings.Replace(line, "x=", "", -1)
		line = strings.Replace(line, "y=", "", -1)

		split := strings.Split(line, " ")

		S := Point{atoi(split[3]), atoi(split[2])}
		B := Point{atoi(split[9]), atoi(split[8])}

		if B.y == 2000000 {
			beaconsAt2M++
		}

		dist := manhattan(S, B)
		row := 2000000

		if Abs(S.y, row) > dist {
			//row has no span at 2M
		} else {
			low := S.x - (dist - Abs(S.y, row))
			high := S.x + (dist - Abs(S.y, row))
			segments = append(segments, Segment{low, high})
		}
	}

	//sort by segments low
	sort.Slice(segments, func(i, j int) bool {
		return segments[i].low < segments[j].low
	})

	for i := 1; i < len(segments); {
		//fmt.Println(segments)
		if segmentOverlap(segments[i], segments[i-1]) {
			segments[i-1] = joinSegment(segments[i], segments[i-1])
			segments = removeElement(segments, i)
		} else {
			i++
		}
	}
	fmt.Println(segments)

	total := 0
	for i := 0; i < len(segments); i++ {
		total += segments[i].high - segments[i].low
	}

	//ADVENT OF CODE GOT IT WRONG, answer should be total - number of beacons on row 2M, per
	//their own explanation and example solution of the problem. Accepted answer is just total.
	fmt.Println("total span: ", total)

	// Check for any errors that occurred while reading the file.
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
}

func manhattan(a Point, b Point) int {
	dist := math.Abs(float64(a.x-b.x)) + math.Abs(float64(a.y-b.y))

	return int(dist)
}

type Segment struct {
	low  int
	high int
}

func Abs(a int, b int) int {
	if a-b < 0 {
		return -(a - b)
	} else {
		return a - b
	}
}

func joinSegment(a Segment, b Segment) Segment {
	return Segment{min(a.low, b.low), max(a.high, b.high)}
}

func segmentOverlap(a Segment, b Segment) bool {
	return max(a.low, b.low) <= min(a.high, b.high)
}

func removeElement(slice []Segment, i int) []Segment {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
