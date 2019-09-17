package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/42Projects/nPuzzle/proto"
	npuzzle "github.com/42Projects/nPuzzle/src"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

var port = flag.Int("port", 9090, "the server port")

type server struct{}

func (s *server) Greets(ctx context.Context, message *pb.Message) (*pb.Message, error) {

	return &pb.Message{Message: "pong!"}, nil
}

func (s *server) Parse(ctx context.Context, message *pb.Message) (*pb.Matrix, error) {

	log.Printf("received parsing request: %#v", message.Message)
	m, err := npuzzle.ParseMatrix(message.Message)
	if err != nil {
		log.Println(err)
		return &pb.Matrix{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	rows := make([]*pb.Matrix_Row, len(m))
	for index, row := range m {

		/* We need unsigned 32bits integer for protobuf */
		uIntRow := make([]uint32, len(m))
		for rowIndex, num := range row {
			uIntRow[rowIndex] = uint32(num)
		}

		rows[index] = &pb.Matrix_Row{Num: uIntRow}
	}

	return &pb.Matrix{
		Success: true,
		Rows:    rows,
	}, nil
}

func (s *server) Solve(ctx context.Context, problem *pb.Problem) (*pb.Result, error) {

	/* Choose heuristic function */
	var heuristic npuzzle.Heuristic
	switch problem.Heuristic {
	case "hamming":
		heuristic = npuzzle.HammingDistance
	case "manhattan":
		heuristic = npuzzle.ManhattanDistance
	case "manhattan + linear conflicts":
		heuristic = npuzzle.ManhattanPlusLinearConflicts
	}

	/* Choose between greedy search and uniform-cost search */
	var goal npuzzle.Goal
	var search npuzzle.Search
	switch problem.Search {
	case "greedy":
		goal = npuzzle.GreedyGoalReached
		search = npuzzle.GreedySearch
	case "uniform-cost":
		goal = npuzzle.UniformCostGoalReached
		search = npuzzle.UniformCostSearch
	}

	/* Convert protobuf unsigned 32bits integer to regular integer */
	size := len(problem.Rows)
	m := make(npuzzle.Matrix, size)
	for y, row := range problem.Rows {
		m[y] = make([]int, size)
		for x, num := range row.Num {
			m[y][x] = int(num)
		}
	}

	log.Printf("received problem:\n - heuristic: %v\n - search: %v\n - matrix: %v\n", problem.Heuristic, problem.Search, m)
	if npuzzle.IsSolvable(m) == false {
		log.Println("failed to solve problem: unsolvable")
		return &pb.Result{
			Success: false,
			Error:   "unsolvable",
		}, nil
	}

	begin := time.Now()
	log.Printf("starting solve on %v...", m)
	res, totalNumberOfStates, maxNumberOfStates, err := m.Solve(heuristic, search, goal, 30*time.Second)
	if err != nil {
		log.Printf("timed ouf after %v", 30*time.Second)
		return &pb.Result{
			Success: false,
			Error:   fmt.Sprintf("timed ouf after %v", 30*time.Second),
		}, nil
	}

	duration := time.Since(begin)
	log.Printf("solved %v in %v seconds", m, duration)
	var path string
	if res.Parent == nil {
		path = "already solved!"
	} else {
		path = npuzzle.StringifyPath(res)
	}

	return &pb.Result{
		Success:     true,
		Time:        duration.String(),
		Moves:       int32(res.Cost),
		TotalStates: int32(totalNumberOfStates),
		MaxStates:   int32(maxNumberOfStates),
		Path:        path,
	}, nil
}

func main() {

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	s := grpc.NewServer()
	pb.RegisterNpuzzleServer(s, &server{})
	log.Printf("starting server on port %v\n", *port)
	log.Fatalf("failed to serve: %v\n", s.Serve(lis))
}
