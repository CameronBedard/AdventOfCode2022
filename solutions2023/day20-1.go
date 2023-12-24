package solutions2023

import (
	"fmt"
	"strings"
)

func Day20p1(lines []string) {
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

	//fmt.Println(directory)
	//initialize inputs
	for _, v := range directory {
		for i := range v.outputs {
			x, ok := directory[v.outputs[i]]
			if ok {
				x.inputs[v.key] = "low"
			}
		}
	}

	low, high := 0, 0
	for b := 0; b < 1000; b++ {
		pulseQueue = append(pulseQueue, Pulse{"button", "low", "broadcaster"})

		for len(pulseQueue) > 0 {
			pulse := pulseQueue[0]
			if pulse.strength == "low" {
				low++
			} else {
				high++
			}
			curr, ok := directory[pulse.to]
			if !ok {
				pulseQueue = pulseQueue[1:]
				continue
			}

			if curr.key == "broadcaster" {
				pulseQueue = queuePulses(directory, pulseQueue, "broadcaster", "low", curr.outputs)
			} else {
				if curr.prefix == "%" {
					//fmt.Println(directory)
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
		}
	}

	fmt.Println("total", low*high)
}

func queuePulses(directory map[string]Module, queue []Pulse, from, strength string, to []string) []Pulse {
	for i := range to {
		queue = append(queue, Pulse{from, strength, to[i]})
		if _, ok := directory[to[i]]; ok {
			directory[to[i]].inputs[from] = strength
		}
	}
	return queue
}

type Module struct {
	prefix  string
	key     string
	onOff   string
	outputs []string
	inputs  map[string]string //module key, strength
}

type Pulse struct {
	from     string
	strength string
	to       string
}
