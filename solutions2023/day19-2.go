package solutions2023

import (
	"fmt"
	"strconv"
	"strings"
)

func Day19p2(lines []string) {
	rules := make(map[string][]string) //rule has key xdf: x<123:A or m<456:xdf or another key lkj or A/R
	parts := make([]Part, 1)           //x, m, a, s
	parts[0] = Part{
		"in",
		map[string]Interval{
			"x": {1, 4000},
			"m": {1, 4000},
			"a": {1, 4000},
			"s": {1, 4000},
		},
	}

	processRules := true
	for i := range lines {
		if lines[i] == "" {
			processRules = false
			break
		}
		if processRules { //rules
			tmp := strings.Split(lines[i], "{")
			tmp[1] = strings.Trim(tmp[1], "}")

			rules[tmp[0]] = strings.Split(tmp[1], ",")
		}
	}

	//change parts to queue where we start with 1-4000 and append any splits
	//rules should be a map with their currRule eval stored
	total := 0
	for len(parts) > 0 {
		currRule := parts[0].rule
		for currRule != "A" && currRule != "R" {
			for _, rule := range rules[currRule] {
				if !strings.Contains(rule, ":") { //no eval to be done
					currRule = rule
					break
				}
				//for each eval we will have 2 resulting parts, 1 passes and 1 fails
				//we continue working on the failing part because it is in this same rule
				//we save the passing part in the parts queue w/ its key as destination node
				currPart, passingPart := processNextEval(rule, parts[0])
				parts[0] = currPart
				parts = append(parts, passingPart)
			}
		}

		//calculating part combos
		partSum := 1
		if currRule == "A" {
			for _, v := range parts[0].component {
				partSum *= v.high - v.low + 1
			}
			total += partSum
		}
		//delete part from queue
		if len(parts) == 1 {
			break
		} else {
			parts = parts[1:]
		}
	}

	fmt.Println("total", total)
}

type Part struct {
	rule      string
	component map[string]Interval
}

type Interval struct {
	low  int
	high int
}

func processNextEval(rule string, part Part) (Part, Part) {
	//rule has key xdf: x<123:A or m<456:xdf
	ruleArr := strings.Split(rule, ":") //ruleArr[1] is dest if eval is true
	num, _ := strconv.Atoi(ruleArr[0][2:])
	letter := ruleArr[0][0:1]

	passingPart := Part{
		rule:      ruleArr[1],
		component: make(map[string]Interval),
	}
	for k, v := range part.component {
		passingPart.component[k] = v
	}

	if strings.Contains(ruleArr[0], "<") { //less than
		passingPart.component[letter] = Interval{passingPart.component[letter].low, num - 1}
		part.component[letter] = Interval{num, part.component[letter].high}
	} else { //greater than
		passingPart.component[letter] = Interval{num + 1, passingPart.component[letter].high}
		part.component[letter] = Interval{part.component[letter].low, num}
	}

	return part, passingPart
}
