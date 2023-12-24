package solutions2023

import (
	"fmt"
	"strings"
)

func Day20p2(lines []string) {
	/*
		broadcaster -> a, b, c
		%a -> b
		%b -> c
		%c -> inv
		&inv -> a
	*/
	//LEARNING ACTIVITY: Solve with modules being interfaces
	directory := make(map[string]Module)
	pulseQueue := make([]Pulse, 0)

	for i := range lines {
		lines[i] = strings.ReplaceAll(lines[i], " ", "")
		tmp := strings.Split(lines[i], "->")
		outputs := strings.Split(tmp[1], ",")

		if tmp[0] == "broadcaster" {
			m := Module{"", "broadcaster", "off", outputs, make(map[string]string)}
			directory[m.key] = m
		} else {
			m := Module{tmp[0][0:1], tmp[0][1:], "off", outputs, make(map[string]string)}
			directory[m.key] = m
		}
	}

	//initialize inputs
	for _, v := range directory {
		for i := range v.outputs {
			x, ok := directory[v.outputs[i]]
			if ok {
				x.inputs[v.key] = "low"
			}
		}
	}

	lcmInputs := make([]string, 0) //gh feeds rx
	for i := range directory["gh"].inputs {
		lcmInputs = append(lcmInputs, directory["gh"].inputs[i])
	}
	fmt.Println(lcmInputs, directory["gh"].inputs)
	factors := make(map[string]int)

	cycles := 0
bigLoop:
	for b := 0; b < 200000000; b++ {
		pulseQueue = append(pulseQueue, Pulse{"button", "low", "broadcaster"})

		for len(pulseQueue) > 0 {
			pulse := pulseQueue[0]
			cycles++

			curr, ok := directory[pulse.to]
			if !ok {
				pulseQueue = pulseQueue[1:]
				continue
			}

			if curr.key == "broadcaster" {
				pulseQueue = queuePulses(directory, pulseQueue, "broadcaster", "low", curr.outputs)
			} else {
				if curr.prefix == "%" {
					if pulse.strength == "high" {
						//nothing, dont queue new pulse
					} else { //pulse low
						if curr.onOff == "off" {
							pulseQueue = queuePulses(directory, pulseQueue, curr.key, "high", curr.outputs)
							directory[curr.key] = Module{curr.prefix, curr.key, "on", curr.outputs, curr.inputs}
						} else {
							pulseQueue = queuePulses(directory, pulseQueue, curr.key, "low", curr.outputs)
							directory[curr.key] = Module{curr.prefix, curr.key, "off", curr.outputs, curr.inputs}
						}
					}
				}
				if curr.prefix == "&" {
					sendLow := true
					for _, v := range curr.inputs {
						if v == "low" {
							sendLow = false
							break
						}
					}
					if sendLow {
						pulseQueue = queuePulses(directory, pulseQueue, curr.key, "low", curr.outputs)
					} else {
						pulseQueue = queuePulses(directory, pulseQueue, curr.key, "high", curr.outputs)
					}
				}
			}
			pulseQueue = pulseQueue[1:]
		} //sim ended

		for k, v := range directory["gh"].inputs {
			if v == "high" {
				fmt.Println(k, v, b)
			}
		}
		if len(lcmInputs) == len(factors) {
			break bigLoop
		}
	}

	product := 1
	for _, v := range factors {
		product *= v
	}

	fmt.Println("total", product)
}
