package solutions2023

import (
	"fmt"
	"sort"
	"strings"
)

func Day25p1(lines []string) {
	graph := make(map[string][]string)

	for i := range lines {
		tmp := strings.Split(lines[i], ": ")
		graph[tmp[0]] = strings.Split(tmp[1], " ")
	}

	for k, v := range graph { //make all edges bidirectional
		for i := range v {
			graph[v[i]] = append(graph[v[i]], k)
		}
	}

	prev := ""
	for k, _ := range graph {
		if prev != "" {
			vis := make(map[string]string)
			BFS(graph, vis, k, prev)
		}
		prev = k
	}

	//freq is updated
	freqArr := make([]nodeFreq, len(freq))
	i := 0
	for k, v := range freq {
		freqArr[i] = nodeFreq{k, v}
		i++
	}

	sort.SliceStable(freqArr, func(i, j int) bool {
		return freqArr[i].freq > freqArr[j].freq
	})

	fmt.Println(freqArr)

	//nodes 0,1 - 2,3 - 4,5 in sorted freqArr should represent 3 most traversed edges, remove from graph
	for j := 0; j < 6; j += 2 {
		for index, v := range graph[freqArr[j].node] {
			if v == freqArr[j+1].node {
				graph[freqArr[j].node] = removeEdge(graph[freqArr[j].node], index)
			}
		}

		for index, v := range graph[freqArr[j+1].node] {
			if v == freqArr[j].node {
				graph[freqArr[j+1].node] = removeEdge(graph[freqArr[j+1].node], index)
			}
		}
	}

	//find different groups via dfs on any random node
	group1 := 0
	for k, _ := range graph {
		vis := make(map[string]bool)
		stringDFS(graph, vis, k)
		group1 = len(vis)
		break
	}

	fmt.Println("total", (len(graph)-group1)*group1)
}

type nodeFreq struct {
	node string
	freq int
}

func removeEdge(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func stringDFS(graph map[string][]string, visited map[string]bool, node string) {
	visited[node] = true
	for _, v := range graph[node] {
		if !visited[v] {
			stringDFS(graph, visited, v)
		}
	}
}

var freq = make(map[string]int)

func BFS(graph map[string][]string, visited map[string]string, start, end string) {
	q := make([]string, 0)
	q = append(q, start)
	visited[start] = "root"
	freq[start]++

	for len(q) > 0 {
		curr, newQ := q[0], q[1:]

		if curr == end { //found shortest path
			for visited[curr] != "root" { //traverse back up the visited arr to root, incr path node freq
				freq[curr]++
				curr = visited[curr]
			}
			return
		}

		for _, v := range graph[curr] { // add neighbours to q
			_, ok := visited[v]
			if !ok {
				visited[v] = curr //store parent in visited arr
				newQ = append(newQ, v)
			}
		}
		q = newQ
	}
}
