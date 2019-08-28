package main

import "testing"

var getTilePositionTestCases = []struct {
	input    Matrix
	expected Point
}{
	0: {Matrix{{2, 3, 0}, {1, 4, 5}, {7, 6, 8}}, Point{2, 0}},
	1: {Matrix{{2, 3, 7}, {1, 4, 5}, {0, 6, 8}}, Point{0, 2}},
	2: {Matrix{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 0, 15}}, Point{2, 3}},
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
	0: {Matrix{{2, 3, 0}, {1, 4, 5}, {7, 6, 8}}, nil},
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

func TestGetTilePosition(t *testing.T) {

	for _, it := range getTilePositionTestCases {
		res := it.input.getTilePosition()
		if PointsMatch(res, it.expected) == false {
			t.Fatalf("Entry %v, got %v, expected %v", it.input, res, it.expected)
		}
	}
}

func TestSlide(t *testing.T) {

	for _, it := range slideDownTestCases {
		res := it.input.slideDown()
		if MatricesMatch(res, it.expected) == false {
			t.Fatalf("SlideDown: entry %v, got %v, expected %v", it.input, res, it.expected)
		}
	}

	for _, it := range slideLeftTestCases {
		res := it.input.slideLeft()
		if MatricesMatch(res, it.expected) == false {
			t.Fatalf("SlideDown: entry %v, got %v, expected %v", it.input, res, it.expected)
		}
	}

	for _, it := range slideRightTestCases {
		res := it.input.slideRight()
		if MatricesMatch(res, it.expected) == false {
			t.Fatalf("SlideDown: entry %v, got %v, expected %v", it.input, res, it.expected)
		}
	}

	for _, it := range slideUpTestCases {
		res := it.input.slideUp()
		if MatricesMatch(res, it.expected) == false {
			t.Fatalf("SlideUp: entry %v, got %v, expected %v", it.input, res, it.expected)
		}
	}
}
