package npuzzle

import (
	"container/heap"
	"math"
)

//Search represents the function that search for the most promising moves (or "neighbours")
type Search func(*Item, Heuristic, *OpenSet, ClosedSet, Matrix)

//Neighbour is a simple struct containing one neighbour and his heuristic score, only used for greedy search
type Neighbour struct {
	m     Matrix
	score int
}

//GreedySearch targets the shortest path from each node disregarding everything else
func GreedySearch(current *Item, heuristic Heuristic, openSet *OpenSet, closedSet ClosedSet, endState Matrix) {

	bestScore := math.MaxInt32
	cost := current.Cost + 1
	var neighbours []Neighbour

	/* Only keep the most promising neighbours */
	for _, direction := range []Direction{Up, Down, Left, Right} {
		if neighbour := current.M.slide(direction); neighbour != nil && closedSet[neighbour.string()] == nil {
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
				M:        it.m,
				Cost:     cost,
				priority: bestScore,
				Parent:   current,
			})
		}
	}
}

//UniformCostSearch considers every path but will prioritize the one with the lesser overall cost
func UniformCostSearch(current *Item, heuristic Heuristic, openSet *OpenSet, closedSet ClosedSet, endState Matrix) {

	/* Add all neighbours to the priority queue, best will be explored first */
	for _, direction := range []Direction{Up, Down, Left, Right} {
		if neighbour := current.M.slide(direction); neighbour != nil && closedSet[neighbour.string()] == nil {
			cost := current.Cost + 1
			priority := cost + heuristic(neighbour, endState)
			heap.Push(openSet, &Item{
				M:        neighbour,
				Cost:     cost,
				priority: priority,
				Parent:   current,
			})
		}
	}
}
