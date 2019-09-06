package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/terry-finkel/nPuzzle/proto"
	npuzzle "github.com/terry-finkel/nPuzzle/src"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

var port = flag.Int("port", 9090, "the server port")

type server struct{}

func (s *server) Greets(ctx context.Context, message *pb.Message) (*pb.Message, error) {

	log.Printf("client handshake: %v\n", message.Message)
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

	log.Printf("received problem: %v\n", problem)

	/* Choose heuristic function */
	var heuristic npuzzle.Heuristic
	switch problem.Heuristic {
	case "hamming":
		heuristic = npuzzle.HammingDistance
	case "manhattan":
		heuristic = npuzzle.ManhattanDistance
	case "manhattan + linear conflicts":
		heuristic = npuzzle.ManhattanPlusLinearConflicts
	default:
		return &pb.Result{
			Success: false,
			Error: fmt.Sprintf("error: %#v isn't recognized as a valid heuristic function\nAvailable heuristics "+
				"are:\n - hamming\n - manhattan\n - manhattan + linear conflicts (default)\n", problem.Heuristic),
		}, nil
	}

	/* Choose between greedy search and uniform-cost search */
	var search npuzzle.Search
	switch problem.Search {
	case "greedy":
		search = npuzzle.GreedySearch
	case "uniform-cost":
		search = npuzzle.UniformCostSearch
	default:
		return &pb.Result{
			Success: false,
			Error: fmt.Sprintf("error: %#v isn't recognized as a valid search option\nAvailable search "+
				"options are:\n - greedy\n - uniform-cost (default)", problem.Search),
		}, nil
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

	if npuzzle.IsSolvable(m) == false {
		log.Println("failed to solve problem: unsolvable")
		return &pb.Result{
			Success: false,
			Error:   "unsolvable",
		}, nil
	}

	begin := time.Now()
	log.Printf("starting solve on %v...", m)
	res, totalNumberOfStates, maxNumberOfStates := m.Solve(problem.Search, heuristic, search)
	duration := time.Since(begin)
	log.Printf("solved %v in %v seconds", m, duration)

	return &pb.Result{
		Success:     true,
		Time:        duration.String(),
		TotalStates: int32(totalNumberOfStates),
		MaxStates:   int32(maxNumberOfStates),
		Moves:       int32(res.Cost),
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
