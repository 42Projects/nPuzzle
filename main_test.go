package main

import (
"testing"
)

var validMatrixTests = []struct {
	input string
	expected Matrix
}{
	0: {"3\n0 1 2\n3 4 5\n6 7 8", Matrix{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}}},
	1: {"3\n0 1 2\n3 4 5\n6 7 8\n\n#comment", Matrix{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}}},
	2: {"4\n15  1 7 8\n2 11   13 5\n9 4 6    12\n3  0  10  14 #comment", Matrix{{15, 1, 7, 8}, {2, 11, 13, 5}, {9, 4, 6, 12}, {3, 0, 10, 14}}},
}

var invalidMatrixTests = []struct {
	input string
}{
	0: {"1"},
	1: {"A"},
	2: {"3\n1 2 3\n4 17 16\n20 20 20"},
	3: {"1 2 3#comment"},
}

func TestParsing(t *testing.T) {

	for _, it := range validMatrixTests {
		res, err := parseFile(it.input)
		if err != nil {
			t.Fatalf("Valid input \"%v\", expected nil error, got: %v", it.input, err)
		}

		for k, row := range res {
			for p, col := range row {
				if col != it.expected[k][p] {
					t.Fatalf("Valid input \"%v\": got %v, wanted %v", it.input, res, it.expected)
				}
			}
		}
	}

	for _, it := range invalidMatrixTests {
		res, _ := parseFile(it.input)
		if res != nil {
			t.Fatalf("Invalid input \"%v\", expected error, got: %v", it.input, res)
		}
	}
}
