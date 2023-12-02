package solutions2022

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func day25p1() {
	// Open the file.
	file, err := os.Open("/Users/cameron.bedard/Documents/FunRepos/AdventOfCode2022/inputs2022/day25.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	digits := make(map[byte]int)
	digits['='] = -2
	digits['-'] = -1
	digits['0'] = 0
	digits['1'] = 1
	digits['2'] = 2

	total := 0

	// Create a scanner to read the file.
	scanner := bufio.NewScanner(file)
	// Read the file line by line.
	for scanner.Scan() {
		line := scanner.Text()

		curr := 0
		pow := len(line) - 1
		for i := 0; i < len(line); i++ {
			curr += digits[line[i]] * int(math.Pow(5, float64(pow)))
			pow--
		}

		total += curr
	}

	fmt.Println("sum:", total)
	fmt.Println("sum:", intToSnafu(total))

	// Check for any errors that occurred while reading the file.
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
}

func intToSnafu(decimal int) string {
	snafu := ""
	for decimal != 0 {
		remainder := decimal % 5

		switch remainder {
		case 0:
			snafu = "0" + snafu
			break
		case 1:
			snafu = "1" + snafu
			break
		case 2:
			snafu = strconv.Itoa(remainder) + snafu
			break
		case 3:
			snafu = "=" + snafu
			decimal += 5
			break
		case 4:
			snafu = "-" + snafu
			decimal += 5
			break
		}
		decimal = decimal / 5
	}
	return snafu
}
