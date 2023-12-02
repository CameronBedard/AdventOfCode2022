package solutions2022

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func day20p2() {
	// Open the file.
	file, err := os.Open("/Users/cameron.bedard/Documents/FunRepos/AdventOfCode2022/inputs/day20.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	a := make([]value, 0)
	index := 0

	// Create a scanner to read the file.
	scanner := bufio.NewScanner(file)

	// Read the file line by line.
	for scanner.Scan() {
		line := scanner.Text()
		x := atoi(line) * 811589153

		a = append(a, value{x, index})
		index++
	}
	start := time.Now()

	for k := 0; k < 10; k++ {
		for i := 0; i < len(a); i++ {
			for j := 0; j < len(a); j++ {
				if a[j].order == i {
					x := a[j]
					a = remove(a, j)

					newI := floorMod(x.val+j, len(a))
					a = insert(a, newI, x)

					break
				}
			}
		}
	}

	for j := 0; j < len(a); j++ {
		if a[j].val == 0 {
			fmt.Println(a[(j+1000)%len(a)].val + a[(j+2000)%len(a)].val + a[(j+3000)%len(a)].val)
			break
		}
	}
	//fmt.Println(a[1000].val + a[2000].val + a[3000].val)
	timeElapsed := time.Since(start)
	fmt.Println("This function took", timeElapsed)

	// Check for any errors that occurred while reading the file.
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
}
