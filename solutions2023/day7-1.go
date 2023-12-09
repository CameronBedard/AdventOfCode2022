package solutions2023

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func Day7p1() {
	// Open the file.
	file, err := os.Open("/Users/cameron.bedard/Documents/FunRepos/AdventOfCode2022/inputs2023/day7.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file.
	scanner := bufio.NewScanner(file)

	// Read the file line by line.
	lines := make([]string, 0)
	hands := make([]entry, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)

		tmp := strings.Split(line, " ")
		x, _ := strconv.Atoi(tmp[1])
		hands = append(hands, entry{tmp[0], x})
	}
	//eof

	sort.SliceStable(hands, func(i, j int) bool {
		return compareHands(hands[i].hand, hands[j].hand) < 0
	})

	fmt.Println(hands)

	total := 0
	for i := range hands {
		total += hands[i].score * (i + 1)
	}

	fmt.Println("total", total)
}

type entry struct {
	hand  string
	score int
}

func compareHands(hand1, hand2 string) int {
	rank := map[uint8]int{
		'A': 14,
		'K': 13,
		'Q': 12,
		'J': 11,
		'T': 10,
		'9': 9,
		'8': 8,
		'7': 7,
		'6': 6,
		'5': 5,
		'4': 4,
		'3': 3,
		'2': 2,
	}

	h1 := make([]int, 15)
	h2 := make([]int, 15)
	for i := range hand1 {
		h1[rank[hand1[i]]]++
		h2[rank[hand2[i]]]++
	}

	h1Score, h2Score := slices.Max(h1), slices.Max(h2)

	//check full houses
	if h1Score == 3 && h2Score == 3 {
		if slices.Contains(h1, 2) {
			h1Score = 4
		}
		if slices.Contains(h2, 2) {
			h2Score = 4
		}
	}

	//check two pair
	if h1Score == 2 && h2Score == 2 {
		h1freq := 1
		h2freq := 1
		for i := range h1 {
			if h1[i] == 2 {
				h1freq++
			}
			if h2[i] == 2 {
				h2freq++
			}
		}
		h1Score = h1freq
		h2Score = h2freq
	}

	if h1Score != h2Score {
		return h1Score - h2Score
	} else {
		//cards have same match type, must be sorted by rank
		for i := range hand1 {
			if hand1[i] == hand2[i] {
				//continue
			} else {
				return rank[hand1[i]] - rank[hand2[i]]
			}
		}
	}
	fmt.Println("two hands the same", hand1, hand2)
	return 0
}
