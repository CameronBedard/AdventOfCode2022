package main

import (
	"bufio"
	"container/ring"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func day22p1() {
	// Open the file.
	file, err := os.Open("/Users/cameron.bedard/Documents/FunRepos/AdventOfCode2022/inputs/day22.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	lineI := 0
	field := make([][]pos, 200)
	instructions := ""

	// Create a scanner to read the file.
	scanner := bufio.NewScanner(file)
	// Read the file line by line.
	for scanner.Scan() {
		line := scanner.Text()

		if lineI < 200 { //parsing cube
			for len(line) < 200 { //make all lines the same length so bounds check is easy
				line = line + " "
			}

			field[lineI] = make([]pos, 200)
			for i := 0; i < 200; i++ {
				field[lineI][i] = pos{int(line[i]), lineI, i, nil, nil, nil, nil}
			}
		}

		if lineI == 201 { //move sequence
			instructions = line
		}

		lineI++
	}

	//assign neighbours
	for i := 0; i < 200; i++ {
		for j := 0; j < 200; j++ {
			if inCube(i-1, j) && (field[i-1][j].val == '.' || field[i-1][j].val == '#') {
				field[i][j].up = &field[i-1][j]
			}
			if inCube(i+1, j) && (field[i+1][j].val == '.' || field[i+1][j].val == '#') {
				field[i][j].down = &field[i+1][j]
			}
			if inCube(i, j-1) && (field[i][j-1].val == '.' || field[i][j-1].val == '#') {
				field[i][j].left = &field[i][j-1]
			}
			if inCube(i, j+1) && (field[i][j+1].val == '.' || field[i][j+1].val == '#') {
				field[i][j].right = &field[i][j+1]
			}
		}
	}

	//stitch sides together
	stitchV(field, 50, 149, 0, 49)
	stitchV(field, 50, 99, 50, 99)
	stitchV(field, 0, 99, 100, 149)
	stitchV(field, 0, 49, 150, 199)

	stitchH(field[0], field[49], 100, 149)
	stitchH(field[0], field[149], 50, 99)
	stitchH(field[100], field[199], 0, 49)

	// run move sequence
	steps, turns := parseInstruction(instructions)
	dirs := NewCircularList("R", "D", "L", "U") //dirs starts at r
	current := field[0][50]

	for i := 0; i < len(steps); i++ {
		next := current

		for steps[i] > 0 {
			//fmt.Println(current, steps[i], dirs.GetCurr())
			switch dirs.GetCurr() {
			case "R":
				if current.right.val != '#' {
					next = *current.right
				}
			case "D":
				if current.down.val != '#' {
					next = *current.down
				}
			case "L":
				if current.left.val != '#' {
					next = *current.left
				}
			case "U":
				if current.up.val != '#' {
					next = *current.up
				}
			}
			steps[i]--
			current = next
		}

		if i < len(turns) {
			if turns[i] == 'R' {
				dirs.GetNext()
			} else {
				dirs.GetPrev()
			}
		}
	}

	current.row++
	current.col++
	facing := 0
	switch dirs.GetCurr() {
	case "R":
		facing = 0
	case "D":
		facing = 1
	case "L":
		facing = 2
	case "U":
		facing = 3
	}
	fmt.Println("final coords:", current.row, current.col)
	fmt.Println("password:", (1000*current.row)+(4*current.col)+facing)
}

func stitchH(topRow []pos, bottomRow []pos, start, stop int) {
	for i := start; i <= stop; i++ {
		topRow[i].up = &bottomRow[i]
		bottomRow[i].down = &topRow[i]
	}
}

func stitchV(field [][]pos, firstCol, lastCol, floor, ceil int) {
	for i := floor; i <= ceil; i++ {
		field[i][lastCol].right = &field[i][firstCol]
		field[i][firstCol].left = &field[i][lastCol]
	}
}

func inCube(row, col int) bool {
	return row >= 0 && col >= 0 && row < 200 && col < 200
}

type pos struct {
	val   int
	row   int
	col   int
	up    *pos
	right *pos
	left  *pos
	down  *pos
}

// parseInstructions author /u/Krethas
func parseInstruction(instruction string) ([]int, []byte) {
	moves, turns := make([]int, 0), make([]byte, 0)
	for len(instruction) > 0 {
		nextTurn := strings.IndexAny(instruction, "LR")
		if nextTurn == -1 {
			if tiles, err := strconv.Atoi(instruction); err == nil {
				moves = append(moves, tiles)
				break
			} else {
				panic(err)
			}
		} else {
			if tiles, err := strconv.Atoi(instruction[:nextTurn]); err == nil {
				moves = append(moves, tiles)
			} else {
				panic(err)
			}
			turns = append(turns, instruction[nextTurn])
			instruction = instruction[nextTurn+1:]
		}
	}
	return moves, turns
}

type CircularList struct {
	list *ring.Ring
}

func NewCircularList(items ...interface{}) *CircularList {
	cl := &CircularList{
		list: ring.New(len(items)),
	}
	for i := 0; i < cl.list.Len(); i++ {
		cl.list.Value = items[i]
		cl.list = cl.list.Next()
	}
	return cl
}

func (cl *CircularList) ShowAll() {
	cl.list.Do(func(x interface{}) {
		fmt.Printf("Item: %v\n", x)
	})
}

func (cl *CircularList) GetCurr() interface{} {
	val := cl.list.Value
	return val
}

func (cl *CircularList) GetNext() interface{} {
	cl.list = cl.list.Next()
	val := cl.list.Value
	return val
}

func (cl *CircularList) GetPrev() interface{} {
	cl.list = cl.list.Prev()
	val := cl.list.Value
	return val
}
