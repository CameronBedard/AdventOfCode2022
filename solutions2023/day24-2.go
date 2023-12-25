package solutions2023

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Day24p2(lines []string) {
	hail := make([]Hailstone, len(lines))

	for i := range lines {
		lines[i] = strings.ReplaceAll(lines[i], " ", "")
		tmp := strings.Split(lines[i], "@")
		pos := strings.Split(tmp[0], ",")
		vel := strings.Split(tmp[1], ",")
		posNum := make([]int, 3)
		velNum := make([]int, 3)

		for j := range posNum {
			x, _ := strconv.Atoi(pos[j])
			y, _ := strconv.Atoi(vel[j])
			posNum[j] = x
			velNum[j] = y
		}
		hail[i] = Hailstone{Point3d{posNum[0], posNum[1], posNum[2]}, Point3d{velNum[0], velNum[1], velNum[2]}}
	}

	maybeX, maybeY, maybeZ := []int{}, []int{}, []int{}
	for i := 0; i < len(hail)-1; i++ {
		for j := i + 1; j < len(hail); j++ {
			a, b := hail[i], hail[j]
			if a.vel.x == b.vel.x {
				nextMaybe := findMatchingVel(b.pos.x-a.pos.x, a.vel.x)
				if len(maybeX) == 0 {
					maybeX = nextMaybe
				} else {
					maybeX = getIntersect(maybeX, nextMaybe)
				}
			}
			if a.vel.y == b.vel.y {
				nextMaybe := findMatchingVel(b.pos.y-a.pos.y, a.vel.y)
				if len(maybeY) == 0 {
					maybeY = nextMaybe
				} else {
					maybeY = getIntersect(maybeY, nextMaybe)
				}
			}
			if a.vel.z == b.vel.z {
				nextMaybe := findMatchingVel(b.pos.z-a.pos.z, a.vel.z)
				if len(maybeZ) == 0 {
					maybeZ = nextMaybe
				} else {
					maybeZ = getIntersect(maybeZ, nextMaybe)
				}
			}
		}
	}

	var result = 0
	if len(maybeX) == len(maybeY) && len(maybeY) == len(maybeZ) && len(maybeZ) == 1 {
		// only one possible velocity in all dimensions
		rockVel := Float3d{float64(maybeX[0]), float64(maybeY[0]), float64(maybeZ[0])}
		hailStoneA, hailStoneB := hail[0], hail[1]
		mA := (float64(hailStoneA.vel.y) - rockVel.y) / (float64(hailStoneA.vel.x) - rockVel.x)
		mB := (float64(hailStoneB.vel.y) - rockVel.y) / (float64(hailStoneB.vel.x) - rockVel.x)
		cA := float64(hailStoneA.pos.y) - (mA * float64(hailStoneA.pos.x))
		cB := float64(hailStoneB.pos.y) - (mB * float64(hailStoneB.pos.x))
		xPos := (cB - cA) / (mA - mB)
		yPos := mA*xPos + cA
		time := (xPos - float64(hailStoneA.pos.x)) / (float64(hailStoneA.vel.x) - rockVel.x)
		zPos := float64(hailStoneA.pos.z) + (float64(hailStoneA.vel.z)-rockVel.z)*time
		result = int(xPos + yPos + zPos)
	}

	fmt.Println("total", result)
}

func intersects(a, b Hailstone) bool {
	as, ad := a.pos, a.vel
	bs, bd := b.pos, b.vel
	det := bd.x*ad.y - bd.y*ad.x

	if det != 0 {
		u := ((bs.y-as.y)*bd.x - (bs.x-as.x)*bd.y) / det
		v := ((bs.y-as.y)*ad.x - (bs.x-as.x)*ad.y) / det

		if u >= 0 && v >= 0 {
			//resX := bs.x + bd.x*v
			//resY := bs.y + bd.y*v
			return true
		}
	}
	return false
}

func getIntersect(a, b []int) []int {
	result := []int{}
	for _, val := range a {
		if slices.Contains(b, val) {
			result = append(result, val)
		}
	}
	return result
}

func findMatchingVel(dvel, pv int) []int {
	match := []int{}
	for v := -1000; v < 1000; v++ {
		if v != pv && dvel%(v-pv) == 0 {
			match = append(match, v)
		}
	}
	return match
}

type Float3d struct {
	x, y, z float64
}
