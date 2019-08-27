package main

import (
	"errors"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type Matrix [][]int

func parseRow(line string, size int, checker *[]bool) ([]int, error) {

	parsed, row, numbers := 0, make([]int, size), strings.Split(line, " ")
	for _, number := range numbers {
		if number == "" {
			continue
		} else if number[0] == '#' {
			break
		}

		n, err := strconv.Atoi(number)

		/* Parsing error cases */
		switch {
		case err != nil:
			return nil, errors.New("invalid file: invalid number")
		case n < 0 || n > size*size-1:
			return nil, errors.New("invalid file: number out of range")
		case n < len(*checker) && (*checker)[n] == true:
			return nil, errors.New("invalid file: duplicate element")
		}

		if n < len(*checker) {
			(*checker)[n] = true
		}

		if parsed < size {
			row[parsed] = n
		}

		parsed += 1
	}

	if parsed != size {
		return nil, errors.New("invalid file: row with not enough / too many elements")
	}

	return row, nil
}

func parseFile(s string) (Matrix, error) {

	var m Matrix
	var checker []bool
	k, size, lines := 0, 0, strings.Split(s, "\n")
	for _, line := range lines {
		switch {

		/* Skip comments */
		case line == "" || line[0] == '#':
			continue

		/* Create matrix with first valid line, throw an error if size isn't valid or present */
		case size == 0:
			var err error
			size, err = strconv.Atoi(line)
			if err != nil {
				return nil, errors.New("invalid file: expected size")
			}
			if size < 3 {
				return nil, errors.New("invalid file: grid is too small")
			}
			m = make([][]int, size)
			checker = make([]bool, size * size)

		/* Parse lines */
		default:
			row, err := parseRow(line, size, &checker)
			if err != nil {
				return nil, err
			}
			m[k] = row
			k += 1
		}
	}

	return m, nil
}

func (m Matrix) solve() {}

func main() {

	var data string

	flag.Parse()
	switch flag.NArg() {
	case 0:
		p, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		data = string(p)
	case 1:
		p, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		data = string(p)
	default:
		log.Fatalf("Please open a valid npz file or use the standard entry")
	}

	m, err := parseFile(data)
	if err != nil {
		log.Fatal(err)
	}

	m.solve()
}
