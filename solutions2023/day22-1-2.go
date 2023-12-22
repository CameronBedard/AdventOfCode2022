package solutions2023

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Day22p1(lines []string) {
	//0,3,171~0,5,171
	bricks := make([]Brick, len(lines))
	for i := range lines {
		tmp := strings.Split(lines[i], "~")
		p1 := strings.Split(tmp[0], ",")
		p2 := strings.Split(tmp[1], ",")

		bricks[i] = Brick{Point3d{atoi(p1[0]), atoi(p1[1]), atoi(p1[2])}, Point3d{atoi(p2[0]), atoi(p2[1]), atoi(p2[2])}}
	}

	//sort by lowest brick first
	sort.SliceStable(bricks, func(i, j int) bool {
		return bricks[i].p1.z < bricks[j].p1.z
	})

	//initial fall
	sim(bricks)

	pt1, pt2 := 0, 0
	for i := range bricks {
		tmp := make([]Brick, len(bricks))
		copy(tmp, bricks)
		//empty brick, run sim and count changes
		tmp[i] = Brick{Point3d{0, 0, 0}, Point3d{0, 0, 0}} //empty brick, run sim and count changes

		fallenBricks, changes := sim(tmp)
		if fallenBricks == 0 {
			pt1++
		}
		pt2 += changes
	}

	fmt.Println("pt1", pt1)
	fmt.Println("pt2", pt2)
}

func sim(bricks []Brick) (int, int) {
	changes, bricksMoved := 0, 0
	for i := range bricks {
		moved := false
		for bricks[i].p1.z > 1 {
			//lower z by 1 and check any overlaps
			newBrick := bricks[i]
			newBrick.p1.z--
			newBrick.p2.z--

			overlap := 0
			for j := i - 1; j >= 0; j-- {
				if overlap3d(newBrick, bricks[j]) {
					overlap++
				}
			}

			if overlap == 0 {
				bricks[i] = newBrick
				changes++
				moved = true
			} else {
				break
			}
		}
		if moved {
			moved = false
			bricksMoved += 1
		}
	}
	return changes, bricksMoved
}

func overlap3d(b1, b2 Brick) bool {
	//overlaps in all 3 dimensions
	return overlap2d(b1.p1.z, b1.p2.z, b2.p1.z, b2.p2.z) && overlap2d(b1.p1.y, b1.p2.y, b2.p1.y, b2.p2.y) && overlap2d(b1.p1.x, b1.p2.x, b2.p1.x, b2.p2.x)
}

func overlap2d(x1, x2, y1, y2 int) bool {
	a1, a2 := min(x1, x2), max(x1, x2)
	b1, b2 := min(y1, y2), max(y1, y2)
	return max(a1, b1) <= min(a2, b2)
}

func atoi(s string) int {
	x, _ := strconv.Atoi(s)
	return x
}

type Point3d struct {
	x int
	y int
	z int
}

type Brick struct {
	p1 Point3d
	p2 Point3d
}
