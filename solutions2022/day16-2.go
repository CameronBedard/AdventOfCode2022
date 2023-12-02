package solutions2022

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func day16p2() {
	// Open the file.
	file, err := os.Open("/Users/cameron.bedard/Documents/FunRepos/AdventOfCode2022/inputs/day16.txt")
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

	maxFlow2(valveToI["AA"], 26, dist, make([]int, 66), nonZero, 0, valves, iToValve)
	fmt.Println(globalMax)
	sort.Slice(distinctPaths, func(i, j int) bool {
		return distinctPaths[i].total > distinctPaths[j].total
	})
	fmt.Println(len(distinctPaths), distinctPaths[0])

	maxPairSum := 0
	for i := 0; i < 50000; i++ {
		vis1 := make(map[string]bool)
		for j := 0; j < len(distinctPaths[i].vis); j++ {
			vis1[distinctPaths[i].vis[j]] = true
		}

		for j := i + 1; j < 50001; j++ {
			localSum := distinctPaths[i].total + distinctPaths[j].total

			if localSum > maxPairSum && pathsDistinct(vis1, distinctPaths[j].vis) {

				if maxPairSum < localSum {
					maxPairSum = localSum

					fmt.Println(distinctPaths[i], distinctPaths[j], maxPairSum)
				}
			}
		}
	}

	fmt.Println(maxPairSum)
}

var distinctPaths = make([]path, 0)

func maxFlow2(current int, time int, dist [][]int, visited []int, nonZero map[int]bool, flow int, valves map[string]valve, iToValve map[int]string) int {
	if time <= 0 {
		paths++
		if globalMax < flow {
			globalMax = flow
		}

		str := make([]string, 0)
		for i := 0; i < len(visited); i++ {
			if visited[i] > 0 {
				str = append(str, valves[iToValve[i]].valveName)
			}
		}
		distinctPaths = append(distinctPaths, path{str, flow})

		return flow
	}

	flow += valves[iToValve[current]].flowRate * time
	visited[current] = time

	maximum := flow
	for i := 0; i < len(dist[current]); i++ {
		if visited[i] == 0 && nonZero[i] {
			newVis := make([]int, len(visited))
			copy(newVis, visited)
			maximum = max(maximum, maxFlow2(i, time-dist[current][i]-1, dist, newVis, nonZero, flow, valves, iToValve))
		}
	}

	if globalMax < flow {
		globalMax = flow
	}

	str := make([]string, 0)
	for i := 0; i < len(visited); i++ {
		if visited[i] > 0 {
			str = append(str, valves[iToValve[i]].valveName)
		}
	}
	distinctPaths = append(distinctPaths, path{str, flow})

	return flow
}

type path struct {
	vis   []string
	total int
}

func pathsDistinct(s1 map[string]bool, s2 []string) bool {
	for i := 0; i < len(s2); i++ {
		if s2[i] != "AA" && s1[s2[i]] {
			return false
		}
	}
	return true
}
