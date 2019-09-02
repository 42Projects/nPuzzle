package main

import "testing"

var generateEndStateTestCases = []struct {
	size     int
	expected Matrix
}{
	0: {3, Matrix{{1, 2, 3}, {8, 0, 4}, {7, 6, 5}}},
	1: {4, Matrix{{1, 2, 3, 4}, {12, 13, 14, 5}, {11, 0, 15, 6}, {10, 9, 8, 7}}},
	2: {5, Matrix{{1, 2, 3, 4, 5}, {16, 17, 18, 19, 6}, {15, 24, 0, 20, 7}, {14, 23, 22, 21, 8}, {13, 12, 11, 10, 9}}},
}

var getTilePositionTestCases = []struct {
	input    Matrix
	tile     int
	expected Point
}{
	0: {Matrix{{2, 3, 0}, {1, 4, 5}, {7, 6, 8}}, 7, Point{0, 2}},
	1: {Matrix{{2, 3, 7}, {1, 4, 5}, {0, 6, 8}}, 5, Point{2, 1}},
	2: {Matrix{{2, 3, 0}, {1, 4, 5}, {7, 6, 8}}, 0, Point{2, 0}},
	3: {Matrix{{2, 3, 7}, {1, 4, 5}, {0, 6, 8}}, 0, Point{0, 2}},
	4: {Matrix{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 0, 15}}, 0, Point{2, 3}},
	5: {Matrix{{1, 2, 0, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 3, 15}}, 3, Point{2, 3}},
}

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

var slideDownTestCases = []struct {
	input    Matrix
	expected Matrix
}{
	0: {Matrix{{2, 3, 8}, {1, 4, 5}, {7, 6, 0}}, nil},
	1: {Matrix{{2, 3, 0}, {1, 4, 5}, {7, 6, 8}}, Matrix{{2, 3, 5}, {1, 4, 0}, {7, 6, 8}}},
	2: {Matrix{{2, 3, 8, 11}, {0, 4, 5, 12}, {15, 7, 6, 1}, {10, 9, 14, 13}}, Matrix{{2, 3, 8, 11}, {15, 4, 5, 12}, {0, 7, 6, 1}, {10, 9, 14, 13}}},
}

var slideLeftTestCases = []struct {
	input    Matrix
	expected Matrix
}{
	0: {Matrix{{0, 3, 2}, {1, 4, 5}, {7, 6, 8}}, nil},
	1: {Matrix{{2, 3, 8}, {1, 4, 5}, {7, 6, 0}}, Matrix{{2, 3, 8}, {1, 4, 5}, {7, 0, 6}}},
	2: {Matrix{{2, 3, 8, 11}, {12, 4, 5, 0}, {15, 7, 6, 1}, {10, 9, 14, 13}}, Matrix{{2, 3, 8, 11}, {12, 4, 0, 5}, {15, 7, 6, 1}, {10, 9, 14, 13}}},
}

var slideRightTestCases = []struct {
	input    Matrix
	expected Matrix
}{
	0: {Matrix{{2, 3, 5}, {1, 4, 0}, {7, 6, 8}}, nil},
	1: {Matrix{{2, 3, 1}, {4, 0, 5}, {7, 6, 8}}, Matrix{{2, 3, 1}, {4, 5, 0}, {7, 6, 8}}},
	2: {Matrix{{2, 3, 8, 11}, {0, 4, 5, 12}, {15, 7, 6, 1}, {10, 9, 14, 13}}, Matrix{{2, 3, 8, 11}, {4, 0, 5, 12}, {15, 7, 6, 1}, {10, 9, 14, 13}}},
}

var slideUpTestCases = []struct {
	input    Matrix
	expected Matrix
}{
	0: {Matrix{{2, 3, 0}, {1, 4, 5}, {7, 6, 8}}, nil},
	1: {Matrix{{2, 3, 8}, {1, 4, 5}, {7, 6, 0}}, Matrix{{2, 3, 8}, {1, 4, 0}, {7, 6, 5}}},
	2: {Matrix{{2, 3, 8, 11}, {0, 4, 5, 12}, {15, 7, 6, 1}, {10, 9, 14, 13}}, Matrix{{0, 3, 8, 11}, {2, 4, 5, 12}, {15, 7, 6, 1}, {10, 9, 14, 13}}},
}

func PointsMatch(p1, p2 Point) bool { return p1.x == p2.x && p1.y == p2.y }

func TestGenerateEndState(t *testing.T) {

	for _, it := range generateEndStateTestCases {
		res := generateEndState(it.size)
		if MatricesMatch(res, it.expected) == false {
			t.Errorf("Size %v, got %v, expected %v", it.size, res, it.expected)
		}
	}
}

func TestGetTilePosition(t *testing.T) {

	for _, it := range getTilePositionTestCases {
		res := it.input.getTilePosition(it.tile)
		if PointsMatch(res, it.expected) == false {
			t.Errorf("Tile %v entry %v, got %v, expected %v", it.tile, it.input, res, it.expected)
		}
	}
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

func TestSlide(t *testing.T) {

	for _, it := range slideDownTestCases {
		res := it.input.slide(Down)
		if MatricesMatch(res, it.expected) == false {
			t.Errorf("SlideDown: entry %v, got %v, expected %v", it.input, res, it.expected)
		}
	}

	for _, it := range slideLeftTestCases {
		res := it.input.slide(Left)
		if MatricesMatch(res, it.expected) == false {
			t.Errorf("SlideDown: entry %v, got %v, expected %v", it.input, res, it.expected)
		}
	}

	for _, it := range slideRightTestCases {
		res := it.input.slide(Right)
		if MatricesMatch(res, it.expected) == false {
			t.Errorf("SlideDown: entry %v, got %v, expected %v", it.input, res, it.expected)
		}
	}

	for _, it := range slideUpTestCases {
		res := it.input.slide(Up)
		if MatricesMatch(res, it.expected) == false {
			t.Errorf("SlideUp: entry %v, got %v, expected %v", it.input, res, it.expected)
		}
	}
}
