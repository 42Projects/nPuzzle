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

//OpenSet is implemented using a priority queue with the help of the heap interface
type OpenSet []*Item

func (os OpenSet) Len() int           { return len(os) }
func (os OpenSet) Less(k, p int) bool { return os[k].priority < os[p].priority }
func (os OpenSet) Swap(k, p int) {

	os[k], os[p] = os[p], os[k]
	os[k].index = k
	os[p].index = p
}

//Pop removes last item from the priority queue
func (os *OpenSet) Pop() interface{} {

	old := *os
	n := len(old)
	item := old[n-1]
	item.index = -1
	*os = old[0 : n-1]
	return item
}

//Push insert an item into the priority queue, priority is handled under the hood by the heap interface
func (os *OpenSet) Push(x interface{}) {

	n := len(*os)
	item := x.(*Item)
	item.index = n
	*os = append(*os, item)
}

//Remove removes an item from the priority queue based on it's index
func (os *OpenSet) Remove(x interface{}) {

	index := x.(*Item).index
	heap.Remove(os, index)
}
