package solutions2023

import (
	"container/heap"
	"fmt"
)

func Day17p2(lines []string) {
	dist := make(map[State]int)
	start := Point{0, 0}
	r := State{start, east, 0}
	d := State{start, south, 0}
	dist[r] = 0
	dist[d] = 0
	pq := PriorityQueue{
		&Item{r, 0, 0},
		&Item{d, 0, 1},
	}
	heap.Init(&pq)

	reverse := map[Point]Point{
		north: south,
		south: north,
		east:  west,
		west:  east,
	}

	low := 20000000
	//start dijkstra
	for pq.Len() > 0 {
		curr := pq.Pop().(*Item)

		//if we already reached this state faster, skip
		if dist[curr.value] < curr.priority {
			continue
		}
		//if we are at end node
		if curr.value.point.y == len(lines)-1 && curr.value.point.x == len(lines[0])-1 && curr.value.currMoves >= 4 {
			low = min(curr.priority, low)
		}

		//get next valid moves
		for _, dir := range [4]Point{north, south, east, west} {
			if dir == reverse[curr.value.currDir] || (curr.value.currDir == dir && curr.value.currMoves >= 10) {
				continue
			}
			//get new cords and check inbounds
			newY, newX := curr.value.point.y+dir.y, curr.value.point.x+dir.x
			if newX >= 0 && newY >= 0 && newY < len(lines) && newX < len(lines[0]) {

				newMoves := curr.value.currMoves
				if dir == curr.value.currDir {
					newMoves = curr.value.currMoves + 1
				} else {
					if newMoves < 4 {
						continue
					} else {
						newMoves = 1
					}
				}
				newState := State{Point{newY, newX}, dir, newMoves}
				newHeat := int(lines[newY][newX]-'0') + curr.priority

				//check if state exists in dist and check if we have new lowest
				prevHeat, ok := dist[newState]
				if ok && prevHeat <= newHeat {
					continue
				}
				dist[newState] = newHeat
				heap.Push(&pq, &Item{value: newState, priority: newHeat})
			}
		}
	}

	fmt.Println("end", low)
}
