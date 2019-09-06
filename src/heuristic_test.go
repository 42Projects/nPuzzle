package npuzzle

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

var linearConflictsTestCases = []struct {
	m1, m2   Matrix
	expected int
}{
	0: {Matrix{{1, 2, 3}, {8, 0, 4}, {7, 6, 5}}, Matrix{{1, 2, 3}, {8, 0, 4}, {7, 6, 5}}, 0},
	1: {Matrix{{2, 7, 6}, {4, 1, 5}, {0, 3, 8}}, Matrix{{1, 2, 3}, {8, 0, 4}, {7, 6, 5}}, 0},
	2: {Matrix{{2, 1, 3}, {8, 0, 4}, {7, 6, 5}}, Matrix{{1, 2, 3}, {8, 0, 4}, {7, 6, 5}}, 1},
	3: {Matrix{{3, 2, 1}, {8, 0, 4}, {7, 6, 5}}, Matrix{{1, 2, 3}, {8, 0, 4}, {7, 6, 5}}, 3},
	4: {Matrix{{3, 6, 1}, {4, 0, 8}, {7, 2, 5}}, Matrix{{1, 2, 3}, {8, 0, 4}, {7, 6, 5}}, 3},
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
		res := HammingDistance(it.m1, it.m2)
		if res != it.expected {
			t.Errorf("Entry %v and %v, got %v, expected %v", it.m1, it.m2, res, it.expected)
		}
	}
}

func TestLinearConflict(t *testing.T) {

	for _, it := range linearConflictsTestCases {
		res := LinearConflicts(it.m1, it.m2)
		if res != it.expected {
			t.Errorf("Entry %v and %v, got %v, expected %v", it.m1, it.m2, res, it.expected)
		}
	}
}

func TestManhattanDistance(t *testing.T) {

	for _, it := range manhattanDistanceTestCases {
		res := ManhattanDistance(it.m1, it.m2)
		if res != it.expected {
			t.Errorf("Entry %v and %v, got %v, expected %v", it.m1, it.m2, res, it.expected)
		}
	}
}
