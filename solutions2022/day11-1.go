package solutions2022

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func day11p1() {
	// Open the file.
	file, err := os.Open("/Users/cameron.bedard/Documents/FunRepos/AdventOfCode2022/inputs2022/day11.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file.
	scanner := bufio.NewScanner(file)

	// Read the file line by line.
	monkeys := [8]Monkey{}
	monkeyCount := [8]int{}
	monkeyIndex := 0
	for scanner.Scan() {
		//line 0
		scanner.Text()
		//line 1
		scanner.Scan()
		line := scanner.Text()
		s1 := strings.Split(line, ":")
		s1 = strings.Split(s1[1], ",")
		queue := NewQueue()
		for i := 0; i < len(s1); i++ {
			queue.Push(atoi(strings.Trim(s1[i], " ")))
		}

		//line 2
		scanner.Scan()
		line = scanner.Text()
		opIsMult := strings.Contains(line, "*")
		s1 = strings.Split(line, " ")
		operation := atoi(s1[len(s1)-1])

		//line 3
		scanner.Scan()
		line = scanner.Text()
		s1 = strings.Split(line, " ")
		testNum := atoi(s1[len(s1)-1])

		//line 4
		scanner.Scan()
		line = scanner.Text()
		s1 = strings.Split(line, " ")
		trueMonkey := atoi(s1[len(s1)-1])

		//line 5
		scanner.Scan()
		line = scanner.Text()
		s1 = strings.Split(line, " ")
		falseMonkey := atoi(s1[len(s1)-1])

		//line 6
		monkeys[monkeyIndex] = Monkey{queue, opIsMult, operation, testNum, trueMonkey, falseMonkey}
		scanner.Scan()
		scanner.Text()
		monkeyIndex++
	}

	//i†ems enter monkey queue first in first out
	//worry from item goes thru operation on inspection, then floor(worry/3)
	//test throws item based on current worry level
	//monkey inspects and throws all items it hasd each turn, a round is monkey 0-7
	for round := 0; round < 20; round++ {
		for m := 0; m < 8; m++ {
			for !monkeys[m].queue.IsEmpty() {
				item := monkeys[m].queue.Pop()

				//peform operation
				if monkeys[m].opIsMult {
					if monkeys[m].operation == 0 {
						item = item * item
					} else {
						item = item * monkeys[m].operation
					}
				} else {
					item = item + monkeys[m].operation
				}
				//divide worry by 3
				item = item / 3

				//test
				if item%monkeys[m].testNum == 0 {
					monkeys[monkeys[m].trueMonkey].queue.Push(item)
				} else {
					monkeys[monkeys[m].falseMonkey].queue.Push(item)
				}

				monkeyCount[m]++
			}
		}
	}

	//fmt.Println("signal strength:", signal)
	fmt.Println(monkeys[0].queue)
	fmt.Printf("%v", monkeyCount)

	// Check for any errors that occurred while reading the file.
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
}

type Monkey struct {
	queue       *Queue
	opIsMult    bool
	operation   int
	testNum     int
	trueMonkey  int
	falseMonkey int
}

// to use q:= NewQueue()
type Queue []int

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Push(x int) {
	*q = append(*q, x)
}

func (q *Queue) IsEmpty() bool {
	h := *q
	l := len(h)
	return l == 0
}

func (q *Queue) Pop() int {
	h := *q
	el := h[0]
	*q = h[1:]

	return el
}
