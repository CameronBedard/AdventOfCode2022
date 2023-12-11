package solutions2023

import "fmt"

func Day11p1(lines []string) {
	rows := make(map[int]bool)
	cols := make(map[int]bool)
	galaxies := make([]Point, 0)
	for i := range lines {
		for j := range lines[0] {
			if lines[i][j] == '#' {
				galaxies = append(galaxies, Point{i, j})
				rows[i] = true
				cols[j] = true
			}
		}
	}

	total := 0
	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			g1 := galaxies[i]
			g2 := galaxies[j]
			distY := 0
			for y := min(g1.y, g2.y) + 1; y <= max(g1.y, g2.y); y++ {
				if rows[y] == false {
					distY += 1000000
				} else {
					distY++
				}
			}
			distX := 0
			for x := min(g1.x, g2.x) + 1; x <= max(g1.x, g2.x); x++ {
				if cols[x] == false {
					distX += 1000000
				} else {
					distX++
				}
			}
			total += distY + distX
		}
	}
	fmt.Println("total", total)
}
