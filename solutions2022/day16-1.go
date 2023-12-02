package solutions2022

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func day16p1() {
	// Open the file.
	file, err := os.Open("/Users/cameron.bedard/Documents/FunRepos/AdventOfCode2022/inputs/day16test.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	valves := make(map[string]valve)
	iToValve := make(map[int]string)
	valveToI := make(map[string]int)
	index := 0

	// Create a scanner to read the file.
	scanner := bufio.NewScanner(file)
	// Read the file line by line.
	for scanner.Scan() {
		line := scanner.Text()
		//Valve DB has flow rate=0; tunnels lead to valves AC, UN, ...
		tokens := strings.Split(line, " ")
		valveName := tokens[1]
		flowRate := atoi(strings.Trim(tokens[4], "rate=;"))

		v := valve{valveName, flowRate, make([]string, 0)}
		for i := 9; i < len(tokens); i++ {
			v.adj = append(v.adj, strings.Trim(tokens[i], ","))
		}

		valves[valveName] = v

		valveToI[valveName] = index
		iToValve[index] = valveName
		index++
	}

	n := len(valves) //66
	nonZero := make(map[int]bool)
	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			dist[i][j] = 9999999
		}
	}

	for name, v := range valves {
		for i := 0; i < len(v.adj); i++ {
			dist[valveToI[name]][valveToI[v.adj[i]]] = 1
		}
		dist[valveToI[name]][valveToI[name]] = 0

		if v.flowRate > 0 {
			nonZero[valveToI[name]] = true
		}
	}

	//floyd warshall
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][j] > dist[i][k]+dist[k][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}

	fmt.Println(maxFlow(valveToI["AA"], 30, dist, make([]int, 66), nonZero, 0, valves, iToValve))
	fmt.Println(globalMax)
	fmt.Println(paths)
}

type valve struct {
	valveName string
	flowRate  int
	adj       []string
}

var globalMax = 0
var paths = 0

func maxFlow(current int, time int, dist [][]int, visited []int, nonZero map[int]bool, flow int, valves map[string]valve, iToValve map[int]string) int {
	if time <= 0 {
		paths++
		if globalMax < flow {
			globalMax = flow
			/*for i := 0; i < len(visited); i++ {
				if visited[i] > 0 {
					fmt.Print(valves[iToValve[i]].valveName, ":", visited[i], " ")
				}
			}
			fmt.Println()*/ // get path history
		}
		return flow
	}

	flow += valves[iToValve[current]].flowRate * time
	visited[current] = time

	maximum := flow
	for i := 0; i < len(dist[current]); i++ {
		if visited[i] == 0 && nonZero[i] {
			newVis := make([]int, len(visited))
			copy(newVis, visited)
			maximum = max(maximum, maxFlow(i, time-dist[current][i]-1, dist, newVis, nonZero, flow, valves, iToValve))
		}
	}

	if globalMax < flow {
		globalMax = flow
	}
	return flow
}
