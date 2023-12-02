package solutions2022

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func day5p2() {
	// Open the file.
	file, err := os.Open("/Users/cameron.bedard/Documents/FunRepos/AdventOfCode2022/inputs2022/day5.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file.
	scanner := bufio.NewScanner(file)

	// Our shipping yard
	yard := make([]Stack, 9)
	for i := 0; i < 9; i++ {
		yard[i] = Stack{}
	}

	// Read the file line by line.
	ind := 0
	for scanner.Scan() {
		line := scanner.Text()

		if ind < 8 {
			fmt.Println("yard: ", line)

			for i := 0; i < 9; i++ {
				s := line[(i * 4) : (i*4)+3]
				if strings.Contains(s, "[") {
					yard[i].Push(s)
				}
			}

		}

		if ind == 8 {
			yard2 := make([]Stack, 9)
			for i := 0; i < 9; i++ {
				yard2[i] = Stack{}

				for !yard[i].IsEmpty() {
					s, success := yard[i].Pop()

					if success {
						yard2[i].Push(s)
					}
				}
			}
			yard = yard2
		}

		if ind > 9 {
			//process stacks
			instructions := strings.Split(line, " ")

			amt := atoi(instructions[1])
			from := atoi(instructions[3]) - 1
			to := atoi(instructions[5]) - 1

			lift := Stack{}
			for i := 0; i < amt; i++ {
				s, success := yard[from].Pop()

				if success {
					lift.Push(s)
				}
			}

			for !lift.IsEmpty() {
				s, success := lift.Pop()

				if success {
					yard[to].Push(s)
				}
			}
		}

		ind++
	}

	for i := 0; i < 9; i++ {
		s, success := yard[i].Pop()

		if !success {
			fmt.Print("fail", " ")
		} else {
			fmt.Print(s, " ")
		}

	}

	// Check for any errors that occurred while reading the file.
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
}
