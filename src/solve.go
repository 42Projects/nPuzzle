package main

import (
	"container/heap"
	"math"
)

type Direction int

//Point represent a point in a 2D plane, x is the abscissa, y the ordinate
type Point struct {
	x, y int
}

const (
	_ Direction = iota
	Down
	Left
	Right
	Up
)

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
				x -= 1
			} else {
				y += 1
			}
		case Left:
			if x == 0 || m[y][x-1] != 0 {
				direction = Up
				y -= 1
			} else {
				x -= 1
			}
		case Right:
			if x == size-1 || m[y][x+1] != 0 {
				direction = Down
				y += 1
			} else {
				x += 1
			}
		case Up:
			if y == 0 || m[y-1][x] != 0 {
				direction = Right
				x += 1
			} else {
				y -= 1
			}
		}
	}

	return m
}

func hammingDistance(m1, m2 Matrix) (hammingDistance int) {

	hammingDistance = 0
	for k, row := range m1 {
		for p, num := range row {
			if num != 0 && num != m2[k][p] {
				hammingDistance += 1
			}
		}
	}

	return
}

func manhattanDistance(m1, m2 Matrix) (manhattanDistance int) {

	manhattanDistance = 0
	for k, row := range m1 {
		for p, num := range row {
			if num != 0 {
				m1pos := Point{p, k}
				m2pos := m2.getTilePosition(num)
				manhattanDistance += int(math.Abs(float64(m1pos.x-m2pos.x))) + int(math.Abs(float64(m1pos.y-m2pos.y)))
			}
		}
	}

	return
}

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
					outOfBound = y == len(m) - 1
				case Left:
					outOfBound = x == 0
				case Right:
					outOfBound = x == len(m) - 1
				case Up:
					outOfBound = y == 0
				}

				if outOfBound == true {
					return nil
				} else {
					p.x = x
					p.y = y
				}
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

func (m Matrix) solve() Matrix {

	openSet := make(PriorityQueue, 0)
	heap.Init(&openSet)

	return m
}
