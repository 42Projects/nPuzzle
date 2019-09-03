package main

import (
	"container/heap"
	"math"
)

//Neighbour is a simple struct containing one neighbour and his heuristic score, only used for greedy search
type Neighbour struct {
	m     Matrix
	score int
}

//Search represents the function that search for the most promising moves (or "neighbours")
type Search func(*Item, Heuristic, *OpenSet, ClosedSet, Matrix)

func greedySearch(current *Item, heuristic Heuristic, openSet *OpenSet, closedSet ClosedSet, endState Matrix) {

	bestScore := math.MaxInt32
	cost := current.cost + 1
	var neighbours []Neighbour

	/* Only keep the most promising neighbours */
	for _, direction := range []Direction{Up, Down, Left, Right} {
		if neighbour := current.m.slide(direction); neighbour != nil && closedSet[neighbour.string()] == nil {
			score := heuristic(neighbour, endState)
			neighbours = append(neighbours, Neighbour{
				m:     neighbour,
				score: score,
			})
			if score < bestScore {
				bestScore = score
			}
		}
	}

	for _, it := range neighbours {
		if it.score == bestScore {
			heap.Push(openSet, &Item{
				m:        it.m,
				cost:     cost,
				priority: bestScore,
				parent:   current,
			})
		}
	}
}

func uniformCostSearch(current *Item, heuristic Heuristic, openSet *OpenSet, closedSet ClosedSet, endState Matrix) {

	/* Add all neighbours to the priority queue, best will be explored first */
	for _, direction := range []Direction{Up, Down, Left, Right} {
		if neighbour := current.m.slide(direction); neighbour != nil && closedSet[neighbour.string()] == nil {
			cost := current.cost + 1
			priority := cost + heuristic(neighbour, endState)
			heap.Push(openSet, &Item{
				m:        neighbour,
				cost:     cost,
				priority: priority,
				parent:   current,
			})
		}
	}
}
