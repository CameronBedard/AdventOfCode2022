package solutions2022

import (
	"bufio"
	"fmt"
	"os"
)

func day17p1() {
	// Open the file.
	file, err := os.Open("/Users/cameron.bedard/Documents/FunRepos/AdventOfCode2022/inputs/day17.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	shapes := make([][]string, 0)
	shapes = append(shapes, []string{"####"})
	shapes = append(shapes, []string{".#.", "###", ".#."})
	shapes = append(shapes, []string{"..#", "..#", "###"})
	shapes = append(shapes, []string{"#", "#", "#", "#"})
	shapes = append(shapes, []string{"##", "##"})

	cave := make([]string, 0)
	for i := 0; i < 10000; i++ {
		cave = append(cave, ".......")
	}

	// Create a scanner to read the file.
	scanner := bufio.NewScanner(file)
	// Read the file line by line.
	scanner.Scan()
	wind := scanner.Text()
	windI := 0
	minLevel := len(cave) //we start with high number

	for shape := 0; shape < 2022; shape++ {
		shapeI := shape % 5
		offset := 2
		level := minLevel - 3 - len(shapes[shapeI])

		for true {
			//wind
			if wind[windI] == '<' {
				if validPush(offset-1, level, shapes[shapeI], cave) {
					offset--
				}
			}
			if wind[windI] == '>' {
				if validPush(offset+1, level, shapes[shapeI], cave) {
					offset++
				}
			}
			windI = (windI + 1) % len(wind)

			//drop
			if validDrop(offset, level+1, shapes[shapeI], cave) {
				level++
			} else {
				placeRock(offset, level, shapes[shapeI], cave)
				if level < minLevel {
					minLevel = level
				}
				break //next shape
			}
		}
	}

	/*for i := 9500; i < len(cave); i++ {
		fmt.Println(cave[i])
	}*/
	fmt.Println("highest rock:", len(cave)-minLevel)

	// Check for any errors that occurred while reading the file.
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
}

func validPush(offset int, level int, shape []string, cave []string) bool {
	if offset < 0 {
		return false
	}
	if offset+len(shape[0]) > 7 {
		return false
	}

	for i := 0; i < len(shape); i++ {
		for j := 0; j < len(shape[i]); j++ {
			if shape[i][j] == '#' && cave[level+i][offset+j] == '#' {
				return false //contact made with another rock
			}
		}
	}
	return true
}

func validDrop(offset int, level int, shape []string, cave []string) bool {
	if level+len(shape) > len(cave) {
		return false //we hit floor
	}

	for i := 0; i < len(shape); i++ {
		for j := 0; j < len(shape[i]); j++ {
			if shape[i][j] == '#' && cave[level+i][offset+j] == '#' {
				return false //contact made with another rock
			}
		}
	}
	return true
}

func placeRock(offset int, level int, shape []string, cave []string) {
	for i := 0; i < len(shape); i++ {
		for j := 0; j < len(shape[i]); j++ {
			if shape[i][j] == '#' {
				cave[level+i] = replaceAtIndex(cave[level+i], rune(shape[i][j]), offset+j)
			}
		}
		//fmt.Println("cave level", level+i, cave[level+i])
	}
}

func replaceAtIndex(str string, replacement rune, index int) string {
	return str[:index] + string(replacement) + str[index+1:]
}
