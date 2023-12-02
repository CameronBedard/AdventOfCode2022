package solutions2022

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

func day15p2() {
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
	s := make([]Point, 0)
	b := make([]Point, 0)

	// Read the file line by line.
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Replace(line, ",", "", -1)
		line = strings.Replace(line, ":", "", -1)
		line = strings.Replace(line, "x=", "", -1)
		line = strings.Replace(line, "y=", "", -1)

		split := strings.Split(line, " ")

		s = append(s, Point{atoi(split[3]), atoi(split[2])})
		b = append(b, Point{atoi(split[9]), atoi(split[8])})
	}

	start := time.Now()

	for y := 0; y <= 4000000; y++ {
		segmentCoverageOfRow(s, b, y)
	}

	timeElapsed := time.Since(start)
	fmt.Println("This function took", timeElapsed)

	// Check for any errors that occurred while reading the file.
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
}

func segmentCoverageOfRow(s []Point, b []Point, row int) {
	segments := make([]Segment, 0)

	for i := 0; i < len(s); i++ {
		S := s[i]
		B := b[i]

		dist := manhattan(S, B)

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

	if len(segments) > 1 {
		fmt.Println("on row:", row, " with segments: ", segments)
		fmt.Println((uint64(segments[0].high+1) * uint64(4000000)) + uint64(row))
	}
}
