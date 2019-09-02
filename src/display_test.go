package main

import "testing"

var stringTestCases = []struct {
	input    Matrix
	expected string
}{
	0: {Matrix{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}}, "0 1 2\n3 4 5\n6 7 8"},
	1: {Matrix{{15, 1, 7, 8}, {2, 11, 13, 5}, {9, 4, 6, 12}, {3, 0, 10, 14}}, "15 1 7 8\n2 11 13 5\n9 4 6 12\n3 0 10 14"},
}

func TestString(t *testing.T) {

	for _, it := range stringTestCases {
		res := it.input.string()
		if res != it.expected {
			t.Errorf("Entry %v, expected %#v, got %#v", it.input, it.expected, res)
		}
	}
}
