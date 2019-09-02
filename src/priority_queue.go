package main

import "container/heap"

//ClosedSet is a hash map of already explored matrices
type ClosedSet map[string]*Item

//Item is our data structure for the priorityQueue
type Item struct {
	m                     Matrix
	cost, index, priority int
	parent                *Item
}

func (it *Item) addNeighboursToQueue(openSet *PriorityQueue, closedSet ClosedSet, endState Matrix) {

	for _, direction := range []Direction{Up, Down, Left, Right} {
		if neighbour := it.m.slide(direction); neighbour != nil && closedSet[neighbour.string()] == nil {
			cost := it.cost + 1
			priority := cost + heuristic(neighbour, endState)
			heap.Push(openSet, &Item{
				m:        neighbour,
				cost:     cost,
				priority: priority,
				parent:   it,
			})
		}
	}
}

//PriorityQueue is implemented using heap interface
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(k, p int) bool { return pq[k].priority < pq[p].priority }
func (pq PriorityQueue) Swap(k, p int) {

	pq[k], pq[p] = pq[p], pq[k]
	pq[k].index = k
	pq[p].index = p
}

//Pop removes last item from the priorityQueue
func (pq *PriorityQueue) Pop() interface{} {

	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

//Push insert an item into the priorityQueue, priority is handled under the hood by the heap interface
func (pq *PriorityQueue) Push(x interface{}) {

	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

//Remove removes an item from the priorityQueue based on it's index
func (pq *PriorityQueue) Remove(x interface{}) {

	index := x.(*Item).index
	heap.Remove(pq, index)
}
