package main

import "container/heap"

type Item struct {
	m                     Matrix
	cost, index, priority int
	parent                *Item
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(k, p int) bool { return pq[k].priority > pq[p].priority }
func (pq PriorityQueue) Swap(k, p int) {

	pq[k], pq[p] = pq[p], pq[k]
	pq[k].index = k
	pq[p].index = p
}

func (pq *PriorityQueue) Pop() interface{} {

	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Push(x interface{}) {

	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Remove(x interface{}) {

	index := x.(*Item).index
	heap.Remove(pq, index)
}
