package solutions2023

import (
	"container/heap"
	"fmt"
)

var (
	north = Point{-1, 0}
	south = Point{1, 0}
	east  = Point{0, 1}
	west  = Point{0, -1}
)

func Day17p1(lines []string) {
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
	fmt.Println(pq)
	low := 20000000

	//start dijkstra
	for pq.Len() > 0 {
		curr := pq.Pop().(*Item)
		//fmt.Println(curr.value, curr.priority)

		//if we already reached this state faster, skip
		if dist[curr.value] < curr.priority {
			continue
		}
		//if we are at end node
		if curr.value.point.y == len(lines)-1 && curr.value.point.x == len(lines[0])-1 {
			low = min(curr.priority, low)
			//fmt.Println(low)
		}

		//get next valid moves
		for _, dir := range [4]Point{north, south, east, west} {
			if dir == reverse[curr.value.currDir] || (curr.value.currDir == dir && curr.value.currMoves == 3) {
				continue
			}
			//get new cords and check inbounds
			newY, newX := curr.value.point.y+dir.y, curr.value.point.x+dir.x
			if newX >= 0 && newY >= 0 && newY < len(lines) && newX < len(lines[0]) {
				newMoves := 1
				if dir == curr.value.currDir {
					newMoves = curr.value.currMoves + 1
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

type State struct {
	point     Point
	currDir   Point
	currMoves int
}

// An Item is something we manage in a priority queue.
type Item struct {
	value    State // The value of the item; arbitrary.
	priority int   // The priority of the item in the queue.
	index    int   // The index of the item in the heap. Needed and managed by pq methods.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value State, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
