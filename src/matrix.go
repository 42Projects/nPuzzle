package npuzzle

import (
	"bytes"
	"container/heap"
	"fmt"
	"strconv"
	"time"
)

//Direction represents cardinal points
type Direction int

//Point represents a point in a 2D plane, x is the abscissa, y the ordinate
type Point struct {
	x, y int
}

//Matrix is a 2d array representation
type Matrix [][]int

//Directions
const (
	_ Direction = iota
	Down
	Left
	Right
	Up
)

func (m Matrix) getTilePosition(tile int) Point {

	for y, row := range m {
		for x, it := range row {
			if it == tile {
				return Point{x, y}
			}
		}
	}

	return Point{-1, -1}
}

func (m Matrix) slide(direction Direction) Matrix {

	slid := make(Matrix, len(m))
	p := Point{}
	for y, row := range m {
		slid[y] = make([]int, len(m))
		for x, num := range row {
			if num == 0 {
				var outOfBound bool
				switch direction {
				case Down:
					outOfBound = y == len(m)-1
				case Left:
					outOfBound = x == 0
				case Right:
					outOfBound = x == len(m)-1
				case Up:
					outOfBound = y == 0
				}

				if outOfBound == true {
					return nil
				}

				p.x = x
				p.y = y
			}

			slid[y][x] = num
		}
	}

	switch direction {
	case Down:
		slid[p.y][p.x], slid[p.y+1][p.x] = slid[p.y+1][p.x], slid[p.y][p.x]
	case Left:
		slid[p.y][p.x], slid[p.y][p.x-1] = slid[p.y][p.x-1], slid[p.y][p.x]
	case Right:
		slid[p.y][p.x], slid[p.y][p.x+1] = slid[p.y][p.x+1], slid[p.y][p.x]
	case Up:
		slid[p.y][p.x], slid[p.y-1][p.x] = slid[p.y-1][p.x], slid[p.y][p.x]
	}

	return slid
}

//Add a string representation to our matrix so we can hash it for the closed set
func (m Matrix) string() string {

	var buff bytes.Buffer
	for k, row := range m {
		if k != 0 {
			buff.WriteByte('\n')
		}

		for p, num := range row {
			if p != 0 {
				buff.WriteByte(' ')
			}

			buff.WriteString(strconv.Itoa(num))
		}
	}

	return buff.String()
}

func generateEndState(size int) Matrix {

	m := make(Matrix, size)
	for k := 0; k < size; k++ {
		m[k] = make([]int, size)
	}

	x, y, direction := 0, 0, Right
	for k := 1; k < size*size; k++ {
		m[y][x] = k
		switch direction {
		case Down:
			if y == size-1 || m[y+1][x] != 0 {
				direction = Left
				x--
			} else {
				y++
			}
		case Left:
			if x == 0 || m[y][x-1] != 0 {
				direction = Up
				y--
			} else {
				x--
			}
		case Right:
			if x == size-1 || m[y][x+1] != 0 {
				direction = Down
				y++
			} else {
				x++
			}
		case Up:
			if y == 0 || m[y-1][x] != 0 {
				direction = Right
				x++
			} else {
				y--
			}
		}
	}

	return m
}

//Goal represents a boolean returning function to check whether goal is reached or not
type Goal func(int, int) bool

//GreedyGoalReached is the goal function for greedy search type
func GreedyGoalReached(cost, priority int) bool { return priority == 0 }

//UniformCostGoalReached is the goal function for uniform-cost search type
func UniformCostGoalReached(cost, priority int) bool { return cost == priority }

//Solve implements A* algorithm
func (m Matrix) Solve(heuristic Heuristic, search Search, goal Goal, timeout time.Duration) (*Item, int, int, error) {

	timer := time.NewTimer(timeout)
	defer timer.Stop()

	closedSet := ClosedSet{}
	openSet := make(OpenSet, 0)
	heap.Init(&openSet)
	endState := generateEndState(len(m))
	current := &Item{
		M:        m,
		Cost:     0,
		priority: heuristic(m, endState),
		Parent:   nil,
	}
	heap.Push(&openSet, current)

	totalNumberOfStates := 1
	maxNumberOfStates := 1
	for openSet.Len() > 0 && goal(current.Cost, current.priority) == false {

		select {
		case <-timer.C:
			return nil, -1, -1, fmt.Errorf("timed out after %v", timeout)
		default:
		}

		closedSet[current.M.string()] = current
		search(current, heuristic, &openSet, closedSet, endState)

		/* We keep track of memory complexity */
		if size := openSet.Len(); size > maxNumberOfStates {
			maxNumberOfStates = size
		}

		current = heap.Pop(&openSet).(*Item)
		totalNumberOfStates++
	}

	return current, totalNumberOfStates, maxNumberOfStates, nil
}
