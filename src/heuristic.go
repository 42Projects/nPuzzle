package main

import "math"

//Heuristic represents a heuristic function
type Heuristic func(Matrix, Matrix) int

//HammingDistance add one move for each tile not in place, it's a pretty weak heuristic
func hammingDistance(m1, m2 Matrix) (hammingDistance int) {

	hammingDistance = 0
	for k, row := range m1 {
		for p, num := range row {
			if num != 0 && num != m2[k][p] {
				hammingDistance++
			}
		}
	}

	return
}

//ManhattanDistance accounts for the distance for each tile to reach it's desired position
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
