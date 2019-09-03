package main

import "testing"

var hammingDistanceTestCases = []struct {
	m1, m2   Matrix
	expected int
}{
	0: {Matrix{{1, 2, 3}, {8, 0, 4}, {7, 6, 5}}, Matrix{{1, 2, 3}, {8, 0, 4}, {7, 6, 5}}, 0},
	1: {Matrix{{1, 2, 3}, {0, 8, 4}, {7, 6, 5}}, Matrix{{1, 2, 3}, {8, 0, 4}, {7, 6, 5}}, 1},
	2: {Matrix{{2, 7, 3}, {4, 1, 5}, {0, 6, 8}}, Matrix{{1, 2, 3}, {8, 0, 4}, {7, 6, 5}}, 6},
	3: {Matrix{{2, 7, 6}, {4, 1, 5}, {0, 3, 8}}, Matrix{{1, 2, 3}, {8, 0, 4}, {7, 6, 5}}, 8},
}

var manhattanDistanceTestCases = []struct {
	m1, m2   Matrix
	expected int
}{
	0: {Matrix{{1, 2, 3}, {8, 0, 4}, {7, 6, 5}}, Matrix{{1, 2, 3}, {8, 0, 4}, {7, 6, 5}}, 0},
	1: {Matrix{{1, 2, 3}, {0, 8, 4}, {7, 6, 5}}, Matrix{{1, 2, 3}, {8, 0, 4}, {7, 6, 5}}, 1},
	2: {Matrix{{1, 2, 5}, {3, 0, 6}, {7, 4, 8}}, Matrix{{1, 2, 3}, {4, 5, 6}, {7, 8, 0}}, 8},
	3: {Matrix{{2, 7, 3}, {4, 1, 5}, {0, 6, 8}}, Matrix{{1, 2, 3}, {8, 0, 4}, {7, 6, 5}}, 12},
	4: {Matrix{{2, 7, 6}, {4, 1, 5}, {0, 3, 8}}, Matrix{{1, 2, 3}, {8, 0, 4}, {7, 6, 5}}, 18},
}

func TestHammingDistance(t *testing.T) {

	for _, it := range hammingDistanceTestCases {
		res := hammingDistance(it.m1, it.m2)
		if res != it.expected {
			t.Errorf("Entry %v and %v, got %v, expected %v", it.m1, it.m2, res, it.expected)
		}
	}
}

func TestManhattanDistance(t *testing.T) {

	for _, it := range manhattanDistanceTestCases {
		res := manhattanDistance(it.m1, it.m2)
		if res != it.expected {
			t.Errorf("Entry %v and %v, got %v, expected %v", it.m1, it.m2, res, it.expected)
		}
	}
}