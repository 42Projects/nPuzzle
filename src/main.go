package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

//Variables for command-line display
var matrixDisplayWidth int
var itemDisplayWidth int

//Flags
var h = flag.String("heuristic", "manhattan + linear conflicts", "choose between available heuristics: [\"hamming\", \"manhattan\", \"manhattan + linear conflicts\"]")
var s = flag.String("search", "uniform-cost", "choose between [\"uniform-cost\", \"greedy\"]")

func parseRow(line string, size int, checker *[]bool) ([]int, error) {

	parsed, row, numbers := 0, make([]int, size), strings.Split(line, " ")
	for _, number := range numbers {
		if number == "" {
			continue
		} else if number[0] == '#' {
			break
		}

		/* Get rid of end of lines comments */
		if index := strings.Index(number, "#"); index != -1 {
			number = number[0:index]
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

		parsed++
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
			checker = make([]bool, size*size)

		/* Parse lines */
		default:
			row, err := parseRow(line, size, &checker)
			if err != nil {
				return nil, err
			}
			m[k] = row
			k++
		}
	}

	return m, nil
}

func main() {

	flag.Parse()

	/* Choose heuristic function */
	var heuristic Heuristic
	switch *h {
	case "hamming":
		heuristic = hammingDistance
	case "manhattan":
		heuristic = manhattanDistance
	case "manhattan + linear conflicts":
		heuristic = manhattanPlusLinearConflicts
	default:
		log.Fatalf("error: %#v isn't recognized as a valid heuristic function\nAvailable heuristics are:\n"+
			" - hamming\n - manhattan\n - manhattan + linear conflicts (default)\n", *h)
	}

	/* Choose between greedy search and uniform-cost search */
	var search Search
	switch *s {
	case "greedy":
		search = greedySearch
	case "uniform-cost":
		search = uniformCostSearch
	default:
		log.Fatalf("error: %#v isn't recognized as a valid search option\nAvailable search options are:\n"+
			" - greedy\n - uniform-cost (default)", *s)
	}

	var data string
	switch flag.NArg() {

	/* Will read from STDIN in case of no argument */
	case 0:
		p, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		data = string(p)

	/* Read from file */
	case 1:
		p, err := ioutil.ReadFile(flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}
		data = string(p)

	/* Only accept one file */
	default:
		log.Fatalf("please add a single valid npz file or use the standard entry")
	}

	m, err := parseFile(data)
	if err != nil {
		log.Fatal(err)
	}

	if isSolvable(m) == false {
		log.Fatal("unsolvable")
	}

	/* Calculate width in pixels for our matrix display */
	size := len(m)
	itemDisplayWidth = len(strconv.Itoa(size*size - 1))
	matrixDisplayWidth = itemDisplayWidth*size + size + 1

	res, totalNumberOfStates, maxNumberOfStates := m.solve(heuristic, search)
	fmt.Printf("\n -------------- SOLVED! ---------------\n")
	fmt.Printf("- heuristic used: %v\n", *h)
	fmt.Printf("- search: %v\n", *s)
	fmt.Printf("- total number of states selected: %v\n", totalNumberOfStates)
	fmt.Printf("- maximum number of states in memory: %v\n", maxNumberOfStates)
	fmt.Printf("- number of moves: %v\n", res.cost)
	fmt.Printf("- path:\n")
	res.displayPath()
}
