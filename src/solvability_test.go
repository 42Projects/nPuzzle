package main

import "testing"

var isSolvableTestCases = []struct {
	input    Matrix
	expected bool
}{
	0: {Matrix{{7, 1, 0}, {4, 6, 2}, {3, 5, 8}}, true},
	1: {Matrix{{3, 6, 7}, {4, 0, 2}, {5, 1, 8}}, false},
	2: {Matrix{{13, 10, 2, 11}, {12, 14, 1, 15}, {9, 5, 8, 0}, {6, 7, 3, 4}}, true},
	3: {Matrix{{23, 8, 22, 12, 14}, {1, 0, 24, 4, 11}, {2, 10, 20, 18, 5}, {9, 17, 19, 13, 7}, {15, 21, 16, 6, 3}}, false},
	4: {Matrix{{12, 23, 9, 0, 20, 6}, {13, 25, 15, 22, 5, 34}, {35, 10, 1, 28, 31, 18}, {29, 14, 27, 19, 30, 11}, {21, 2, 32, 4, 16, 8}, {26, 33, 7, 3, 24, 17}}, true},
	5: {Matrix{{33, 35, 9, 63, 28, 13, 14, 65, 20}, {58, 6, 1, 8, 29, 31, 41, 73, 23}, {25, 57, 34, 5, 46, 37, 51, 39, 64}, {4, 80, 32, 72, 47, 79, 62, 15, 24}, {53, 12, 50, 26, 2, 18, 30, 74, 0}, {70, 66, 21, 60, 42, 69, 11, 78, 75}, {48, 49, 45, 22, 36, 76, 40, 52, 43}, {56, 68, 54, 19, 3, 55, 38, 44, 10}, {59, 67, 27, 77, 16, 71, 7, 17, 61}}, false},
}

func TestIsSolvable(t *testing.T) {

	for _, it := range isSolvableTestCases {
		res := isSolvable(it.input)
		if res != it.expected {
			t.Errorf("Entry %v, expected %v, got %v", it.input, it.expected, res)
		}
	}
}
