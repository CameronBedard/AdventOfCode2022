package solutions2023

import (
	"fmt"
	"strconv"
	"strings"
)

func Day19p1(lines []string) {
	rules := make(map[string][]string) //rule has key xdf: x<123:A or m<456:xdf or another key lkj or A/R
	parts := make([]map[string]int, 0) //x, m, a, s

	processRules := true
	for i := range lines {
		if lines[i] == "" {
			processRules = false
			continue
		}
		if processRules { //rules
			tmp := strings.Split(lines[i], "{")
			tmp[1] = strings.Trim(tmp[1], "}")

			rules[tmp[0]] = strings.Split(tmp[1], ",")
		} else { //parts
			tmp := strings.Split(strings.Trim(lines[i], "{}"), ",")
			parts = append(parts, make(map[string]int))
			for j := range tmp {
				x, _ := strconv.Atoi(tmp[j][2:])
				parts[len(parts)-1][tmp[j][0:1]] = x
			}
		}
	}

	total := 0
	for _, part := range parts {
		partSum := 0
		for _, v := range part {
			partSum += v
		}

		currRule := "in"
		for currRule != "A" && currRule != "R" {
			for _, rule := range rules[currRule] {
				dest := processNextRule(rule, part)

				if dest != "false" {
					currRule = dest
					break
				}
			}
		}

		//fmt.Println(currRule, part)
		if currRule == "A" {
			total += partSum
		}
	}

	fmt.Println("total", total)
}

func processNextRule(rule string, part map[string]int) string {
	//rule has key xdf: x<123:A or m<456:xdf or another key lkj or A/R
	if !strings.Contains(rule, ":") {
		//no eval
		return rule
	}

	dest := "false"
	ruleArr := strings.Split(rule, ":")    //ruleArr[1] is dest if eval is true
	if strings.Contains(ruleArr[0], "<") { //less than
		num, _ := strconv.Atoi(ruleArr[0][2:])
		letter := ruleArr[0][0:1]
		if part[letter] < num {
			return ruleArr[1]
		} else {
			return dest
		}
	} else { //greater than
		num, _ := strconv.Atoi(ruleArr[0][2:])
		letter := ruleArr[0][0:1]
		if part[letter] > num {
			return ruleArr[1]
		} else {
			return dest
		}
	}
}
