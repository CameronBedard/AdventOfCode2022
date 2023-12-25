package solutions2023

import (
	"fmt"
	"strconv"
	"strings"
)

func Day24p1(lines []string) {
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

	sum := 0
	for i := 0; i < len(hail)-1; i++ {
		for j := i + 1; j < len(hail); j++ {
			as, ad := hail[i].pos, hail[i].vel
			bs, bd := hail[j].pos, hail[j].vel
			det := bd.x*ad.y - bd.y*ad.x

			if det != 0 {
				u := ((bs.y-as.y)*bd.x - (bs.x-as.x)*bd.y) / det
				v := ((bs.y-as.y)*ad.x - (bs.x-as.x)*ad.y) / det

				if u >= 0 && v >= 0 {
					resX := bs.x + bd.x*v
					resY := bs.y + bd.y*v

					start := 200000000000000
					end := 400000000000000
					if resX >= start && resY >= start && resX <= end && resY <= end {
						sum++
						//fmt.Println("Intersects:", resY, resX)
					}
				}
			}

		}
	}

	fmt.Println("total", sum)
}

type Hailstone struct {
	pos Point3d
	vel Point3d
}
