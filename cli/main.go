package main

import (
	"bytes"
	"flag"
	"fmt"
	npuzzle "github.com/terry-finkel/nPuzzle/src"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

//Variables for command-line display
var matrixDisplayWidth int
var itemDisplayWidth int

//Flags
var d = flag.Bool("display", true, "display solution")
var h = flag.String("heuristic", "manhattan + linear conflicts", "choose between available heuristics: [\"hamming\", \"manhattan\", \"manhattan + linear conflicts\"]")
var s = flag.String("search", "uniform-cost", "choose between [\"uniform-cost\", \"greedy\"]")
var t = flag.Duration("timeout", 60*time.Second, "set a time out duration")

func displayPath(it *npuzzle.Item) {

	if it.Parent != nil {
		displayPath(it.Parent)
		fmt.Printf("%v|\n", strings.Repeat(" ", matrixDisplayWidth/2))
		fmt.Printf("%vv\n", strings.Repeat(" ", matrixDisplayWidth/2))
	}

	display(it.M)
}

func display(m npuzzle.Matrix) {

	var buff bytes.Buffer
	buff.WriteString(fmt.Sprintf("%v\n", strings.Repeat("-", matrixDisplayWidth)))
	for _, row := range m {
		buff.WriteString("|")
		for k, num := range row {
			if k != 0 {
				buff.WriteString(" ")
			}
			if num == 0 {
				buff.WriteString("\033[1;33m")
			}
			buff.WriteString(fmt.Sprintf("%*v", itemDisplayWidth, num))
			if num == 0 {
				buff.WriteString("\033[0m")
			}
		}

		buff.WriteString("|\n")
	}

	buff.WriteString(strings.Repeat("-", matrixDisplayWidth))
	fmt.Println(buff.String())
}

func main() {

	flag.Parse()

	/* Choose heuristic function */
	var heuristic npuzzle.Heuristic
	switch *h {
	case "hamming":
		heuristic = npuzzle.HammingDistance
	case "manhattan":
		heuristic = npuzzle.ManhattanDistance
	case "manhattan + linear conflicts":
		heuristic = npuzzle.ManhattanPlusLinearConflicts
	default:
		log.Fatalf("error: %#v isn't recognized as a valid heuristic function\nAvailable heuristics are:\n"+
			" - hamming\n - manhattan\n - manhattan + linear conflicts (default)\n", *h)
	}

	/* Choose between greedy search and uniform-cost search */
	var goal npuzzle.Goal
	var search npuzzle.Search
	switch *s {
	case "greedy":
		goal = npuzzle.GreedyGoalReached
		search = npuzzle.GreedySearch
	case "uniform-cost":
		goal = npuzzle.UniformCostGoalReached
		search = npuzzle.UniformCostSearch
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

	m, err := npuzzle.ParseMatrix(data)
	if err != nil {
		log.Fatal(err)
	}

	if npuzzle.IsSolvable(m) == false {
		log.Fatal("unsolvable")
	}

	/* Calculate width in pixels for our matrix display */
	size := len(m)
	itemDisplayWidth = len(strconv.Itoa(size*size - 1))
	matrixDisplayWidth = itemDisplayWidth*size + size + 1

	fmt.Printf("starting solve on %v...\n", m)
	begin := time.Now()
	res, totalNumberOfStates, maxNumberOfStates, err := m.Solve(heuristic, search, goal, *t)
	if err != nil {
		log.Fatal(err)
	}

	duration := time.Since(begin)
	fmt.Printf("- solved in: %v\n", duration)
	fmt.Printf("- heuristic used: %v\n", *h)
	fmt.Printf("- search: %v\n", *s)
	fmt.Printf("- total number of states selected: %v\n", totalNumberOfStates)
	fmt.Printf("- maximum number of states in memory: %v\n", maxNumberOfStates)
	fmt.Printf("- number of moves: %v\n", res.Cost)
	if res.Parent == nil {
		fmt.Printf("- solving sequence: already solved!\n")
	} else {
		fmt.Printf("- solving sequence: %v\n", npuzzle.StringifyPath(res))
	}

	if *d == true {
		fmt.Printf("- path:\n")
		displayPath(res)
	}
}
