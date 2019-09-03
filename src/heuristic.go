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

func getIndex(num int, row []int) int {

	for index, it := range row {
		if it == num {
			return index
		}
	}

	return -1
}

func inlineConflicts(current, goal []int) int {

	inlineConflicts := 0
	for c1, num1 := range current {
		offset := c1 + 1
		for c2, num2 := range current[offset:] {
			if num1 != 0 && num2 != 0 {
				g1 := getIndex(num1, goal)
				g2 := getIndex(num2, goal)

				/* We need to adjust c2 position as it results from the slice of a line */
				if g1 != -1 && g2 != -1 && (c1 < (c2+offset)) != (g1 < g2) {
					inlineConflicts++
				}
			}
		}
	}

	return inlineConflicts
}

//LinearConflicts calculates conflicts when two tiles are in the same column or row (and their final one) but their goal position is inverted
func linearConflicts(m1, m2 Matrix) (linearConflict int) {

	linearConflict = 0
	for k, row := range m1 {
		linearConflict += inlineConflicts(row, m2[k])
	}

	size := len(m1)
	for k := 0; k < size; k++ {
		column := make([]int, size)
		goal := make([]int, size)
		for p := 0; p < size; p++ {
			column[p] = m1[p][k]
			goal[p] = m2[p][k]
		}

		linearConflict += inlineConflicts(column, goal)
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

//Cumulate manhattan distance heuristic with the cost of linear conflicts (which adds at least 2 moves to the solution)
func manhattanPlusLinearConflicts(m1, m2 Matrix) (distance int) {

	return manhattanDistance(m1, m2) + 2*linearConflicts(m1, m2)
}
